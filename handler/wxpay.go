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
	"mimi/djq/service"
	"strconv"
	"io"
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
	bs, err := ioutil.ReadAll(Notify4unified_order(c.Request.Body))
	if err != nil {
		panic(err)
	}
	c.Writer.Write(bs)
}

func WxpayNotifyRefund(c *gin.Context) {
	bs, err := ioutil.ReadAll(Notify4refund(c.Request.Body))
	if err != nil {
		panic(err)
	}
	c.Writer.Write(bs)
}

func WxpayDownloadBill(c *gin.Context) {
	w, err := wxpay.DownloadBill(c.Query("billDate"))
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, util.BuildSuccessResult(w))
}

func Notify4unified_order(r io.Reader) io.Reader {
	var err error
	client := wxpay.NewDefaultClient()
	if config.Get("running_state") == "test" {
		client.ApiKey, err = wxpay.GetSignKey()
		if err != nil {
			panic(errors.Wrap(err, "获取测试API_KEY失败"))
		}
	}
	params := client.Decode(r)
	fmt.Println(params)
	p2 := make(wxpay.Params)
	if params["return_code"] == "FAIL" {
		log.Println(params)
		p2.SetString("return_code", "FAIL")
		p2.SetString("return_msg", "接收失败")
	} else if !client.CheckSign(params) {
		p2.SetString("return_code", "FAIL")
		p2.SetString("return_msg", "签名失败")
	} else {
		payOrderNumber := params["out_trade_no"]
		//timeEnd := params["time_end"]
		totalFeeStr := params["total_fee"]
		cashFeeStr := params["cash_fee"]
		appId := params["appid"]
		mchId := params["mch_id"]

		if client.AppId != appId || client.MchId != mchId || totalFeeStr != cashFeeStr || payOrderNumber == "" {
			p2.SetString("return_code", "FAIL")
			p2.SetString("return_msg", "参数格式校验错误")
		} else if params["result_code"] == "FAIL" {
			p2.SetString("return_code", "SUCCESS")
			p2.SetString("return_msg", "")
			//serviceObj := &service.CashCouponOrder{}
			//_, err := serviceObj.CancelOrder(payOrderNumber)
			//if err != nil {
			//	err = errors.Wrap(err, payOrderNumber)
			//	log.Println(err)
			//	//cache.Set(cache.CacheNameWxpayErrorPayOrderNumberCancel + payOrderNumber, err.Error(), time.Hour * 24 * 7)
			//	p2.SetString("return_code", "FAIL")
			//	p2.SetString("return_msg", "系统异常")
			//} else {
			//	p2.SetString("return_code", "SUCCESS")
			//	p2.SetString("return_msg", "")
			//}
			////cache.Set(cache.CacheNameWxpayPayOrderNumberCancel + payOrderNumber, idListStr, time.Hour * 24 * 7)
		} else {
			totalFee, err := strconv.Atoi(totalFeeStr)
			if err != nil {
				p2.SetString("return_code", "FAIL")
				p2.SetString("return_msg", "参数格式校验错误")
			} else {
				serviceObj := &service.CashCouponOrder{}
				_, err := serviceObj.ConfirmOrder(payOrderNumber, totalFee)
				if err != nil {
					err = errors.Wrap(err, payOrderNumber + "_" + strconv.Itoa(totalFee))
					log.Println(err)
					//cache.Set(cache.CacheNameWxpayErrorPayOrderNumberConfirm + payOrderNumber, err.Error(), time.Hour * 24 * 7)
					p2.SetString("return_code", "FAIL")
					p2.SetString("return_msg", "系统异常")
				} else {
					p2.SetString("return_code", "SUCCESS")
					p2.SetString("return_msg", "")
				}
				//cache.Set(cache.CacheNameWxpayPayOrderNumberConfirm + payOrderNumber, idListStr, time.Hour * 24 * 7)
			}
		}
	}
	fmt.Println(p2)
	return client.Encode(p2)
}

func Notify4refund(r io.Reader) io.Reader {
	var err error
	client := wxpay.NewDefaultClient()
	if config.Get("running_state") == "test" {
		client.ApiKey, err = wxpay.GetSignKey()
		if err != nil {
			panic(errors.Wrap(err, "获取测试API_KEY失败"))
		}
	}
	params := client.Decode(r)
	p2 := make(wxpay.Params)
	if params["return_code"] == "FAIL" {
		p2.SetString("return_code", "FAIL")
		p2.SetString("return_msg", "接收失败")

	} else {
		reqInfo := params["req_info"]

		newP, err := client.Aes256EcbDecrypt(reqInfo)
		if err != nil {
			log.Println(err)
			p2.SetString("return_code", "FAIL")
			p2.SetString("return_msg", "解密失败")
		} else {
			appId := params["appid"]
			mchId := params["mch_id"]
			payOrderNumber := newP["out_trade_no"]
			refundOrderNumber := newP["out_refund_no"]
			settlementRefundFeeStr := newP["settlement_refund_fee"]
			refundFeeStr := newP["refund_fee"]

			if client.AppId != appId || client.MchId != mchId || settlementRefundFeeStr != refundFeeStr || refundOrderNumber == "" || payOrderNumber == "" {
				p2.SetString("return_code", "FAIL")
				p2.SetString("return_msg", "参数格式校验错误")
			} else if newP["refund_status"] == "SUCCESS" {
				serviceObj := &service.Refund{}
				err = serviceObj.ConfirmByRefundOrderNumber(refundOrderNumber)
				if err != nil {
					log.Println(err)
					p2.SetString("return_code", "FAIL")
					p2.SetString("return_msg", "系统异常")
				} else {
					p2.SetString("return_code", "SUCCESS")
					p2.SetString("return_msg", "")
				}
			} else {
				serviceObj := &service.Refund{}
				err = serviceObj.FailCloseByRefundOrderNumber(refundOrderNumber)
				if err != nil {
					log.Println(err)
					p2.SetString("return_code", "FAIL")
					p2.SetString("return_msg", "系统异常")
				} else {
					p2.SetString("return_code", "SUCCESS")
					p2.SetString("return_msg", "")
				}
				p2.SetString("return_code", "SUCCESS")
				p2.SetString("return_msg", "")
			}
		}
	}
	if p2.GetString("return_code") != "SUCCESS" {
		log.Println(params)
		log.Println(p2.GetString("return_msg"))
	}
	return client.Encode(p2)
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