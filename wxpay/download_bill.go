package wxpay

import (
	"mimi/djq/config"
	"mimi/djq/util"
	"github.com/pkg/errors"
	"net/http"
	"io/ioutil"
	"fmt"
)

func DownloadBill(billDate string) ( string, error) {
	var err error

	url := "https://api.mch.weixin.qq.com/pay/downloadbill"

	appId := config.Get("wxpay_appid") // 微信公众平台应用ID
	mchId := config.Get("wxpay_mch_id") // 微信支付商户平台商户号
	apiKey := config.Get("wxpay_key") // 微信支付商户平台API密钥

	if config.Get("running_state") == "test" {
		url = "https://api.mch.weixin.qq.com/sandboxnew/pay/downloadbill"
		apiKey, err = GetSignKey()
		if err != nil {
			return "", errors.Wrap(err,"获取测试API_KEY失败")
		}
	}

	c := NewClient(appId, mchId, apiKey)

	params := make(Params)

	params.SetString("appid",c.AppId)//公众账号ID	appid	是	String(32)	wx8888888888888888	微信分配的公众账号ID（企业号corpid即为此appId）
	params.SetString("mch_id",c.MchId)//商户号	mch_id	是	String(32)	1900000109	微信支付分配的商户号
	//设备号	device_info	否	String(32)	013467007045764	微信支付分配的终端设备号
	params.SetString("nonce_str",util.BuildUUID())//随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
	params.SetString("sign_type","MD5")//签名类型	sign_type	否	String(32)	HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	params.SetString("bill_date",billDate)//对账单日期	bill_date	是	String(8)	20140603	下载对账单的日期，格式：20140603
	params.SetString("bill_type","ALL")//账单类型	bill_type	是	String(8)	ALL
	//ALL，返回当日所有订单信息，默认值
	//SUCCESS，返回当日成功支付的订单
	//REFUND，返回当日退款订单
	//RECHARGE_REFUND，返回当日充值退款订单（相比其他对账单多一栏“返还手续费”）

	params.SetString("sign", c.Sign(params)) //签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	通过签名算法计算得出的签名值，详见签名生成算法


	// 发送查询企业付款请求
	resp, err := http.Post(url, bodyType, c.Encode(params))
	if err != nil {
		return "", err
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
	//p2,err := c.Post(url, params, false)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(p2)
	return string(bs),err
	//return c.Post(url, params, false)
}
//应用场景
//商户可以通过该接口下载历史交易清单。比如掉单、系统错误等导致商户侧和微信侧数据不一致，通过对账单核对后可校正支付状态。
//注意：
//1、微信侧未成功下单的交易不会出现在对账单中。支付成功后撤销的交易会出现在对账单中，跟原支付单订单号一致；
//2、微信在次日9点启动生成前一天的对账单，建议商户10点后再获取；
//3、对账单中涉及金额的字段单位为“元”。
//
//4、对账单接口只能下载三个月以内的账单。
//接口链接
//https://api.mch.weixin.qq.com/pay/downloadbill
//是否需要证书
//不需要。
//请求参数
//字段名	变量名	必填	类型	示例值	描述
//公众账号ID	appid	是	String(32)	wx8888888888888888	微信分配的公众账号ID（企业号corpid即为此appId）
//商户号	mch_id	是	String(32)	1900000109	微信支付分配的商户号
//设备号	device_info	否	String(32)	013467007045764	微信支付分配的终端设备号
//随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
//签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法
//签名类型	sign_type	否	String(32)	HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
//对账单日期	bill_date	是	String(8)	20140603	下载对账单的日期，格式：20140603
//账单类型	bill_type	是	String(8)	ALL
//ALL，返回当日所有订单信息，默认值
//SUCCESS，返回当日成功支付的订单
//REFUND，返回当日退款订单
//RECHARGE_REFUND，返回当日充值退款订单（相比其他对账单多一栏“返还手续费”）
//压缩账单	tar_type	否	String(8)	GZIP	非必传参数，固定值：GZIP，返回格式为.gzip的压缩包账单。不传则默认为数据流形式。
//<xml>
//<appid>wx2421b1c4370ec43b</appid>
//<bill_date>20141110</bill_date>
//<bill_type>ALL</bill_type>
//<mch_id>10000100</mch_id>
//<nonce_str>21df7dc9cd8616b56919f20d9f679233</nonce_str>
//<sign>332F17B766FC787203EBE9D6E40457A1</sign>
//</xml>
//返回结果
//失败时，返回以下字段
//
//字段名	变量名	必填	类型	示例值	描述
//返回状态码	return_code	是	String(16)	FAIL	FAIL
//返回信息	return_msg	否	String(128)	签名失败
//返回信息，如非空，为错误原因
//如：签名失败、参数格式错误等。
//成功时，数据以文本表格的方式返回，第一行为表头，后面各行为对应的字段内容，字段内容跟查询订单或退款结果一致，具体字段说明可查阅相应接口。
//第一行为表头，根据请求下载的对账单类型不同而不同(由bill_type决定),目前有：
//
//当日所有订单
//交易时间,公众账号ID,商户号,子商户号,设备号,微信订单号,商户订单号,用户标识,交易类型,交易状态,付款银行,货币种类,总金额,代金券或立减优惠金额,微信退款单号,商户退款单号,退款金额,代金券或立减优惠退款金额，退款类型，退款状态,商品名称,商户数据包,手续费,费率
//
//当日成功支付的订单
//交易时间,公众账号ID,商户号,子商户号,设备号,微信订单号,商户订单号,用户标识,交易类型,交易状态,付款银行,货币种类,总金额,代金券或立减优惠金额,商品名称,商户数据包,手续费,费率
//
//当日退款的订单
//交易时间,公众账号ID,商户号,子商户号,设备号,微信订单号,商户订单号,用户标识,交易类型,交易状态,付款银行,货币种类,总金额,代金券或立减优惠金额,退款申请时间,退款成功时间,微信退款单号,商户退款单号,退款金额,代金券或立减优惠退款金额,退款类型,退款状态,商品名称,商户数据包,手续费,费率
//
//从第二行起，为数据记录，各参数以逗号分隔，参数前增加`符号，为标准键盘1左边键的字符，字段顺序与表头一致。
//倒数第二行为订单统计标题，最后一行为统计数据
//总交易单数，总交易额，总退款金额，总代金券或立减优惠退款金额，手续费总金额
//举例如下：
//交易时间,公众账号ID,商户号,子商户号,设备号,微信订单号,商户订单号,用户标识,交易类型,交易状态,付款银行,货币种类,总金额,代金券或立减优惠金额,微信退款单号,商户退款单号,退款金额,代金券或立减优惠退款金额,退款类型,退款状态,商品名称,商户数据包,手续费,费率
//`2014-11-1016：33：45,`wx2421b1c4370ec43b,`10000100,`0,`1000,`1001690740201411100005734289,`1415640626,`085e9858e3ba5186aafcbaed1,`MICROPAY,`SUCCESS,`CFT,`CNY,`0.01,`0.0,`0,`0,`0,`0,`,`,`被扫支付测试,`订单额外描述,`0,`0.60%
//`2014-11-1016：46：14,`wx2421b1c4370ec43b,`10000100,`0,`1000,`1002780740201411100005729794,`1415635270,`085e9858e90ca40c0b5aee463,`MICROPAY,`SUCCESS,`CFT,`CNY,`0.01,`0.0,`0,`0,`0,`0,`,`,`被扫支付测试,`订单额外描述,`0,`0.60%
//总交易单数,总交易额,总退款金额,总代金券或立减优惠退款金额,手续费总金额
//`2,`0.02,`0.0,`0.0,`0
//
//错误码
//名称	描述	原因	解决方案
//SYSTEMERROR	下载失败	系统超时	请尝试再次查询。
//invalid bill_type	参数错误	请求参数未按指引进行填写	参数错误，请重新检查
//data format error
//missing parameter
//SIGN ERROR
//NO Bill Exist	账单不存在	当前商户号没有已成交的订单，不生成对账单	请检查当前商户号在指定日期内是否有成功的交易。
//Bill Creating	账单未生成	当前商户号没有已成交的订单或对账单尚未生成	请先检查当前商户号在指定日期内是否有成功的交易，如指定日期有交易则表示账单正在生成中，请在上午10点以后再下载。
//CompressGZip Error	账单压缩失败	账单压缩失败，请稍后重试	账单压缩失败，请稍后重试
//UnCompressGZip Error	账单解压失败	账单解压失败，请稍后重试	账单解压失败，请稍后重试
