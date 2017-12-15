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
	"strconv"
	"strings"
	"time"
)

func UserLogin(c *gin.Context) {
	//if !util.GeetestCheck(c) {
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(util.ErrParamException.Error()))
	//	return
	//}

	mobile := c.PostForm("mobile")
	captcha := c.PostForm("captcha")
	promotionalPartnerId := c.PostForm("promotionalPartnerId")

	sn, err := session.GetUi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	value, err := sn.Get(session.SessionNameUiUserCaptcha)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	if value != captcha {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("手机验证码不正确"))
		return
	}

	count, err := sn.Get(session.SessionNameUiUserLoginCount)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	newCount := 0
	if count != "" {
		newCount, err = strconv.Atoi(count)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
			return
		}
	}
	newCount++
	if newCount > 10 {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("登录频繁，请休息5分钟再来"))
		return
	}
	err = sn.SetTemp(session.SessionNameUiUserLoginCount, strconv.Itoa(newCount), time.Minute*5)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}

	obj := &model.User{}
	obj.PromotionalPartnerId = promotionalPartnerId
	obj.Mobile = mobile
	serviceObj := &service.User{}
	obj, err = serviceObj.Register(obj)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	err = sn.Set(session.SessionNameUiUserId, obj.Id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	err = sn.Set(session.SessionNameUiUserMobile, obj.Mobile)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameUiUserId, Value: obj.Id, Path: "/", MaxAge: sn.CookieMaxAge})
	http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameUiUserMobile, Value: obj.Mobile, Path: "/", MaxAge: sn.CookieMaxAge})
	result := util.BuildSuccessResult(obj)
	c.JSON(http.StatusOK, result)

}

func UserLogout(c *gin.Context) {
	sn, err := session.GetUi(c.Writer, c.Request)
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
	http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameUiUserId, Value: "", Path: "/", MaxAge: -1})
	http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameUiUserMobile, Value: "", Path: "/", MaxAge: -1})
	http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameUiUserOpenId, Value: "", Path: "/", MaxAge: -1})
	http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameSiShopAccountId, Value: "", Path: "/", MaxAge: -1})
	http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameSiShopAccountName, Value: "", Path: "/", MaxAge: -1})
	http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameSiShopAccountOpenId, Value: "", Path: "/", MaxAge: -1})

	result := util.BuildSuccessResult(nil)
	c.JSON(http.StatusOK, result)
}

func UserCheckLogin(c *gin.Context) {
	if c.Request.URL.Path != "/ui/login" && c.Request.URL.Path != "/ui/logout" {
		sn, err := session.GetUi(c.Writer, c.Request)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
			return
		}
		id, err := sn.Get(session.SessionNameUiUserId)
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

func UserActionShareTimelineResponse(c *gin.Context) {
	sn, err := session.GetUi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	id, err := sn.Get(session.SessionNameUiUserId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	serviceObj := &service.User{}
	if id == "" {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("未知用户"))
		return
	}
	err = serviceObj.SharedActionResponse(id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult("")
	c.JSON(http.StatusOK, result)
}

func UserGet4Ui(c *gin.Context) {
	sn, err := session.GetUi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	id, err := sn.Get(session.SessionNameUiUserId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	serviceObj := &service.User{}
	result := service.ResultGet(serviceObj, id)
	c.JSON(http.StatusOK, result)
}

func UserList(c *gin.Context) {
	argObj := &arg.User{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	locked := c.Query("locked")
	if locked == "true" {
		argObj.LockedOnly = true
	} else if locked == "false" {
		argObj.UnlockedOnly = true
	}

	argObj.OrderBy = "mobile"
	argObj.DisplayNames = []string{"id", "promotionalPartnerId", "mobile", "presentChance", "shared", "locked"}
	serviceObj := &service.User{}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func UserGet(c *gin.Context) {
	serviceObj := &service.User{}
	//obj, err := service.Get(serviceObj,c.Param("id"))
	//if err != nil {
	//	log.Println(err)
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
	//	return
	//}
	//result := util.BuildSuccessResult(obj)
	result := service.ResultGet(serviceObj, c.Param("id"))
	c.JSON(http.StatusOK, result)
}

//func UserGetSelf(c *gin.Context) {
//	serviceObj := &service.User{}
//	sn, err := session.GetUi(c.Writer, c.Request)
//	if err != nil {
//		log.Println(err)
//		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
//		return
//	}
//	id, err := sn.Get(session.SessionNameUiUserId)
//	if err != nil {
//		log.Println(err)
//		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
//		return
//	}
//	obj, err := service.Get(serviceObj, id)
//	if err != nil {
//		log.Println(err)
//		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
//		return
//	}
//	result := util.BuildSuccessResult(obj)
//	c.JSON(http.StatusOK, result)
//}

func UserPost(c *gin.Context) {
	obj := &model.User{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.User{}
	//obj, err := serviceObj.Add(user)
	//if err != nil {
	//	log.Println(err)
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
	//	return
	//}
	//result := util.BuildSuccessResult(obj.Id)
	result := service.ResultAdd(serviceObj, obj)
	c.JSON(http.StatusOK, result)
}

func UserPatch(c *gin.Context) {
	obj := &model.User{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.User{}
	//obj, err = serviceObj.Update(obj)
	//if err != nil {
	//	log.Println(err)
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
	//	return
	//}
	//result := util.BuildSuccessResult(obj.Id)
	result := service.ResultUpdate(serviceObj, obj)
	c.JSON(http.StatusOK, result)
}

func UserDelete(c *gin.Context) {

	//serviceObj := &service.User{}
	//count, err := serviceObj.Delete(ids...)
	//if err != nil {
	//	log.Println(err)
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
	//	return
	//}
	//result := util.BuildSuccessResult(count)
	ids := strings.Split(c.Query("ids"), constant.Split4Id)

	serviceObj := &service.User{}
	argObj := &arg.User{}
	argObj.IdsIn = ids
	result := service.ResultBatchDelete(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}
