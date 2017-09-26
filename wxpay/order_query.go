package wxpay

import (
	"github.com/pkg/errors"
	"log"
	"mimi/djq/config"
	"mimi/djq/util"
	"strconv"
)

func OrderQuery(payOrderNumber string) (Params, error) {
	url := "https://api.mch.weixin.qq.com/pay/orderquery"

	appId := config.Get("wxpay_appid")  // 微信公众平台应用ID
	mchId := config.Get("wxpay_mch_id") // 微信支付商户平台商户号
	apiKey := config.Get("wxpay_key")   // 微信支付商户平台API密钥

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
	params.SetString("appid", c.AppId)  //公众账号ID	appid	是	String(32)	wxd678efh567hg6787	微信支付分配的公众账号ID（企业号corpid即为此appId）
	params.SetString("mch_id", c.MchId) //商户号	mch_id	是	String(32)	1230000109	微信支付分配的商户号
	//微信订单号	transaction_id	二选一	String(32)	1009660380201506130728806387	微信的订单号，建议优先使用
	params.SetString("out_trade_no", payOrderNumber) //商户订单号	out_trade_no	String(32)	20150806125346	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。 详见商户订单号
	params.SetString("nonce_str", util.BuildUUID())  //随机字符串	nonce_str	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	随机字符串，不长于32位。推荐随机数生成算法
	params.SetString("sign_type", "MD5")             //签名类型	sign_type	否	String(32)	HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	params.SetString("sign", c.Sign(params))         //签名	sign	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	通过签名算法计算得出的签名值，详见签名生成算法

	return c.Post(url, params, false)
}

func OrderQueryResult(payOrderNumber string) (tradeState string, totalFee int, err error) {
	params, err := OrderQuery(payOrderNumber)
	if err != nil {
		return
	}
	client := NewDefaultClient()
	if params["return_code"] == "FAIL" {
		err = errors.New("查询订单失败")
		log.Println(errors.Wrap(err, params["return_msg"]))
	} else if !client.CheckSign(params) {
		err = errors.New("查询订单失败")
		log.Println(params, errors.Wrap(err, "校验签名不匹配"))
	} else if params["result_code"] == "FAIL" {
		err = errors.New("查询订单失败")
		log.Println(params, errors.Wrap(err, params["err_code_des"]))
	} else {
		tradeState = params["trade_state"]
		if tradeState == "SUCCESS" {
			totalFee, err = strconv.Atoi(params["total_fee"])
			if err != nil {
				log.Println(params, err)
			}
		}
	}
	return
}

//应用场景
//该接口提供所有微信支付订单的查询，商户可以通过查询订单接口主动查询订单状态，完成下一步的业务逻辑。
//需要调用查询接口的情况：
//◆ 当商户后台、网络、服务器等出现异常，商户系统最终未接收到支付通知；
//◆ 调用支付接口后，返回系统错误或未知交易状态情况；
//◆ 调用刷卡支付API，返回USERPAYING的状态；
//◆ 调用关单或撤销接口API之前，需确认支付状态；
//接口链接
//https://api.mch.weixin.qq.com/pay/orderquery
//
//是否需要证书
//不需要
//
//请求参数
//
//字段名	变量名	必填	类型	示例值	描述
//公众账号ID	appid	是	String(32)	wxd678efh567hg6787	微信支付分配的公众账号ID（企业号corpid即为此appId）
//商户号	mch_id	是	String(32)	1230000109	微信支付分配的商户号
//微信订单号	transaction_id	二选一	String(32)	1009660380201506130728806387	微信的订单号，建议优先使用
//商户订单号	out_trade_no	String(32)	20150806125346	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。 详见商户订单号
//随机字符串	nonce_str	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	随机字符串，不长于32位。推荐随机数生成算法
//签名	sign	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	通过签名算法计算得出的签名值，详见签名生成算法
//签名类型	sign_type	否	String(32)	HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5

//返回结果
//字段名	变量名	必填	类型	示例值	描述
//返回状态码	return_code	是	String(16)	SUCCESS
//SUCCESS/FAIL
//此字段是通信标识，非交易标识，交易是否成功需要查看trade_state来判断
//返回信息	return_msg	否	String(128)	签名失败
//返回信息，如非空，为错误原因
//签名失败
//参数格式校验错误
//以下字段在return_code为SUCCESS的时候有返回
//
//字段名	变量名	必填	类型	示例值	描述
//公众账号ID	appid	是	String(32)	wxd678efh567hg6787	微信分配的公众账号ID
//商户号	mch_id	是	String(32)	1230000109	微信支付分配的商户号
//随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
//签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法
//业务结果	result_code	是	String(16)	SUCCESS	SUCCESS/FAIL
//错误代码	err_code	否	String(32)	SYSTEMERROR	错误码
//错误代码描述	err_code_des	否	String(128)	系统错误	结果信息描述
//以下字段在return_code 、result_code、trade_state都为SUCCESS时有返回 ，如trade_state不为 SUCCESS，则只返回out_trade_no（必传）和attach（选传）。
//
//字段名	变量名	必填	类型	示例值	描述
//设备号	device_info	否	String(32)	013467007045764	微信支付分配的终端设备号，
//用户标识	openid	是	String(128)	oUpF8uMuAJO_M2pxb1Q9zNjWeS6o	用户在商户appid下的唯一标识
//是否关注公众账号	is_subscribe	否	String(1)	Y	用户是否关注公众账号，Y-关注，N-未关注，仅在公众账号类型支付有效
//交易类型	trade_type	是	String(16)	JSAPI	调用接口提交的交易类型，取值如下：JSAPI，NATIVE，APP，MICROPAY，详细说明见参数规定
//交易状态	trade_state	是	String(32)	SUCCESS
//SUCCESS—支付成功
//REFUND—转入退款
//NOTPAY—未支付
//CLOSED—已关闭
//REVOKED—已撤销（刷卡支付）
//USERPAYING--用户支付中
//PAYERROR--支付失败(其他原因，如银行返回失败)
//支付状态机请见下单API页面
//付款银行	bank_type	是	String(16)	CMC	银行类型，采用字符串类型的银行标识
//标价金额	total_fee	是	Int	100	订单总金额，单位为分
//应结订单金额	settlement_total_fee	否	Int	100	当订单使用了免充值型优惠券后返回该参数，应结订单金额=订单金额-免充值优惠券金额。
//标价币种	fee_type	否	String(8)	CNY	货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
//现金支付金额	cash_fee	是	Int	100	现金支付金额订单现金支付金额，详见支付金额
//现金支付币种	cash_fee_type	否	String(16)	CNY	货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
//代金券金额	coupon_fee	否	Int	100	“代金券”金额<=订单金额，订单金额-“代金券”金额=现金支付金额，详见支付金额
//代金券使用数量	coupon_count	否	Int	1	代金券使用数量
//代金券类型	coupon_type_$n	否	String	CASH
//CASH--充值代金券
//NO_CASH---非充值优惠券
//开通免充值券功能，并且订单使用了优惠券后有返回（取值：CASH、NO_CASH）。$n为下标,从0开始编号，举例：coupon_type_$0
//代金券ID	coupon_id_$n	否	String(20)	10000 	代金券ID, $n为下标，从0开始编号
//单个代金券支付金额	coupon_fee_$n	否	Int	100	单个代金券支付金额, $n为下标，从0开始编号
//微信支付订单号	transaction_id	是	String(32)	1009660380201506130728806387	微信支付订单号
//商户订单号	out_trade_no	是	String(32)	20150806125346	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
//附加数据	attach	否	String(128)	深圳分店	附加数据，原样返回
//支付完成时间	time_end	是	String(14)	20141030133525	订单支付时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
//交易状态描述	trade_state_desc	是	String(256)	支付失败，请重新下单支付	对当前查询订单状态的描述和下一步操作的指引

//错误码
//名称	描述	原因	解决方案
//ORDERNOTEXIST	此交易订单号不存在	查询系统中不存在此交易订单号	该API只能查提交支付交易返回成功的订单，请商户检查需要查询的订单号是否正确
//SYSTEMERROR	系统错误	后台系统返回错误	系统异常，请再调用发起查询
