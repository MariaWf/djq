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
	"strconv"
)

func AdvertisementList4Index(c *gin.Context) {
	advertisementList := make([]*model.Advertisement, 0, 5)
	for i := 0; i < 5; i++ {
		advertisement := &model.Advertisement{}
		advertisement.Id = "id" + strconv.Itoa(i)
		advertisement.Name = "name" + strconv.Itoa(i)
		advertisement.Image = "Image" + strconv.Itoa(i)
		advertisement.Link = "link" + strconv.Itoa(i)
		advertisementList = append(advertisementList, advertisement)
	}
	c.JSON(http.StatusOK, util.BuildSuccessResult(advertisementList))
}

func AdvertisementList(c *gin.Context) {
	argObj := &arg.Advertisement{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.Advertisement{}
	argObj.ShowColumnNames = []string{"id", "name", "image", "link", "priority", "hide", "description"}
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