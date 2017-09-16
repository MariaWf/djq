package handler

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"mimi/djq/wxpay"
	"net/http"
	"mimi/djq/util"
	"log"
	"io/ioutil"
	"html/template"
	"mimi/djq/session"
	"strings"
	"mimi/djq/config"
	"encoding/json"
	"github.com/pkg/errors"
)

var ErrWxpayGetOpenIdFail = errors.New("获取微信OpenId失败")

func Wxpay(c *gin.Context) {
	values := make(map[string]interface{})
	t, _ := template.ParseFiles("html/template/wxpay.html")
	t.Execute(c.Writer, values)
}

func WxpayGetOpenId(c *gin.Context) {
	sn, err := session.GetUi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrWxpayGetOpenIdFail.Error()))
		return
	}
	openId, err := sn.Get(session.SessionNameUiUserOpenId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrWxpayGetOpenIdFail.Error()))
		return
	}
	if openId == "" {
		code := c.Query("code")
		if code == "" {
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
			return
		}
		url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid={appId}&secret={appSecret}&code={code}&grant_type=authorization_code"
		url = strings.Replace(url, "{appId}", config.Get("wxpay_appid"), -1)
		url = strings.Replace(url, "{appSecret}", config.Get("wxpay_app_secret"), -1)
		url = strings.Replace(url, "{code}", code, -1)

		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrWxpayGetOpenIdFail.Error()))
			return
		}
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrWxpayGetOpenIdFail.Error()))
			return
		}
		values := make(map[string]interface{})
		fmt.Println(string(bs))
		err = json.Unmarshal(bs, &values)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrWxpayGetOpenIdFail.Error()))
			return
		}
		if values["errmsg"] != nil {
			log.Println(values)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrWxpayGetOpenIdFail.Error()))
			return
		} else {
			openId = values["openid"].(string)
			sn.Set(session.SessionNameUiUserOpenId, openId)
			http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameUiUserOpenId, Value: openId, Path: "/", MaxAge: sn.CookieMaxAge})
		}
	}
	c.JSON(http.StatusOK, util.BuildSuccessResult(openId))
}

func WxpayNotifyUnifiedOrder(c *gin.Context) {
	fmt.Println(c.Request.Method)
	var err error
	client := wxpay.NewDefaultClient()
	if config.Get("running_state") == "test" {
		client.ApiKey, err = wxpay.GetSignKey()
		if err != nil {
			panic(errors.Wrap(err, "获取测试API_KEY失败"))
		}
	}
	params := client.Decode(c.Request.Body)
	fmt.Println(client.CheckSign(params), params)

	p2 := make(wxpay.Params)
	p2.SetString("return_code", "SUCCESS")
	p2.SetString("return_msg", "")

	bs, err := ioutil.ReadAll(client.Encode(p2))
	if err != nil {
		panic(err)
	}
	c.Writer.Write(bs)
}

func WxpayConfig(c *gin.Context) {
	client := wxpay.NewDefaultClient()
	url := c.Request.Referer()
	fmt.Println(url)
	//url = "http://djq.tunnel.qydev.com/open/wxpay/"
	params, err := wxpay.GetConfigSignatureMap(url)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	params.SetString("appId", client.AppId)
	params.SetString("title", "标题mimi")
	params.SetString("link", "www.baidu.com")
	params.SetString("desc", "描述mimi")
	params.SetString("imgUrl", "https://www.baidu.com/img/bd_logo1.png")
	result := util.BuildSuccessResult(params)
	c.JSON(http.StatusOK, result)
}

func WxpayQuery(c *gin.Context) {
	c.Query("payOrderNumber")
	params, err := wxpay.OrderQuery(c.Query("payOrderNumber"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	fmt.Println(params)
	result := util.BuildSuccessResult(params)
	c.JSON(http.StatusOK, result)
}