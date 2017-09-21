package wxpay

import (
	"testing"
	"log"
)

func TestClient_WithCert(t *testing.T) {
	appId  := "" // 微信公众平台应用ID
	mchId  := "" // 微信支付商户平台商户号
	apiKey := "" // 微信支付商户平台API密钥

	// 微信支付商户平台证书路径
	certFile   := "cert/apiclient_cert.pem"
	keyFile    := "cert/apiclient_key.pem"
	rootcaFile := "cert/rootca.pem"

	c := NewClient(appId, mchId, apiKey)

	// 附着商户证书
	err := c.WithCert(certFile, keyFile, rootcaFile)
	if err != nil {
		log.Fatal(err)
	}

	params := make(Params)
	// 查询企业付款接口请求参数
	params.SetString("appid", c.AppId)
	params.SetString("mch_id", c.MchId)
	params.SetString("nonce_str", "5K8264ILTKCH16CQ2502SI8ZNMTM67VS")  // 随机字符串
	params.SetString("partner_trade_no", "10000098201411111234567890") // 商户订单号
	params.SetString("sign", c.Sign(params))                           // 签名

	// 查询企业付款接口请求URL
	url := "https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo"

	// 发送查询企业付款请求
	ret, err := c.Post(url, params, true)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(ret)
}

func TestClient_CheckSign(t *testing.T) {
	//appid:wxfdde12865d9e1e4c bank_type:CFT out_trade_no:b03dcf56813b4e08b8f2395dd0f248b7
	//result_code:SUCCESS device_info:WEB fee_type:CNY openid:oDKzK07AOK4UEBy5rUwodTAljRGM
	//return_code:SUCCESS transaction_id:4200000028201709193011266765 cash_fee:200 sign:96C34FB6037D06F59BC1182AD9F4AC76
	//total_fee:200 is_subscribe:Y mch_id:1488733562 nonce_str:813edcce64d342d0a6678299304901ce time_end:20170919154412 trade_type:JSAPI
	client := NewDefaultClient()
	params := make(Params)
	params["appid"] = "wxfdde12865d9e1e4c"
	params["bank_type"] = "CFT"
	params["out_trade_no"] = "b03dcf56813b4e08b8f2395dd0f248b7"
	params["result_code"] = "SUCCESS"
	params["device_info"] = "WEB"
	params["fee_type"] = "CNY"
	params["openid"] = "oDKzK07AOK4UEBy5rUwodTAljRGM"
	params["return_code"] = "SUCCESS"
	params["transaction_id"] = "4200000028201709193011266765"
	params["cash_fee"] = "200"
	params["total_fee"] = "200"
	params["is_subscribe"] = "Y"
	params["mch_id"] = "1488733562"
	params["nonce_str"] = "813edcce64d342d0a6678299304901ce"
	params["time_end"] = "20170919154412"
	params["trade_type"] = "JSAPI"
	params["sign"] = "96C34FB6037D06F59BC1182AD9F4AC76"
	t.Log(client.Sign(params))
	t.Log(client.CheckSign(params))
}