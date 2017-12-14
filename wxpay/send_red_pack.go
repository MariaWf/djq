package wxpay

import (
	"github.com/pkg/errors"
	"log"
	"mimi/djq/config"
	"mimi/djq/util"
	"strconv"
)

func SendRedPack(openId string, totalAmount int) (Params, error) {
	var err error
	url := "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack"

	appId := config.Get("wxpay_appid")  // 微信公众平台应用ID
	mchId := config.Get("wxpay_mch_id") // 微信支付商户平台商户号
	apiKey := config.Get("wxpay_key")   // 微信支付商户平台API密钥

	if config.Get("running_state") == "test" {
		url = "https://api.mch.weixin.qq.com/sandboxnew/mmpaymkttransfers/sendredpack"
		apiKey, err = GetSignKey()
		if err != nil {
			return nil, errors.Wrap(err, "获取测试API_KEY失败")
		}
	}

	c := NewClient(appId, mchId, apiKey)

	// 微信支付商户平台证书路径
	certFile := config.Get("wxpay_cert_file")
	keyFile := config.Get("wxpay_key_file")
	rootcaFile := config.Get("wxpay_rootca_file")

	// 附着商户证书
	err = c.WithCert(certFile, keyFile, rootcaFile)
	if err != nil {
		log.Println(err)
		return nil, errors.Wrap(err, "获取商户证书失败")
	}

	sendName := "摩设共享设计平台"
	wishing := "恭喜发财"
	actName := "恭喜发财"
	remark := "恭喜发财"
	clientIp := config.Get("external_ip")
	params := make(Params)

	params.SetString("nonce_str", util.BuildUUID())       //随机字符串	nonce_str	是	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	String(32)	随机字符串，不长于32位
	params.SetString("mch_billno", util.BuildUUID()[:28]) //商户订单号	mch_billno	是	10000098201411111234567890	String(28)
	//商户订单号（每个订单号必须唯一。取值范围：0~9，a~z，A~Z）
	//接口根据商户订单号支持重入，如出现超时可再调用。
	params.SetString("mch_id", c.MchId)     //商户号	mch_id	是	10000098	String(32)	微信支付分配的商户号
	params.SetString("wxappid", c.AppId)    //公众账号appid	wxappid	是	wx8888888888888888	String(32)	微信分配的公众账号ID（企业号corpid即为此appId）。接口传入的所有appid应该为公众号的appid（在mp.weixin.qq.com申请的），不能为APP的appid（在open.weixin.qq.com申请的）。
	params.SetString("send_name", sendName) //商户名称	send_name	是	天虹百货	String(32)	红包发送者名称
	params.SetString("re_openid", openId)   //用户openid	re_openid	是	oxTWIuGaIt6gTKsQRLau2M0yL16E	String(32)
	//接受红包的用户
	//用户在wxappid下的openid
	params.SetString("total_amount", strconv.Itoa(totalAmount)) //付款金额	total_amount	是	1000	int	付款金额，单位分
	params.SetString("total_num", strconv.Itoa(1))              //红包发放总人数	total_num	是	1	int
	//红包发放总人数
	//total_num=1
	params.SetString("wishing", wishing)      //红包祝福语	wishing	是	感谢您参加猜灯谜活动，祝您元宵节快乐！	String(128)	红包祝福语
	params.SetString("client_ip", clientIp)   //Ip地址	client_ip	是	192.168.0.1	String(15)	调用接口的机器Ip地址
	params.SetString("act_name", actName)     //活动名称	act_name	是	猜灯谜抢红包活动	String(32)	活动名称
	params.SetString("remark", remark)        //备注	remark	是	猜越多得越多，快来抢！	String(256)	备注信息
	params.SetString("scene_id", "PRODUCT_4") //场景id	scene_id	否	PRODUCT_8	String(32)
	//发放红包使用场景，红包金额大于200时必传
	//PRODUCT_1:商品促销
	//PRODUCT_2:抽奖
	//PRODUCT_3:虚拟物品兑奖
	//PRODUCT_4:企业内部福利
	//PRODUCT_5:渠道分润
	//PRODUCT_6:保险回馈
	//PRODUCT_7:彩票派奖
	//PRODUCT_8:税务刮奖
	//活动信息	risk_info	否	posttime%3d123123412%26clientversion%3d234134%26mobile%3d122344545%26deviceid%3dIOS	String(128)
	//posttime:用户操作的时间戳
	//mobile:业务系统账号的手机号，国家代码-手机号。不需要+号
	//deviceid :mac 地址或者设备唯一标识
	//clientversion :用户操作的客户端版本
	//把值为非空的信息用key=value进行拼接，再进行urlencode
	//urlencode(posttime=xx& mobile =xx&deviceid=xx)
	//资金授权商户号	consume_mch_id	否	1222000096	String(32)
	//资金授权商户号
	//服务商替特约商户发放时使用
	params.SetString("sign", c.Sign(params)) //签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法

	return c.Post(url, params, true)

}

func SendRedPackResult(openId string, totalAmount int) (err error) {
	params, err := SendRedPack(openId, totalAmount)
	if err != nil {
		return
	}
	client := NewDefaultClient()
	if params["return_code"] == "FAIL" {
		err = errors.New("发红包失败")
		log.Println(errors.Wrap(err, params["return_msg"]))
	} else if params["sign"] !="" && !client.CheckSign(params) {
		err = errors.New("发红包失败")
		log.Println(params, errors.Wrap(err, "校验签名不匹配"))
	} else if params["result_code"] == "FAIL" {
		err = errors.New("发红包失败")
		log.Println(params, errors.Wrap(err, params["err_code_des"]))
	}
	return
}

//发放规则
//1.发送频率限制------默认1800/min
//2.发送个数上限------按照默认1800/min算
//3.金额上限------根据传入场景id不同默认上限不同，可以在商户平台产品设置进行设置和申请，最大不大于4999元/个
//4.其他的“量”上的限制还有哪些？------用户当天的领取上限次数,默认是10
//5.如果量上满足不了我们的需求，如何提高各个上限？------金额上限和用户当天领取次数上限可以在商户平台进行设置
//注意-红包金额大于200时，请求参数scene_id必传，参数说明见下文。
//注意2-根据监管要求，新申请商户号使用现金红包需要满足两个条件：1、入驻时间超过90天 2、连续正常交易30天。
//接口调用请求说明
//请求Url	https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack
//是否需要证书	是（证书及使用说明详见商户证书）
//请求方式	POST
//请求参数
//字段名	字段	必填	示例值	类型	说明
//随机字符串	nonce_str	是	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	String(32)	随机字符串，不长于32位
//签名	sign	是	C380BEC2BFD727A4B6845133519F3AD6	String(32)	详见签名生成算法
//商户订单号	mch_billno	是	10000098201411111234567890	String(28)
//商户订单号（每个订单号必须唯一。取值范围：0~9，a~z，A~Z）
//接口根据商户订单号支持重入，如出现超时可再调用。
//商户号	mch_id	是	10000098	String(32)	微信支付分配的商户号
//公众账号appid	wxappid	是	wx8888888888888888	String(32)	微信分配的公众账号ID（企业号corpid即为此appId）。接口传入的所有appid应该为公众号的appid（在mp.weixin.qq.com申请的），不能为APP的appid（在open.weixin.qq.com申请的）。
//商户名称	send_name	是	天虹百货	String(32)	红包发送者名称
//用户openid	re_openid	是	oxTWIuGaIt6gTKsQRLau2M0yL16E	String(32)
//接受红包的用户
//用户在wxappid下的openid
//付款金额	total_amount	是	1000	int	付款金额，单位分
//红包发放总人数	total_num	是	1	int
//红包发放总人数
//total_num=1
//红包祝福语	wishing	是	感谢您参加猜灯谜活动，祝您元宵节快乐！	String(128)	红包祝福语
//Ip地址	client_ip	是	192.168.0.1	String(15)	调用接口的机器Ip地址
//活动名称	act_name	是	猜灯谜抢红包活动	String(32)	活动名称
//备注	remark	是	猜越多得越多，快来抢！	String(256)	备注信息
//场景id	scene_id	否	PRODUCT_8	String(32)
//发放红包使用场景，红包金额大于200时必传
//PRODUCT_1:商品促销
//PRODUCT_2:抽奖
//PRODUCT_3:虚拟物品兑奖
//PRODUCT_4:企业内部福利
//PRODUCT_5:渠道分润
//PRODUCT_6:保险回馈
//PRODUCT_7:彩票派奖
//PRODUCT_8:税务刮奖
//活动信息	risk_info	否	posttime%3d123123412%26clientversion%3d234134%26mobile%3d122344545%26deviceid%3dIOS	String(128)
//posttime:用户操作的时间戳
//mobile:业务系统账号的手机号，国家代码-手机号。不需要+号
//deviceid :mac 地址或者设备唯一标识
//clientversion :用户操作的客户端版本
//把值为非空的信息用key=value进行拼接，再进行urlencode
//urlencode(posttime=xx& mobile =xx&deviceid=xx)
//资金授权商户号	consume_mch_id	否	1222000096	String(32)
//资金授权商户号
//服务商替特约商户发放时使用
//数据示例：
//<xml>
//<sign><![CDATA[E1EE61A91C8E90F299DE6AE075D60A2D]]></sign>
//<mch_billno><![CDATA[0010010404201411170000046545]]></mch_billno>
//<mch_id><![CDATA[888]]></mch_id>
//<wxappid><![CDATA[wxcbda96de0b165486]]></wxappid>
//<send_name><![CDATA[send_name]]></send_name>
//<re_openid><![CDATA[onqOjjmM1tad-3ROpncN-yUfa6uI]]></re_openid>
//<total_amount><![CDATA[200]]></total_amount>
//<total_num><![CDATA[1]]></total_num>
//<wishing><![CDATA[恭喜发财]]></wishing>
//<client_ip><![CDATA[127.0.0.1]]></client_ip>
//<act_name><![CDATA[新年红包]]></act_name>
//<remark><![CDATA[新年红包]]></remark>
//<scene_id><![CDATA[PRODUCT_2]]></scene_id>
//<consume_mch_id><![CDATA[10000097]]></consume_mch_id>
//<nonce_str><![CDATA[50780e0cca98c8c8e814883e5caa672e]]></nonce_str>
//<risk_info>posttime%3d123123412%26clientversion%3d234134%26mobile%3d122344545%26deviceid%3dIOS</risk_info>
//</xml>
//返回参数
//字段名	变量名	必填	示例值	类型	说明
//返回状态码	return_code	是	SUCCESS	String(16)
//SUCCESS/FAIL
//此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
//返回信息	return_msg	否	签名失败	String(128)
//返回信息，如非空，为错误原因
//签名失败
//参数格式校验错误
//以下字段在return_code为SUCCESS的时候有返回
//签名	sign	是	C380BEC2BFD727A4B6845133519F3AD6	String(32)	生成签名方式详见签名生成算法
//业务结果	result_code	是	SUCCESS	String(16)	SUCCESS/FAIL
//错误代码	err_code	否	SYSTEMERROR	String(32)	错误码信息
//错误代码描述	err_code_des	否	系统错误	String(128)	结果信息描述
//以下字段在return_code和result_code都为SUCCESS的时候有返回
//商户订单号	mch_billno	是	10000098201411111234567890	String(28)
//商户订单号（每个订单号必须唯一）
//组成：mch_id+yyyymmdd+10位一天内不能重复的数字
//商户号	mch_id	是	10000098	String(32)	微信支付分配的商户号
//公众账号appid	wxappid	是	wx8888888888888888	String(32)	商户appid，接口传入的所有appid应该为公众号的appid（在mp.weixin.qq.com申请的），不能为APP的appid（在open.weixin.qq.com申请的）。
//用户openid	re_openid	是	oxTWIuGaIt6gTKsQRLau2M0yL16E	String(32)
//接受收红包的用户
//用户在wxappid下的openid
//付款金额	total_amount	是	1000	int	付款金额，单位分
//微信单号	send_listid	是	100000000020150520314766074200	String(32)	红包订单的微信单号
//成功示例：
//<xml>
//<return_code><![CDATA[SUCCESS]]></return_code>
//<return_msg><![CDATA[发放成功.]]></return_msg>
//<result_code><![CDATA[SUCCESS]]></result_code>
//<err_code><![CDATA[0]]></err_code>
//<err_code_des><![CDATA[发放成功.]]></err_code_des>
//<mch_billno><![CDATA[0010010404201411170000046545]]></mch_billno>
//<mch_id>10010404</mch_id>
//<wxappid><![CDATA[wx6fa7e3bab7e15415]]></wxappid>
//<re_openid><![CDATA[onqOjjmM1tad-3ROpncN-yUfa6uI]]></re_openid>
//<total_amount>1</total_amount>
//</xml>
//失败示例：
//<xml>
//<return_code><![CDATA[FAIL]]></return_code>
//<return_msg><![CDATA[系统繁忙,请稍后再试.]]></return_msg>
//<result_code><![CDATA[FAIL]]></result_code>
//<err_code><![CDATA[268458547]]></err_code>
//<err_code_des><![CDATA[系统繁忙,请稍后再试.]]></err_code_des>
//<mch_billno><![CDATA[0010010404201411170000046542]]></mch_billno>
//<mch_id>10010404</mch_id>
//<wxappid><![CDATA[wx6fa7e3bab7e15415]]></wxappid>
//<re_openid><![CDATA[onqOjjmM1tad-3ROpncN-yUfa6uI]]></re_openid>
//<total_amount>1</total_amount>
//</xml>
//
//错误码
//错误码	错误描述	原因	解决方式
//NO_AUTH	发放失败，此请求可能存在风险，已被微信拦截	用户账号异常，被拦截	请提醒用户检查自身帐号是否异常。使用常用的活跃的微信号可避免这种情况。
//SENDNUM_LIMIT	该用户今日领取红包个数超过限制	该用户今日领取红包个数超过你在微信支付商户平台配置的上限	如有需要、请在微信支付商户平台【api安全】中重新配置 【每日同一用户领取本商户红包不允许超过的个数】。
//ILLEGAL_APPID	非法appid，请确认是否为公众号的appid，不能为APP的appid	错误传入了app的appid	接口传入的所有appid应该为公众号的appid（在mp.weixin.qq.com申请的），不能为APP的appid（在open.weixin.qq.com申请的）。
//MONEY_LIMIT	红包金额发放限制	发送红包金额不再限制范围内	每个红包金额必须大于1元，小于200元（可联系微信支付wxhongbao@tencent.com申请调高额度）
//SEND_FAILED	红包发放失败,请更换单号再重试	该红包已经发放失败	如果需要重新发放，请更换单号再发放
//FATAL_ERROR	openid和原始单参数不一致	更换了openid，但商户单号未更新	请商户检查代码实现逻辑
//金额和原始单参数不一致	更换了金额，但商户单号未更新	请商户检查代码实现逻辑	请检查金额、商户订单号是否正确
//CA_ERROR	CA证书出错，请登录微信支付商户平台下载证书	请求携带的证书出错	到商户平台下载证书，请求带上证书后重试
//SIGN_ERROR	签名错误	1、没有使用商户平台设置的商户API密钥进行加密（有可能之前设置过密钥，后来被修改了，没有使用新的密钥进行加密）。
//2、加密前没有按照文档进行参数排序（可参考文档）
//3、把值为空的参数也进行了签名。可到（http://mch.weixin.qq.com/wiki/tools/signverify/ ）验证。
//4、如果以上3步都没有问题，把请求串中(post的数据）里面中文都去掉，换成英文，试下，看看是否是编码问题。（post的数据要求是utf8）	1. 到商户平台重新设置新的密钥后重试
//2. 检查请求参数把空格去掉重试
//3. 中文不需要进行encode，使用CDATA
//4. 按文档要求生成签名后再重试
//在线签名验证工具：http://mch.weixin.qq.com/wiki/tools/signverify/
//SYSTEMERROR	请求已受理，请稍后使用原单号查询发放结果	系统无返回明确发放结果	使用原单号调用接口，查询发放结果，如果使用新单号调用接口，视为新发放请求
//XML_ERROR	输入xml参数格式错误	请求的xml格式错误，或者post的数据为空	检查请求串，确认无误后重试
//FREQ_LIMIT	超过频率限制,请稍后再试	受频率限制	请对请求做频率控制（可联系微信支付wxhongbao@tencent.com申请调高）
//NOTENOUGH	帐号余额不足，请到商户平台充值后再重试	账户余额不足	充值后重试
//OPENID_ERROR	openid和appid不匹配	openid和appid不匹配	发红包的openid必须是本appid下的openid
//MSGAPPID_ERROR	触达消息给用户appid有误	msgappid与主、子商户号的绑定关系校验失败	检查下msgappid是否填写错误，msgappid需要跟主、子商户号 有绑定关系
//ACCEPTMODE_ERROR	主、子商户号关系校验失败	服务商模式下主商户号与子商户号关系校验失败	确认传入的主商户号与子商户号是否有受理关系
//PROCESSING	请求已受理，请稍后使用原单号查询发放结果	发红包流程正在处理	二十分钟后查询,按照查询结果成功失败进行处理
//PARAM_ERROR	act_name字段必填,并且少于32个字符	请求的act_name字段填写错误	填写正确的act_name后重试
//发放金额、最小金额、最大金额必须相等	请求的金额相关字段填写错误	按文档要求填写正确的金额后重试
//红包金额参数错误	红包金额过大	修改金额重试
//appid字段必填,最长为32个字符	请求的appid字段填写错误	填写正确的appid后重试
//订单号字段必填,最长为28个字符	请求的mch_billno字段填写错误	填写正确的billno后重试
//client_ip必须是合法的IP字符串	请求的client_ip填写不正确	填写正确的IP后重试
//输入的商户号有误	请求的mchid字段非法（或者没填）	填写对应的商户号再重试
//找不到对应的商户号	请求的mchid字段填写错误	填写正确的mchid字段后重试
//nick_name字段必填，并且少于16字符	请求的nick_name字段错误	按文档填写正确的nick_name后重试
//nonce_str字段必填,并且少于32字符	请求的nonce_str字段填写不正确	按文档要求填写正确的nonce_str值后重试
//re_openid字段为必填并且少于32个字符	请求的re_openid字段非法	填写对re_openid后重试
//remark字段为必填,并且少于256字符	请求的remark字段填写错误	填写正确的remark后重试
//send_name字段为必填并且少于32字符	请求的send_name字段填写不正确	按文档填写正确的send_name字段后重试
//total_num必须为1	total_num字段值不为1	修改total_num值为1后重试
//wishing字段为必填,并且少于128个字符	缺少wishing字段	填写wishing字段再重试
//商户号和wxappid不匹配	商户号和wxappid不匹配	请修改Mchid或wxappid参数
