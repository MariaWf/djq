package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/service"
	"mimi/djq/session"
	"mimi/djq/util"
	"net/http"
	"strings"
	"mimi/djq/constant"
)

func UserLogin(c *gin.Context) {
	if !util.GeetestCheck(c) {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(util.ErrParamException.Error()))
		return
	}

	obj := &model.User{}
	err := c.Bind(obj)

	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	//serviceObj := &service.User{}
	//obj, err = serviceObj.CheckLogin(obj)
	//if err != nil {
	//	log.Println(err)
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
	//	return
	//}
	//
	//sn, err := session.GetUi(c.Writer, c.Request)
	//if err != nil {
	//	log.Println(err)
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
	//	return
	//}
	//if err = sn.Set(session.SessionNameUiUserId, obj.Id); err != nil {
	//	log.Println(err)
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
	//	return
	//} else {
	//	http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameUiUserId, Value: obj.Id, Path: "/", MaxAge: sn.CookieMaxAge})
	//}
	//if err = sn.Set(session.SessionNameUiUserName, obj.Name); err != nil {
	//	log.Println(err)
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
	//	return
	//} else {
	//	http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameUiUserName, Value: obj.Name, Path: "/", MaxAge: sn.CookieMaxAge})
	//}
	//if codeList := obj.GetPermissionCodeList(); codeList != nil && len(codeList) != 0 {
	//	if err = sn.Set(session.SessionNameUiPermission, strings.Join(codeList, constant.Split4Permission)); err != nil {
	//		log.Println(err)
	//		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
	//		return
	//	} else {
	//		http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameUiPermission, Value: strings.Join(codeList, constant.Split4Permission), Path: "/", MaxAge: sn.CookieMaxAge})
	//	}
	//}

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
	argObj.DisplayNames = []string{"id",  "mobile", "presentChance", "locked"}
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
	result := service.ResultGet(serviceObj,c.Param("id"))
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
	result := service.ResultAdd(serviceObj,obj)
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
	result := service.ResultUpdate(serviceObj,obj)
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
