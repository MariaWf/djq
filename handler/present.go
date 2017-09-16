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

func PresentList(c *gin.Context) {
	argObj := &arg.Present{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	argObj.OrderBy = "weight desc,name"
	serviceObj := &service.Present{}
	argObj.DisplayNames = []string{"id", "name", "image", "address", "stock", "requirement", "weight", "expiryDate", "hide"}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func PresentGet(c *gin.Context) {
	serviceObj := &service.Present{}
	result := service.ResultGet(serviceObj, c.Param("id"))
	c.JSON(http.StatusOK, result)
}

func PresentPost(c *gin.Context) {
	obj := &model.Present{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.Present{}
	result := service.ResultAdd(serviceObj, obj)
	c.JSON(http.StatusOK, result)
}

func PresentPatch(c *gin.Context) {
	obj := &model.Present{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.Present{}
	result := service.ResultUpdate(serviceObj, obj, "name", "image", "address", "stock", "requirement", "weight", "expiryDate", "hide")
	c.JSON(http.StatusOK, result)
}

func PresentDelete(c *gin.Context) {
	ids := strings.Split(c.Query("ids"), constant.Split4Id)

	serviceObj := &service.Present{}
	argObj := &arg.Present{}
	argObj.IdsIn = ids
	result := service.ResultBatchDelete(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}


func PresentUploadImage(c *gin.Context) {
	commonUploadImage(c, "present")
}