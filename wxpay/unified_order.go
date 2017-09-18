package wxpay

import (
	"mimi/djq/config"
	"mimi/djq/util"
	"github.com/pkg/errors"
	"strconv"
)

func UnifiedOrder(payOrderNumber string, totalFee int, clientIp string,openId string) (Params, error) {
	var err error
	// 查询企业付款接口请求URL
	url := "https://api.mch.weixin.qq.com/pay/unifiedorder"

	appId := config.Get("wxpay_appid") // 微信公众平台应用ID
	mchId := config.Get("wxpay_mch_id") // 微信支付商户平台商户号
	apiKey := config.Get("wxpay_key") // 微信支付商户平台API密钥

	if config.Get("running_state") == "test" {
		url = "https://api.mch.weixin.qq.com/sandboxnew/pay/unifiedorder"
		apiKey, err = GetSignKey()
		if err != nil {
			return nil, errors.Wrap(err,"获取测试API_KEY失败")
		}
	}

	c := NewClient(appId, mchId, apiKey)

	params := make(Params)

	totalFee = 179

	params.SetString("appid", c.AppId)//公众账号ID	appid	是	String(32)	wxd678efh567hg6787	微信支付分配的公众账号ID（企业号corpid即为此appId）
	params.SetString("mch_id", c.MchId)//商户号	mch_id	是	String(32)	1230000109	微信支付分配的商户号
	params.SetString("device_info", "WEB") //设备号	device_info	否	String(32)	013467007045764	自定义参数，可以为终端设备号(门店号或收银设备ID)，PC网页或公众号内支付可以传"WEB"
	params.SetString("nonce_str",  util.BuildUUID()) //随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，长度要求在32位以内。推荐随机数生成算法
	params.SetString("sign_type", "MD5") //签名类型	sign_type	否	String(32)	MD5	签名类型，默认为MD5，支持HMAC-SHA256和MD5。
	params.SetString("body", "摩设-代金券") //商品描述	body	是	String(128)	腾讯充值中心-QQ会员充值
	//商品简单描述，该字段请按照规范传递，具体请见参数规定
	params.SetString("detail", "") //商品详情	detail	否	String(6000)	 	商品详细描述，对于使用单品优惠的商户，改字段必须按照规范上传，详见“单品优惠参数说明”
	params.SetString("attach", "") //附加数据	attach	否	String(127)	深圳分店	附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用。
	params.SetString("out_trade_no", payOrderNumber) //商户订单号	out_trade_no	是	String(32)	20150806125346	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。详见商户订单号
	params.SetString("fee_type", "CNY") //标价币种	fee_type	否	String(16)	CNY	符合ISO 4217标准的三位字母代码，默认人民币：CNY，详细列表请参见货币类型
	params.SetString("total_fee", strconv.Itoa(totalFee)) //标价金额	total_fee	是	Int	88	订单总金额，单位为分，详见支付金额
	params.SetString("spbill_create_ip", clientIp) //终端IP	spbill_create_ip	是	String(16)	123.12.12.123	APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP。
	params.SetString("time_start", "") //交易起始时间	time_start	否	String(14)	20091225091010	订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	params.SetString("time_expire", "") //交易结束时间	time_expire	否	String(14)	20091227091010
	//订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。其他详见时间规则
	//注意：最短失效时间间隔必须大于5分钟
	params.SetString("goods_tag", "") //订单优惠标记	goods_tag	否	String(32)	WXG	订单优惠标记，使用代金券或立减优惠功能时需要的参数，说明详见代金券或立减优惠
	params.SetString("notify_url", util.PathAppend(config.Get("server_root_url"), "open/wxpay/notify4UnifiedOrder")) //通知地址	notify_url	是	String(256)	http://www.weixin.qq.com/wxpay/pay.php	异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。
	params.SetString("trade_type", "JSAPI") //交易类型	trade_type	是	String(16)	JSAPI	取值如下：JSAPI，NATIVE，APP等，说明详见参数规定
	params.SetString("product_id", "") //商品ID	product_id	否	String(32)	12235413214070356458058	trade_type=NATIVE时（即扫码支付），此参数必传。此参数为二维码中包含的商品ID，商户自行定义。
	params.SetString("limit_pay", "") //指定支付方式	limit_pay	否	String(32)	no_credit	上传此参数no_credit--可限制用户不能使用信用卡支付
	params.SetString("openid", openId) //用户标识	openid	否	String(128)	oUpF8uMuAJO_M2pxb1Q9zNjWeS6o	trade_type=JSAPI时（即公众号支付），此参数必传，此参数为微信用户在商户对应appid下的唯一标识。openid如何获取，可参考【获取openid】。企业号请使用【企业号OAuth2.0接口】获取企业号内成员userid，再调用【企业号userid转openid接口】进行转换
	params.SetString("scene_info", "") //+场景信息	scene_info	否	String(256)
	//{"store_info" : {
	//"id": "SZTX001",
	//"name": "腾大餐厅",
	//"area_code": "440305",
	//"address": "科技园中一路腾讯大厦" }}
	//该字段用于上报场景信息，目前支持上报实际门店信息。该字段为JSON对象数据，对象格式为{"store_info":{"id": "门店ID","name": "名称","area_code": "编码","address": "地址" }} ，字段详细说明请点击行前的+展开

	params.SetString("sign", c.Sign(params)) //签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	通过签名算法计算得出的签名值，详见签名生成算法


	// 发送查询企业付款请求
	return c.Post(url, params, false)
}

//应用场景
//除被扫支付场景以外，商户系统先调用该接口在微信支付服务后台生成预支付交易单，返回正确的预支付交易回话标识后再按扫码、JSAPI、APP等不同场景生成交易串调起支付。

//接口链接
//URL地址：https://api.mch.weixin.qq.com/pay/unifiedorder

//是否需要证书
//否

//请求参数
//公众账号ID	appid	是	String(32)	wxd678efh567hg6787	微信支付分配的公众账号ID（企业号corpid即为此appId）
//商户号	mch_id	是	String(32)	1230000109	微信支付分配的商户号
//设备号	device_info	否	String(32)	013467007045764	自定义参数，可以为终端设备号(门店号或收银设备ID)，PC网页或公众号内支付可以传"WEB"
//随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，长度要求在32位以内。推荐随机数生成算法
//签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	通过签名算法计算得出的签名值，详见签名生成算法
//签名类型	sign_type	否	String(32)	MD5	签名类型，默认为MD5，支持HMAC-SHA256和MD5。
//商品描述	body	是	String(128)	腾讯充值中心-QQ会员充值
//商品简单描述，该字段请按照规范传递，具体请见参数规定
//商品详情	detail	否	String(6000)	 	商品详细描述，对于使用单品优惠的商户，改字段必须按照规范上传，详见“单品优惠参数说明”
//附加数据	attach	否	String(127)	深圳分店	附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用。
//商户订单号	out_trade_no	是	String(32)	20150806125346	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。详见商户订单号
//标价币种	fee_type	否	String(16)	CNY	符合ISO 4217标准的三位字母代码，默认人民币：CNY，详细列表请参见货币类型
//标价金额	total_fee	是	Int	88	订单总金额，单位为分，详见支付金额
//终端IP	spbill_create_ip	是	String(16)	123.12.12.123	APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP。
//交易起始时间	time_start	否	String(14)	20091225091010	订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
//交易结束时间	time_expire	否	String(14)	20091227091010
//订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。其他详见时间规则
//注意：最短失效时间间隔必须大于5分钟
//订单优惠标记	goods_tag	否	String(32)	WXG	订单优惠标记，使用代金券或立减优惠功能时需要的参数，说明详见代金券或立减优惠
//通知地址	notify_url	是	String(256)	http://www.weixin.qq.com/wxpay/pay.php	异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。
//交易类型	trade_type	是	String(16)	JSAPI	取值如下：JSAPI，NATIVE，APP等，说明详见参数规定
//商品ID	product_id	否	String(32)	12235413214070356458058	trade_type=NATIVE时（即扫码支付），此参数必传。此参数为二维码中包含的商品ID，商户自行定义。
//指定支付方式	limit_pay	否	String(32)	no_credit	上传此参数no_credit--可限制用户不能使用信用卡支付
//用户标识	openid	否	String(128)	oUpF8uMuAJO_M2pxb1Q9zNjWeS6o	trade_type=JSAPI时（即公众号支付），此参数必传，此参数为微信用户在商户对应appid下的唯一标识。openid如何获取，可参考【获取openid】。企业号请使用【企业号OAuth2.0接口】获取企业号内成员userid，再调用【企业号userid转openid接口】进行转换
//+场景信息	scene_info	否	String(256)
//{"store_info" : {
//"id": "SZTX001",
//"name": "腾大餐厅",
//"area_code": "440305",
//"address": "科技园中一路腾讯大厦" }}
//该字段用于上报场景信息，目前支持上报实际门店信息。该字段为JSON对象数据，对象格式为{"store_info":{"id": "门店ID","name": "名称","area_code": "编码","address": "地址" }} ，字段详细说明请点击行前的+展开


//返回结果
//字段名	变量名	必填	类型	示例值	描述
//返回状态码	return_code	是	String(16)	SUCCESS
//SUCCESS/FAIL
//此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
//返回信息	return_msg	否	String(128)	签名失败
//返回信息，如非空，为错误原因
//签名失败
//参数格式校验错误
//以下字段在return_code为SUCCESS的时候有返回
//
//字段名	变量名	必填	类型	示例值	描述
//公众账号ID	appid	是	String(32)	wx8888888888888888	调用接口提交的公众账号ID
//商户号	mch_id	是	String(32)	1900000109	调用接口提交的商户号
//设备号	device_info	否	String(32)	013467007045764	自定义参数，可以为请求支付的终端设备号等
//随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	微信返回的随机字符串
//签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	微信返回的签名值，详见签名算法
//业务结果	result_code	是	String(16)	SUCCESS	SUCCESS/FAIL
//错误代码	err_code	否	String(32)	SYSTEMERROR	详细参见下文错误列表
//错误代码描述	err_code_des	否	String(128)	系统错误	错误信息描述
//以下字段在return_code 和result_code都为SUCCESS的时候有返回
//
//字段名	变量名	必填	类型	示例值	描述
//交易类型	trade_type	是	String(16)	JSAPI	交易类型，取值为：JSAPI，NATIVE，APP等，说明详见参数规定
//预支付交易会话标识	prepay_id	是	String(64)	wx201410272009395522657a690389285100	微信生成的预支付会话标识，用于后续接口调用中使用，该值有效期为2小时
//二维码链接	code_url	否	String(64)	URl：weixin：//wxpay/s/An4baqw	trade_type为NATIVE时有返回，用于生成二维码，展示给用户进行扫码支付

//错误码
//名称	描述	原因	解决方案
//NOAUTH	商户无此接口权限	商户未开通此接口权限	请商户前往申请此接口权限
//NOTENOUGH	余额不足	用户帐号余额不足	用户帐号余额不足，请用户充值或更换支付卡后再支付
//ORDERPAID	商户订单已支付	商户订单已支付，无需重复操作	商户订单已支付，无需更多操作
//ORDERCLOSED	订单已关闭	当前订单已关闭，无法支付	当前订单已关闭，请重新下单
//SYSTEMERROR	系统错误	系统超时	系统异常，请用相同参数重新调用
//APPID_NOT_EXIST	APPID不存在	参数中缺少APPID	请检查APPID是否正确
//MCHID_NOT_EXIST	MCHID不存在	参数中缺少MCHID	请检查MCHID是否正确
//APPID_MCHID_NOT_MATCH	appid和mch_id不匹配	appid和mch_id不匹配	请确认appid和mch_id是否匹配
//LACK_PARAMS	缺少参数	缺少必要的请求参数	请检查参数是否齐全
//OUT_TRADE_NO_USED	商户订单号重复	同一笔交易不能多次提交	请核实商户订单号是否重复提交
//SIGNERROR	签名错误	参数签名结果不正确	请检查签名参数和方法是否都符合签名算法要求
//XML_FORMAT_ERROR	XML格式错误	XML格式错误	请检查XML参数格式是否正确
//REQUIRE_POST_METHOD	请使用post方法	未使用post传递参数 	请检查请求参数是否通过post方法提交
//POST_DATA_EMPTY	post数据为空	post数据不能为空	请检查post数据是否为空
//NOT_UTF8	编码格式错误	未使用指定编码格式	请使用UTF-8编码格式