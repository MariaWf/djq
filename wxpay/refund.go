package wxpay

import (
	"mimi/djq/config"
	"github.com/pkg/errors"
	"mimi/djq/util"
	"strconv"
	"log"
)

func Refund(payOrderNumber string, totalFee int, refundOrderNumber string,refundFee int) (Params, error) {
	var err error
	url := "https://api.mch.weixin.qq.com/secapi/pay/refund"

	appId := config.Get("wxpay_appid") // 微信公众平台应用ID
	mchId := config.Get("wxpay_mch_id") // 微信支付商户平台商户号
	apiKey := config.Get("wxpay_key") // 微信支付商户平台API密钥

	if config.Get("running_state") == "test" {
		url = "https://api.mch.weixin.qq.com/sandboxnew/secapi/pay/refund"
		apiKey, err = GetSignKey()
		if err != nil {
			return nil, errors.Wrap(err,"获取测试API_KEY失败")
		}
	}

	c := NewClient(appId, mchId, apiKey)

	// 微信支付商户平台证书路径
	certFile   := config.Get("wxpay_cert_file")
	keyFile    := config.Get("wxpay_key_file")
	rootcaFile := config.Get("wxpay_rootca_file")

	// 附着商户证书
	err = c.WithCert(certFile, keyFile, rootcaFile)
	if err != nil {
		log.Println(err)
		return nil ,errors.Wrap(err,"获取商户证书失败")
	}

	params := make(Params)

	params.SetString("appid",c.AppId)//公众账号ID	appid	是	String(32)	wx8888888888888888	微信分配的公众账号ID（企业号corpid即为此appId）
	params.SetString("mch_id",c.MchId)//商户号	mch_id	是	String(32)	1900000109	微信支付分配的商户号
	params.SetString("nonce_str",util.BuildUUID())//随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
	params.SetString("sign_type","MD5")//签名类型	sign_type	否	String(32)	HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	//微信订单号	transaction_id	二选一	String(28)	1217752501201407033233368018	微信生成的订单号，在支付通知中有返回
	params.SetString("out_trade_no",payOrderNumber)//商户订单号	out_trade_no	String(32)	1217752501201407033233368018	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
	params.SetString("out_refund_no",refundOrderNumber)//商户退款单号	out_refund_no	是	String(64)	1217752501201407033233368018	商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	params.SetString("total_fee",strconv.Itoa(totalFee))//订单金额	total_fee	是	Int	100	订单总金额，单位为分，只能为整数，详见支付金额
	params.SetString("refund_fee",strconv.Itoa(refundFee))//退款金额	refund_fee	是	Int	100	退款总金额，订单总金额，单位为分，只能为整数，详见支付金额
	params.SetString("refund_fee_type","CNY")//货币种类	refund_fee_type	否	String(8)	CNY	货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	params.SetString("refund_desc","")//退款原因	refund_desc	否	String(80)	商品已售完	若商户传入，会在下发给用户的退款消息中体现退款原因
	//退款资金来源	refund_account	否	String(30)	REFUND_SOURCE_RECHARGE_FUNDS
	//仅针对老资金流商户使用
	//REFUND_SOURCE_UNSETTLED_FUNDS---未结算资金退款（默认使用未结算资金退款）
	//REFUND_SOURCE_RECHARGE_FUNDS---可用余额退款
	params.SetString("sign", c.Sign(params))//签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法

	return c.Post(url, params, true)

}

func RefundResult(payOrderNumber string, totalFee int, refundOrderNumber string,refundFee int) (err error) {
	params, err := Refund(payOrderNumber,totalFee,refundOrderNumber,refundFee)
	if err != nil {
		return
	}
	client := NewDefaultClient()
	if params["return_code"] == "FAIL" {
		err = errors.New("退款失败")
		log.Println(errors.Wrap(err, params["return_msg"]))
	} else if !client.CheckSign(params) {
		err = errors.New("退款失败")
		log.Println(params, errors.Wrap(err, "校验签名不匹配"))
	} else if params["result_code"] == "FAIL" {
		err = errors.New("退款失败")
		log.Println(params, errors.Wrap(err, params["err_code_des"]))
	}
	return
}
//应用场景
//当交易发生之后一段时间内，由于买家或者卖家的原因需要退款时，卖家可以通过退款接口将支付款退还给买家，微信支付将在收到退款请求并且验证成功之后，按照退款规则将支付款按原路退到买家帐号上。
//注意：
//1、交易时间超过一年的订单无法提交退款
//2、微信支付退款支持单笔交易分多次退款，多次退款需要提交原支付订单的商户订单号和设置不同的退款单号。申请退款总金额不能超过订单金额。 一笔退款失败后重新提交，请不要更换退款单号，请使用原商户退款单号
//
//3、请求频率限制：150qps，即每秒钟正常的申请退款请求次数不超过150次
//错误或无效请求频率限制：6qps，即每秒钟异常或错误的退款申请请求不超过6次
//4、每个支付订单的部分退款次数不能超过50次
//接口地址
//接口链接：https://api.mch.weixin.qq.com/secapi/pay/refund
//
//是否需要证书
//请求需要双向证书。 详见证书使用
//请求参数
//字段名	变量名	必填	类型	示例值	描述
//公众账号ID	appid	是	String(32)	wx8888888888888888	微信分配的公众账号ID（企业号corpid即为此appId）
//商户号	mch_id	是	String(32)	1900000109	微信支付分配的商户号
//随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
//签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法
//签名类型	sign_type	否	String(32)	HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
//微信订单号	transaction_id	二选一	String(28)	1217752501201407033233368018	微信生成的订单号，在支付通知中有返回
//商户订单号	out_trade_no	String(32)	1217752501201407033233368018	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
//商户退款单号	out_refund_no	是	String(64)	1217752501201407033233368018	商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
//订单金额	total_fee	是	Int	100	订单总金额，单位为分，只能为整数，详见支付金额
//退款金额	refund_fee	是	Int	100	退款总金额，订单总金额，单位为分，只能为整数，详见支付金额
//货币种类	refund_fee_type	否	String(8)	CNY	货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
//退款原因	refund_desc	否	String(80)	商品已售完	若商户传入，会在下发给用户的退款消息中体现退款原因
//退款资金来源	refund_account	否	String(30)	REFUND_SOURCE_RECHARGE_FUNDS
//仅针对老资金流商户使用
//REFUND_SOURCE_UNSETTLED_FUNDS---未结算资金退款（默认使用未结算资金退款）
//REFUND_SOURCE_RECHARGE_FUNDS---可用余额退款

//返回结果
//字段名	变量名	必填	类型	示例值	描述
//返回状态码	return_code	是	String(16)	SUCCESS	SUCCESS/FAIL
//返回信息	return_msg	否	String(128)	签名失败
//返回信息，如非空，为错误原因
//签名失败
//参数格式校验错误
//以下字段在return_code为SUCCESS的时候有返回
//
//字段名	变量名	必填	类型	示例值	描述
//业务结果	result_code	是	String(16)	SUCCESS
//SUCCESS/FAIL
//SUCCESS退款申请接收成功，结果通过退款查询接口查询
//FAIL 提交业务失败
//错误代码	err_code	否	String(32)	SYSTEMERROR	列表详见错误码列表
//错误代码描述	err_code_des	否	String(128)	系统超时	结果信息描述
//公众账号ID	appid	是	String(32)	wx8888888888888888	微信分配的公众账号ID
//商户号	mch_id	是	String(32)	1900000109	微信支付分配的商户号
//随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位
//签名	sign	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	签名，详见签名算法
//微信订单号	transaction_id	是	String(28)	4007752501201407033233368018	微信订单号
//商户订单号	out_trade_no	是	String(32)	33368018	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
//商户退款单号	out_refund_no	是	String(64)	121775250	商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
//微信退款单号	refund_id	是	String(32)	2007752501201407033233368018	微信退款单号
//退款金额	refund_fee	是	Int	100	退款总金额,单位为分,可以做部分退款
//应结退款金额	settlement_refund_fee	否	Int	100	去掉非充值代金券退款金额后的退款金额，退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额
//标价金额	total_fee	是	Int	100	订单总金额，单位为分，只能为整数，详见支付金额
//应结订单金额	settlement_total_fee	否	Int	100	去掉非充值代金券金额后的订单总金额，应结订单金额=订单金额-非充值代金券金额，应结订单金额<=订单金额。
//标价币种	fee_type	否	String(8)	CNY	订单金额货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
//现金支付金额	cash_fee	是	Int	100	现金支付金额，单位为分，只能为整数，详见支付金额
//现金支付币种	cash_fee_type	否	String(16)	CNY	货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
//现金退款金额	cash_refund_fee	否	Int	100	现金退款金额，单位为分，只能为整数，详见支付金额
//代金券类型	coupon_type_$n	否	String(8)	CASH
//CASH--充值代金券
//NO_CASH---非充值代金券
//订单使用代金券时有返回（取值：CASH、NO_CASH）。$n为下标,从0开始编号，举例：coupon_type_0
//代金券退款总金额	coupon_refund_fee	否	Int	100	代金券退款金额<=退款金额，退款金额-代金券或立减优惠退款金额为现金，说明详见代金券或立减优惠
//单个代金券退款金额	coupon_refund_fee_$n	否	Int	100	代金券退款金额<=退款金额，退款金额-代金券或立减优惠退款金额为现金，说明详见代金券或立减优惠
//退款代金券使用数量	coupon_refund_count	否	Int	1	退款代金券使用数量
//退款代金券ID	coupon_refund_id_$n	否	String(20)	10000 	退款代金券ID, $n为下标，从0开始编号

//错误码
//名称	描述	原因	解决方案
//SYSTEMERROR	接口返回错误	系统超时等	请不要更换商户退款单号，请使用相同参数再次调用API。
//BIZERR_NEED_RETRY	退款业务流程错误，需要商户触发重试来解决	并发情况下，业务被拒绝，商户重试即可解决	请不要更换商户退款单号，请使用相同参数再次调用API。
//TRADE_OVERDUE	订单已经超过退款期限	订单已经超过可退款的最大期限(支付后一年内可退款)	请选择其他方式自行退款
//ERROR	业务错误	申请退款业务发生错误	该错误都会返回具体的错误原因，请根据实际返回做相应处理。
//USER_ACCOUNT_ABNORMAL	退款请求失败	用户帐号注销	此状态代表退款申请失败，商户可自行处理退款。
//INVALID_REQ_TOO_MUCH	无效请求过多	连续错误请求数过多被系统短暂屏蔽	请检查业务是否正常，确认业务正常后请在1分钟后再来重试
//NOTENOUGH	余额不足	商户可用退款余额不足	此状态代表退款申请失败，商户可根据具体的错误提示做相应的处理。
//INVALID_TRANSACTIONID	无效transaction_id	请求参数未按指引进行填写	请求参数错误，检查原交易号是否存在或发起支付交易接口返回失败
//PARAM_ERROR	参数错误	请求参数未按指引进行填写	请求参数错误，请重新检查再调用退款申请
//APPID_NOT_EXIST	APPID不存在	参数中缺少APPID	请检查APPID是否正确
//MCHID_NOT_EXIST	MCHID不存在	参数中缺少MCHID	请检查MCHID是否正确
//REQUIRE_POST_METHOD	请使用post方法	未使用post传递参数 	请检查请求参数是否通过post方法提交
//SIGNERROR	签名错误	参数签名结果不正确	请检查签名参数和方法是否都符合签名算法要求
//XML_FORMAT_ERROR	XML格式错误	XML格式错误	请检查XML参数格式是否正确
//FREQUENCY_LIMITED	频率限制	2个月之前的订单申请退款有频率限制	该笔退款未受理，请降低频率后重试
