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

func ShopClassificationList4Open(c *gin.Context) {
	argObj := &arg.ShopClassification{}
	argObj.NotIncludeHide = true
	argObj.OrderBy = "priority desc"
	serviceObj := &service.ShopClassification{}
	argObj.DisplayNames = []string{"id", "name"}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func ShopClassificationList(c *gin.Context) {
	argObj := &arg.ShopClassification{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	argObj.OrderBy = "priority desc"

	serviceObj := &service.ShopClassification{}
	argObj.DisplayNames = []string{"id", "name", "priority", "hide","description"}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func ShopClassificationGet(c *gin.Context) {
	serviceObj := &service.ShopClassification{}
	result := service.ResultGet(serviceObj, c.Param("id"))
	c.JSON(http.StatusOK, result)
}

func ShopClassificationPost(c *gin.Context) {
	obj := &model.ShopClassification{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.ShopClassification{}
	result := service.ResultAdd(serviceObj, obj)
	c.JSON(http.StatusOK, result)
}

func ShopClassificationPatch(c *gin.Context) {
	obj := &model.ShopClassification{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.ShopClassification{}
	result := service.ResultUpdate(serviceObj, obj, "name", "priority", "hide","description")
	c.JSON(http.StatusOK, result)
}

func ShopClassificationDelete(c *gin.Context) {
	ids := strings.Split(c.Query("ids"), constant.Split4Id)

	serviceObj := &service.ShopClassification{}
	count, err := serviceObj.Delete(ids...)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(count)
	c.JSON(http.StatusOK, result)
}
