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

func AdvertisementList4Open(c *gin.Context) {
	argObj := &arg.Advertisement{}
	argObj.NotIncludeHide = true
	argObj.OrderBy = "priority desc"
	serviceObj := &service.Advertisement{}
	argObj.DisplayNames = []string{"image", "link"}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func AdvertisementList(c *gin.Context) {
	argObj := &arg.Advertisement{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	argObj.OrderBy = "priority desc"

	serviceObj := &service.Advertisement{}
	argObj.DisplayNames = []string{"id", "name", "image", "link", "priority", "hide", "description"}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func AdvertisementGet(c *gin.Context) {
	serviceObj := &service.Advertisement{}
	result := service.ResultGet(serviceObj, c.Param("id"))
	c.JSON(http.StatusOK, result)
}

func AdvertisementPost(c *gin.Context) {
	obj := &model.Advertisement{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.Advertisement{}
	result := service.ResultAdd(serviceObj, obj)
	c.JSON(http.StatusOK, result)
}

func AdvertisementPatch(c *gin.Context) {
	obj := &model.Advertisement{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.Advertisement{}
	result := service.ResultUpdate(serviceObj, obj, "name", "image", "link", "priority", "hide", "description")
	c.JSON(http.StatusOK, result)
}

func AdvertisementDelete(c *gin.Context) {
	ids := strings.Split(c.PostForm("ids"), constant.Split4Id)

	serviceObj := &service.Advertisement{}
	argObj := &arg.Advertisement{}
	argObj.IdsIn = ids
	result := service.ResultBatchDelete(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func AdvertisementUploadImage(c *gin.Context) {
	commonUploadImage(c, "advertisement")
}