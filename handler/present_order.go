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

//func PresentOrderCount4Ui(c *gin.Context) {
//	sn, err := session.GetUi(c.Writer, c.Request)
//	if err != nil {
//		log.Println(err)
//		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
//		return
//	}
//	userId, err := sn.Get(session.SessionNameUiUserId)
//	if err != nil || userId == "" {
//		log.Println(err)
//		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
//		return
//	}
//	argObj := &arg.PresentOrder{}
//	argObj.UserIdEqual = userId
//	argObj.StatusEqual = strconv.Itoa(constant.PresentOrderStatusWaiting2Receive)
//	serviceObj := &service.PresentOrder{}
//	count, err := service.Count(serviceObj, argObj)
//	if err != nil {
//		log.Println(err)
//		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
//		return
//	}
//	result := util.BuildSuccessResult(count)
//	c.JSON(http.StatusOK, result)
//}

func PresentOrderList4Ui(c *gin.Context) {
	sn, err := session.GetUi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	userId, err := sn.Get(session.SessionNameUiUserId)
	if err != nil || userId == "" {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	argObj := &arg.PresentOrder{}
	argObj.UserIdEqual = userId
	argObj.OrderBy = "status,number"
	argObj.DisplayNames = []string{"id", "presentId", "userId", "number", "status"}
	serviceObj := &service.PresentOrder{}
	list, err := service.Find(serviceObj, argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	if len(list) > 0 {
		presentIds := make([]string, 0, 10)
		for _, v := range list {
			exist := false
			pId := v.(*model.PresentOrder).PresentId
			for _, p := range presentIds {
				if p == pId {
					exist = true
				}
			}
			if !exist {
				presentIds = append(presentIds, pId)
			}
		}
		servicePresent := &service.Present{}
		argPresent := &arg.Present{}
		argPresent.IdsIn = presentIds
		presents, err := service.Find(servicePresent, argPresent)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
			return
		}
		for _, v := range list {
			pId := v.(*model.PresentOrder).PresentId
			for _, p := range presents {
				if pId == p.(*model.Present).Id {
					v.(*model.PresentOrder).Present = p.(*model.Present)
					break
				}
			}
		}
	}
	result := util.BuildSuccessResult(util.BuildDefaultPageVO(list))
	c.JSON(http.StatusOK, result)
}

func PresentOrderComplete4Si(c *gin.Context) {
	id := c.PostForm("id")
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
	serviceObj := &service.PresentOrder{}
	err = serviceObj.Complete(id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult("")
	c.JSON(http.StatusOK, result)
}

func PresentOrderPost4Ui(c *gin.Context) {
	ids := c.PostForm("ids")
	if ids == "" {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("未知礼品"))
		return
	}
	sn, err := session.GetUi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	userId, err := sn.Get(session.SessionNameUiUserId)
	if err != nil || userId == "" {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	serviceObj := &service.PresentOrder{}
	presentOrder, err := serviceObj.Random(userId, ids)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	presentId := ""
	if presentOrder != nil {
		presentId = presentOrder.PresentId
	}
	result := util.BuildSuccessResult(presentId)
	c.JSON(http.StatusOK, result)
}

func PresentOrderList(c *gin.Context) {
	argObj := &arg.PresentOrder{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	argObj.OrderBy = "number,status"
	serviceObj := &service.PresentOrder{}
	argObj.DisplayNames = []string{"id", "presentId", "userId", "number", "status"}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func PresentOrderGet(c *gin.Context) {
	serviceObj := &service.PresentOrder{}
	obj, err := service.Get(serviceObj, c.Param("id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	servicePresent := &service.Present{}
	present, err := service.Get(servicePresent, obj.(*model.PresentOrder).PresentId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	serviceUser := &service.User{}
	user, err := service.Get(serviceUser, obj.(*model.PresentOrder).UserId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(map[string]interface{}{"presentOrder": obj, "present": present, "user": user})
	//result := service.ResultGet(serviceObj, c.Param("id"))
	c.JSON(http.StatusOK, result)
}

func PresentOrderPost(c *gin.Context) {
	obj := &model.PresentOrder{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.PresentOrder{}
	result := service.ResultAdd(serviceObj, obj)
	c.JSON(http.StatusOK, result)
}

func PresentOrderPatch(c *gin.Context) {
	obj := &model.PresentOrder{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.PresentOrder{}
	result := service.ResultUpdate(serviceObj, obj, "status")
	c.JSON(http.StatusOK, result)
}

func PresentOrderDelete(c *gin.Context) {
	ids := strings.Split(c.Query("ids"), constant.Split4Id)

	serviceObj := &service.PresentOrder{}
	argObj := &arg.PresentOrder{}
	argObj.IdsIn = ids
	result := service.ResultBatchDelete(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}
