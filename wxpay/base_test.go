package wxpay

import (
	"testing"
	"mimi/djq/util"
	"math/rand"
)

func TestGetSignKey(t *testing.T) {
	signKey, err := GetSignKey()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(signKey)
	}
}

func TestUnifiedOrder(t *testing.T) {
	payOrderNumber := util.BuildUUID()
	totalFee := rand.Intn(10000) + 1
	clientIp := "192.168.1.1"
	openId := ""
	params, err := UnifiedOrder(payOrderNumber, totalFee, clientIp,openId)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(params)
		t.Log(params.GetString("prepay_id"))
	}
}

func TestGetAccessToken(t *testing.T) {
	obj, err := GetAccessToken()
	if err!=nil{
		t.Error(err)
	}else{
		t.Log(obj)
	}
}

func TestGetJsapiTicket(t *testing.T) {
	obj, err := GetJsapiTicket()
	if err!=nil{
		t.Error(err)
	}else{
		t.Log(obj)
	}
}

func TestGetSignatureMap(t *testing.T) {
	obj,err := GetConfigSignatureMap("http://www.baidu.com")
	if err!=nil{
		t.Error(err)
	}else{
		t.Log(obj)
	}
}