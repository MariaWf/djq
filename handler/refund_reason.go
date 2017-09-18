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

func RefundReasonList4Ui(c *gin.Context) {
	argObj := &arg.RefundReason{}
	argObj.NotIncludeHide = true
	argObj.OrderBy = "priority desc"

	serviceObj := &service.RefundReason{}
	argObj.DisplayNames = []string{"id", "priority", "hide", "description"}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func RefundReasonList(c *gin.Context) {
	argObj := &arg.RefundReason{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	argObj.OrderBy = "priority desc"

	serviceObj := &service.RefundReason{}
	argObj.DisplayNames = []string{"id", "priority", "hide", "description"}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func RefundReasonGet(c *gin.Context) {
	serviceObj := &service.RefundReason{}
	result := service.ResultGet(serviceObj, c.Param("id"))
	c.JSON(http.StatusOK, result)
}

func RefundReasonPost(c *gin.Context) {
	obj := &model.RefundReason{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.RefundReason{}
	result := service.ResultAdd(serviceObj, obj)
	c.JSON(http.StatusOK, result)
}

func RefundReasonPatch(c *gin.Context) {
	obj := &model.RefundReason{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.RefundReason{}
	result := service.ResultUpdate(serviceObj, obj, "priority", "hide", "description")
	c.JSON(http.StatusOK, result)
}

func RefundReasonDelete(c *gin.Context) {
	ids := strings.Split(c.Query("ids"), constant.Split4Id)

	serviceObj := &service.RefundReason{}
	argObj := &arg.RefundReason{}
	argObj.IdsIn = ids
	result := service.ResultBatchDelete(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}