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

func ShopIntroductionImageList(c *gin.Context) {
	argObj := &arg.ShopIntroductionImage{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	argObj.OrderBy = "priority desc"

	serviceObj := &service.ShopIntroductionImage{}
	argObj.DisplayNames = []string{"id", "shopId", "priority", "contentUrl", "hide"}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func ShopIntroductionImageGet(c *gin.Context) {
	serviceObj := &service.ShopIntroductionImage{}
	result := service.ResultGet(serviceObj, c.Param("id"))
	c.JSON(http.StatusOK, result)
}

func ShopIntroductionImagePost(c *gin.Context) {
	obj := &model.ShopIntroductionImage{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.ShopIntroductionImage{}
	result := service.ResultAdd(serviceObj, obj)
	c.JSON(http.StatusOK, result)
}

func ShopIntroductionImagePatch(c *gin.Context) {
	obj := &model.ShopIntroductionImage{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.ShopIntroductionImage{}
	result := service.ResultUpdate(serviceObj, obj, "priority", "contentUrl", "hide")
	c.JSON(http.StatusOK, result)
}

func ShopIntroductionImageDelete(c *gin.Context) {
	ids := strings.Split(c.PostForm("ids"), constant.Split4Id)

	serviceObj := &service.ShopIntroductionImage{}
	argObj := &arg.ShopIntroductionImage{}
	argObj.IdsIn = ids
	result := service.ResultBatchDelete(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}