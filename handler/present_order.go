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
)

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
	result := util.BuildSuccessResult(map[string]interface{}{"presentOrder":obj,"present":present,"user":user})
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
