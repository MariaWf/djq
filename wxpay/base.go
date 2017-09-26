package wxpay

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"mimi/djq/cache"
	"mimi/djq/config"
	"mimi/djq/util"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GetSignKey() (signKey string, err error) {
	cacheName := "sandBoxSignKey"
	if signKey, err = cache.Get(cacheName); err != nil {
		err = errors.Wrap(err, "获取测试API_KEY失败")
	} else if signKey == "" {
		//url := "https://api.mch.weixin.qq.com/sandboxnew/pay/getsignkey"
		//values := make(map[string][]string)
		//response, err := http.PostForm(url, values)
		//if err != nil {
		//	panic(err)
		//}
		//defer response.Body.Close()    //请求完了关闭回复主体
		//body, err := ioutil.ReadAll(response.Body)
		//fmt.Println(string(body))
		appId := config.Get("wxpay_appid")  // 微信公众平台应用ID
		mchId := config.Get("wxpay_mch_id") // 微信支付商户平台商户号
		apiKey := config.Get("wxpay_key")   // 微信支付商户平台API密钥

		c := NewClient(appId, mchId, apiKey)

		//// 微信支付商户平台证书路径
		//certFile   := config.Get("wxpay_cert_file")
		//keyFile    := config.Get("wxpay_key_file")
		//rootcaFile := config.Get("wxpay_rootca_file")
		//
		//// 附着商户证书
		//err := c.WithCert(certFile, keyFile, rootcaFile)
		//if err != nil {
		//	log.Fatal(err)
		//}

		params := make(Params)
		params.SetString("mch_id", c.MchId)
		params.SetString("nonce_str", util.BuildUUID()) // 随机字符串
		//params.SetString("nonce_str", util.BuildUUID())  // 随机字符串
		params.SetString("sign", c.Sign(params)) // 签名

		// 查询企业付款接口请求URL
		url := "https://api.mch.weixin.qq.com/sandboxnew/pay/getsignkey"

		// 发送查询企业付款请求
		p1, err := c.Post(url, params, false)
		if err != nil {
			return "", errors.Wrap(err, "获取测试API_KEY失败")
		}
		signKey = p1.GetString("sandbox_signkey")
		if err := cache.Set(cacheName, signKey, time.Second*7000); err != nil {
			return "", errors.Wrap(err, "获取测试API_KEY失败")
		}
	}
	return
}

func GetAccessToken() (string, error) {
	cacheName := "wxAccessToken"
	var accessToken string
	var err error
	if accessToken, err = cache.Get(cacheName); err != nil {
		return "", errors.Wrap(err, "获取微信AccessToken异常")
	} else if accessToken == "" {
		url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid={APP_ID}&secret={APP_SECRET}"
		url = strings.Replace(url, "{APP_ID}", config.Get("wxpay_appid"), -1)
		url = strings.Replace(url, "{APP_SECRET}", config.Get("wxpay_app_secret"), -1)
		response, err := http.PostForm(url, make(map[string][]string))
		if err != nil {
			return "", errors.Wrap(err, "获取微信AccessToken异常")
		}
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", errors.Wrap(err, "获取微信AccessToken异常")
		}
		var result map[string]interface{}
		err = json.Unmarshal(bytes, &result)
		if err != nil {
			return "", errors.Wrap(err, "获取微信AccessToken异常")
		}
		if result["errcode"] != nil {
			return "", errors.Wrap(err, "获取微信AccessToken异常")
		}
		accessToken = result["access_token"].(string)
		if err := cache.Set(cacheName, accessToken, time.Second*7000); err != nil {
			return "", errors.Wrap(err, "获取微信AccessToken异常")
		}
	}
	return accessToken, nil
}

func GetJsapiTicket() (string, error) {
	cacheName := "wxJsapiTicket"
	var jsapiTicket string
	var err error
	if jsapiTicket, err = cache.Get(cacheName); err != nil {
		return "", errors.Wrap(err, "获取微信JsapiTicket异常")
	} else if jsapiTicket == "" {
		url := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token={ACCESS_TOKEN}&type=jsapi"
		accessToken, err := GetAccessToken()
		if err != nil {
			return "", errors.Wrap(err, "获取微信JsapiTicket异常")
		}
		url = strings.Replace(url, "{ACCESS_TOKEN}", accessToken, -1)
		response, err := http.PostForm(url, make(map[string][]string))
		if err != nil {
			return "", errors.Wrap(err, "获取微信JsapiTicket异常")
		}
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", errors.Wrap(err, "获取微信JsapiTicket异常")
		}
		var result map[string]interface{}
		err = json.Unmarshal(bytes, &result)
		if err != nil {
			return "", errors.Wrap(err, "获取微信JsapiTicket异常")
		}
		if result["errmsg"].(string) != "ok" {
			return "", errors.Wrap(err, "获取微信JsapiTicket异常")
		}
		jsapiTicket = result["ticket"].(string)
		if err := cache.Set(cacheName, jsapiTicket, time.Second*7000); err != nil {
			return "", errors.Wrap(err, "获取微信JsapiTicket异常")
		}
	}
	return jsapiTicket, nil
}

func GetConfigSignatureMap(url string) (Params, error) {
	result := make(Params)
	ticket, err := GetJsapiTicket()
	if err != nil {
		return result, err
	}
	result["jsapi_ticket"] = ticket
	result["noncestr"] = util.BuildUUID()
	result["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	result["url"] = url

	//result["jsapi_ticket"] = "sM4AOVdWfPE4DxkXGEs8VMCPGGVi4C3VM0P37wVUCFvkVAy_90u5h9nbSlYy3-Sl-HhTdfl2fzFy1AOcHKP7qg"
	//result["noncestr"] = "Wm3WZYTPz0wzccnW"
	//result["timestamp"] = "1414587457"
	//result["url"] = "http://mp.weixin.qq.com?params=value"
	//0f9de62fce790f9a083d5c99e95740ceb90c27ed
	result["signature"] = Signature(result)
	return result, nil
}

// 生成签名
func Signature(params Params) string {
	var keys = make([]string, 0, len(params))
	for k, _ := range params {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	for i, k := range keys {
		if len(params.GetString(k)) > 0 {
			buf.WriteString(k)
			buf.WriteString(`=`)
			buf.WriteString(params.GetString(k))
			if i < len(keys)-1 {
				buf.WriteString(`&`)
			}
		}
	}
	sum := sha1.Sum(buf.Bytes())
	str := hex.EncodeToString(sum[:])
	return strings.ToLower(str)
}

// 生成签名
func Signature4Pay(params Params) string {
	var keys = make([]string, 0, len(params))
	for k, _ := range params {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	for i, k := range keys {
		if len(params.GetString(k)) > 0 {
			buf.WriteString(k)
			buf.WriteString(`=`)
			buf.WriteString(params.GetString(k))
			if i < len(keys)-1 {
				buf.WriteString(`&`)
			}
		}
	}
	sum := md5.Sum(buf.Bytes())
	str := hex.EncodeToString(sum[:])
	return strings.ToUpper(str)
}
