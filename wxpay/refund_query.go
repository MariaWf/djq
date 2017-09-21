package wxpay

import (
	"strconv"
	"mimi/djq/config"
	"mimi/djq/util"
	"log"
	"github.com/pkg/errors"
)

func RefundQuery(RefundOrderNumber string) (Params, error) {
	url := "https://api.mch.weixin.qq.com/pay/orderquery"

	appId := config.Get("wxpay_appid") // 微信公众平台应用ID
	mchId := config.Get("wxpay_mch_id") // 微信支付商户平台商户号
	apiKey := config.Get("wxpay_key") // 微信支付商户平台API密钥

	var err error
	if config.Get("running_state") == "test" {
		url = "https://api.mch.weixin.qq.com/sandboxnew/pay/orderquery"
		apiKey, err = GetSignKey()
		if err != nil {
			return nil, err
		}
	}

	c := NewClient(appId, mchId, apiKey)

	params := make(Params)

	//字段名	变量名	必填	类型	示例值	描述
	params.SetString("appid", c.AppId)//公众账号ID	appid	是	String(32)	wxd678efh567hg6787	微信支付分配的公众账号ID（企业号corpid即为此appId）
	params.SetString("mch_id", c.MchId)//商户号	mch_id	是	String(32)	1230000109	微信支付分配的商户号
	//微信退款号	transaction_id	四选一	String(32)	1217752501201407033233368018	微信退款号查询的优先级是： refund_id > out_refund_no > transaction_id > out_trade_no
	//商户退款号	out_trade_no	String(32)	1217752501201407033233368018	商户系统内部退款号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
	//商户退款单号	out_refund_no	String(64)	1217752501201407033233368018	商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	//微信退款单号	refund_id	String(32)	1217752501201407033233368018
	params.SetString("out_refund_no", RefundOrderNumber)
	params.SetString("nonce_str", util.BuildUUID())//随机字符串	nonce_str	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	随机字符串，不长于32位。推荐随机数生成算法
	params.SetString("sign_type", "MD5")//签名类型	sign_type	否	String(32)	HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	params.SetString("sign", c.Sign(params))//签名	sign	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	通过签名算法计算得出的签名值，详见签名生成算法

	return c.Post(url, params, false)
}

func RefundQueryResult(RefundOrderNumber string) (refundState string, refundFee int, err error) {
	params, err := RefundQuery(RefundOrderNumber)
	if err != nil {
		return
	}
	client := NewDefaultClient()
	if params["return_code"] == "FAIL" {
		err = errors.New("查询退款失败")
		log.Println(errors.Wrap(err, params["return_msg"]))
	} else if !client.CheckSign(params) {
		err = errors.New("查询退款失败")
		log.Println(params, errors.Wrap(err, "校验签名不匹配"))
	} else if params["result_code"] == "FAIL" {
		err = errors.New("查询退款失败")
		log.Println(params, errors.Wrap(err, params["err_code_des"]))
	} else {
		refundState = params["refund_status_0"]
		if refundState == "SUCCESS" {
			refundFee, err = strconv.Atoi(params["refund_fee_0"])
			if err != nil {
				log.Println(params, err)
			}
		}
	}
	return
}



//应用场景
//提交退款申请后，通过调用该接口查询退款状态。退款有一定延时，用零钱支付的退款20分钟内到账，银行卡支付的退款3个工作日后重新查询退款状态。
//
//注意：如果单个支付退款部分退款次数超过20次请使用退款单号查询
//退款状态机
//退款状态变化如下：
//
//接口地址
//接口链接：https://api.mch.weixin.qq.com/pay/refundquery
//
//是否需要证书
//不需要。
//
//请求参数
//字段名	变量名	必填	类型	示例值	描述
//公众账号ID	appid	是	String(32)	wx8888888888888888	微信支付分配的公众账号ID（企业号corpid即为此appId）
//商户号	mch_id	是	String(32)	1900000109	微信支付分配的商户号
//随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
//签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法
//签名类型	sign_type	否	String(32)	HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
//微信退款号	transaction_id	四选一	String(32)	1217752501201407033233368018	微信退款号查询的优先级是： refund_id > out_refund_no > transaction_id > out_trade_no
//商户退款号	out_trade_no	String(32)	1217752501201407033233368018	商户系统内部退款号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
//商户退款单号	out_refund_no	String(64)	1217752501201407033233368018	商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
//微信退款单号	refund_id	String(32)	1217752501201407033233368018
//微信生成的退款单号，在申请退款接口有返回

//返回数据
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
//FAIL
//错误码	err_code	是	String(32)	SYSTEMERROR	错误码详见第6节
//错误描述	err_code_des	是	String(32)	系统错误	结果信息描述
//公众账号ID	appid	是	String(32)	wx8888888888888888	微信分配的公众账号ID（企业号corpid即为此appId）
//商户号	mch_id	是	String(32)	1900000109	微信支付分配的商户号
//随机字符串	nonce_str	是	String(28)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位
//签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名算法
//微信退款号	transaction_id	是	String(32)	1217752501201407033233368018	微信退款号
//商户退款号	out_trade_no	是	String(32)	1217752501201407033233368018	商户系统内部退款号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
//退款金额	total_fee	是	Int	100	退款总金额，单位为分，只能为整数，详见支付金额
//应结退款金额	settlement_total_fee	否	Int	100	当退款使用了免充值型优惠券后返回该参数，应结退款金额=退款金额-免充值优惠券金额。
//货币种类	fee_type	否	String(8)	CNY	退款金额货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
//现金支付金额	cash_fee	是	Int	100	现金支付金额，单位为分，只能为整数，详见支付金额
//退款笔数	refund_count	是	Int	1	退款记录数
//商户退款单号	out_refund_no_$n	是	String(32)	1217752501201407033233368018	商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
//微信退款单号	refund_id_$n	是	String(32)	1217752501201407033233368018	微信退款单号
//退款渠道	refund_channel_$n	否	String(16)	ORIGINAL
//ORIGINAL—原路退款
//BALANCE—退回到余额
//OTHER_BALANCE—原账户异常退到其他余额账户
//OTHER_BANKCARD—原银行卡异常退到其他银行卡
//申请退款金额	refund_fee_$n	是	Int	100	退款总金额,单位为分,可以做部分退款
//退款金额	settlement_refund_fee_$n	否	Int	100	退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额
//代金券类型	coupon_type_$n	否	String	CASH
//CASH--充值代金券
//NO_CASH---非充值优惠券
//开通免充值券功能，并且退款使用了优惠券后有返回（取值：CASH、NO_CASH）。$n为下标,从0开始编号，举例：coupon_type_$0
//总代金券退款金额	coupon_refund_fee_$n	否	Int	100	代金券退款金额<=退款金额，退款金额-代金券或立减优惠退款金额为现金，说明详见代金券或立减优惠
//退款代金券使用数量	coupon_refund_count_$n	否	Int	1	退款代金券使用数量 ,$n为下标,从0开始编号
//退款代金券ID	coupon_refund_id_$n_$m	否	String(20)	10000 	退款代金券ID, $n为下标，$m为下标，从0开始编号
//单个代金券退款金额	coupon_refund_fee_$n_$m	否	Int	100	单个退款代金券支付金额, $n为下标，$m为下标，从0开始编号
//退款状态	refund_status_$n	是	String(16)	SUCCESS
//退款状态：
//SUCCESS—退款成功
//REFUNDCLOSE—退款关闭。
//PROCESSING—退款处理中
//CHANGE—退款异常，退款到银行发现用户的卡作废或者冻结了，导致原路退款银行卡失败，可前往商户平台（pay.weixin.qq.com）-交易中心，手动处理此笔退款。$n为下标，从0开始编号。
//退款资金来源	refund_account_$n	否	String(30)	REFUND_SOURCE_RECHARGE_FUNDS
//REFUND_SOURCE_RECHARGE_FUNDS---可用余额退款/基本账户
//REFUND_SOURCE_UNSETTLED_FUNDS---未结算资金退款
//$n为下标，从0开始编号。
//退款入账账户	refund_recv_accout_$n	是	String(64)	招商银行信用卡0403	取当前退款单的退款入账方
//1）退回银行卡：
//{银行名称}{卡类型}{卡尾号}
//2）退回支付用户零钱:
//支付用户零钱
//3）退还商户:
//商户基本账户
//商户结算银行账户
//4）退回支付用户零钱通:
//支付用户零钱通
//退款成功时间	refund_success_time_$n	否	String(20)	2016-07-25 15:26:26	退款成功时间，当退款状态为退款成功时有返回。$n为下标，从0开始编号。
//
//错误码
//名称	描述	原因	解决方案
//SYSTEMERROR	接口返回错误	系统超时	请尝试再次掉调用API。
//REFUNDNOTEXIST	退款退款查询失败	退款号错误或退款状态不正确	请检查退款号是否有误以及退款状态是否正确，如：未支付、已支付未退款
//INVALID_TRANSACTIONID	无效transaction_id	请求参数未按指引进行填写	请求参数错误，检查原交易号是否存在或发起支付交易接口返回失败
//PARAM_ERROR	参数错误	请求参数未按指引进行填写	请求参数错误，请检查参数再调用退款申请
//APPID_NOT_EXIST	APPID不存在	参数中缺少APPID	请检查APPID是否正确
//MCHID_NOT_EXIST	MCHID不存在	参数中缺少MCHID	请检查MCHID是否正确
//REQUIRE_POST_METHOD	请使用post方法	未使用post传递参数 	请检查请求参数是否通过post方法提交
//SIGNERROR	签名错误	参数签名结果不正确	请检查签名参数和方法是否都符合签名算法要求
//XML_FORMAT_ERROR	XML格式错误	XML格式错误	请检查XML参数格式是否正确
