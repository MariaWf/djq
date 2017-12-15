package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"mimi/djq/constant"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/service"
	"mimi/djq/session"
	"mimi/djq/util"
	"net/http"
	"strings"
)

func ShopAccountLogin(c *gin.Context) {
	if !util.GeetestCheck(c) {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(util.ErrParamException.Error()))
		return
	}

	obj := &model.ShopAccount{}
	err := c.Bind(obj)

	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	serviceObj := &service.ShopAccount{}
	obj, err = serviceObj.CheckLogin(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}

	sn, err := session.GetSi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	if err = sn.Set(session.SessionNameSiShopAccountId, obj.Id); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	} else {
		http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameSiShopAccountId, Value: obj.Id, Path: "/", MaxAge: sn.CookieMaxAge})
	}
	if err = sn.Set(session.SessionNameSiShopAccountName, obj.Name); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameSiShopAccountName, Value: obj.Name, Path: "/", MaxAge: sn.CookieMaxAge})

	result := util.BuildSuccessResult(obj)
	c.JSON(http.StatusOK, result)
}

func ShopAccountLogout(c *gin.Context) {
	sn, err := session.GetSi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	if err = sn.Del(); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameSiShopAccountId, Value: "", Path: "/", MaxAge: -1})
	http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameSiShopAccountName, Value: "", Path: "/", MaxAge: -1})
	http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameSiShopAccountOpenId, Value: "", Path: "/", MaxAge: -1})
	result := util.BuildSuccessResult(nil)
	c.JSON(http.StatusOK, result)
}

func ShopAccountCheckLogin(c *gin.Context) {
	if c.Request.URL.Path != "/si/login" && c.Request.URL.Path != "/si/logout" {
		sn, err := session.GetSi(c.Writer, c.Request)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
			return
		}
		id, err := sn.Get(session.SessionNameSiShopAccountId)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
			return
		}
		if id == "" {
			c.AbortWithStatusJSON(http.StatusOK, util.BuildNeedLoginResult())
		}
	}
}

func ShopAccountGetSelf(c *gin.Context) {
	serviceObj := &service.ShopAccount{}
	sn, err := session.GetSi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	id, err := sn.Get(session.SessionNameSiShopAccountId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	obj, err := service.Get(serviceObj, id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	obj.(*model.ShopAccount).Password = ""
	result := util.BuildSuccessResult(obj)
	//result := service.ResultGet(serviceObj, c.Param("id"))
	c.JSON(http.StatusOK, result)
}

func ShopAccountGetMoney4Si(c *gin.Context) {
	sn, err := session.GetSi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	shopAccountId, err := sn.Get(session.SessionNameSiShopAccountId)
	if err != nil || shopAccountId == "" {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	openId, err := sn.Get(session.SessionNameSiShopAccountOpenId)
	if err != nil{
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	if openId == "" {
		sn2, err2 := session.GetUi(c.Writer, c.Request)
		if err2 != nil {
			log.Println(err2)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
			return
		}
		openId, err = sn2.Get(session.SessionNameUiUserOpenId)
		if err != nil{
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
			return
		}
	}
	if openId == "" {
		//util.CookieCleanAll(c.Writer)
		http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameUiUserId, Value: "", Path: "/", MaxAge: -1})
		http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameUiUserMobile, Value: "", Path: "/", MaxAge: -1})
		http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameUiUserOpenId, Value: "", Path: "/", MaxAge: -1})
		http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameSiShopAccountId, Value: "", Path: "/", MaxAge: -1})
		http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameSiShopAccountName, Value: "", Path: "/", MaxAge: -1})
		http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameSiShopAccountOpenId, Value: "", Path: "/", MaxAge: -1})
		c.AbortWithStatusJSON(http.StatusOK, util.BuildNeedLoginResult())
		//c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("未知微信openId"))
		return
	}
	serviceObj := &service.ShopAccount{}
	money, err := serviceObj.GetMoney(shopAccountId, openId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(money)
	c.JSON(http.StatusOK, result)
}

func ShopAccountActionGetPresentOrderOrCashCouponOrder4Si(c *gin.Context) {
	sn, err := session.GetSi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	shopAccountId, err := sn.Get(session.SessionNameSiShopAccountId)
	if err != nil || shopAccountId == "" {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	number := strings.ToUpper(c.Query("number"))
	if number == "" {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("编号不能为空"))
		return
	}
	var result *util.ResultVO
	if strings.Index(number, "C") == 0 {
		serviceCashCouponOrder := &service.CashCouponOrder{}
		argCashCouponOrder := &arg.CashCouponOrder{}
		argCashCouponOrder.NumberEqual = number
		argCashCouponOrder.PageSize = 1
		//argCashCouponOrder.StatusEqual = strconv.Itoa(constant.CashCouponOrderStatusPaidNotUsed)
		list, err := service.Find(serviceCashCouponOrder, argCashCouponOrder)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
			return
		}
		if len(list) == 0 {
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("找不到匹配代金券订单"))
			return
		}
		cashCouponOrder := list[0].(*model.CashCouponOrder)
		if cashCouponOrder.Status == constant.CashCouponOrderStatusUsed {
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("代金券订单已消费"))
			return
		}
		if cashCouponOrder.Status != constant.CashCouponOrderStatusPaidNotUsed {
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("代金券订单并非处于可消费状态"))
			return
		}
		serviceCashCoupon := &service.CashCoupon{}
		cashCoupon, err := service.Get(serviceCashCoupon, cashCouponOrder.CashCouponId)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
			return
		}
		serviceObj := &service.ShopAccount{}
		obj, err := service.Get(serviceObj, shopAccountId)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
			return
		}
		if cashCoupon.(*model.CashCoupon).ShopId != obj.(*model.ShopAccount).ShopId {
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("代金券不属于本店"))
			return
		}
		cashCouponOrder.CashCoupon = cashCoupon.(*model.CashCoupon)
		serviceUser := &service.User{}
		user, err := service.Get(serviceUser, cashCouponOrder.UserId)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
			return
		}
		cashCouponOrder.User = user.(*model.User)
		result = util.BuildSuccessResult(cashCouponOrder)
	} else if strings.Index(number, "P") == 0 {
		servicePresentOrder := &service.PresentOrder{}
		argPresentOrder := &arg.PresentOrder{}
		argPresentOrder.NumberEqual = number
		argPresentOrder.PageSize = 1
		//argPresentOrder.StatusEqual = strconv.Itoa(constant.PresentOrderStatusWaiting2Receive)
		list, err := service.Find(servicePresentOrder, argPresentOrder)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
			return
		}
		if len(list) == 0 {
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("找不到匹配代抽奖记录"))
			return
		}
		presentOrder := list[0].(*model.PresentOrder)
		if presentOrder.Status == constant.PresentOrderStatusReceived {
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("奖品已被领取"))
			return
		}
		servciePresent := &service.Present{}
		present, err := service.Get(servciePresent, presentOrder.PresentId)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
			return
		}
		presentOrder.Present = present.(*model.Present)
		serviceUser := &service.User{}
		user, err := service.Get(serviceUser, presentOrder.UserId)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
			return
		}
		presentOrder.User = user.(*model.User)
		result = util.BuildSuccessResult(presentOrder)
	} else {
		result = util.BuildFailResult("无法识别编号")
	}
	c.JSON(http.StatusOK, result)
}

func ShopAccountList(c *gin.Context) {
	argObj := &arg.ShopAccount{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	argObj.OrderBy = "name"

	serviceObj := &service.ShopAccount{}
	argObj.DisplayNames = []string{"id", "shopId", "name", "description", "moneyChance", "totalMoney", "locked"}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func ShopAccountGet(c *gin.Context) {
	serviceObj := &service.ShopAccount{}
	obj, err := service.Get(serviceObj, c.Param("id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	obj.(*model.ShopAccount).Password = ""
	result := util.BuildSuccessResult(obj)
	c.JSON(http.StatusOK, result)
}

func ShopAccountPost(c *gin.Context) {
	obj := &model.ShopAccount{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.ShopAccount{}
	_, err = serviceObj.Add(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(obj)
	c.JSON(http.StatusOK, result)
}

func ShopAccountPatch(c *gin.Context) {
	obj := &model.ShopAccount{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.ShopAccount{}
	_, err = serviceObj.Update(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(obj)
	//result := service.ResultUpdate(serviceObj, obj, "name", "description", "moneyChance", "totalMoney", "locked")
	c.JSON(http.StatusOK, result)
}

func ShopAccountDelete(c *gin.Context) {
	ids := strings.Split(c.Query("ids"), constant.Split4Id)

	serviceObj := &service.ShopAccount{}
	argObj := &arg.ShopAccount{}
	argObj.IdsIn = ids
	result := service.ResultBatchDelete(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}
