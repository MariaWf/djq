package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"mimi/djq/constant"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/service"
	"mimi/djq/util"
	"net/http"
	"strings"
	"mimi/djq/session"
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