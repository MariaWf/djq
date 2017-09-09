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
	"mimi/djq/config"
	"path/filepath"
	"time"
	"os"
)

func AdvertisementList4Open(c *gin.Context) {
	//advertisementList := make([]*model.Advertisement, 0, 5)
	//for i := 0; i < 5; i++ {
	//	advertisement := &model.Advertisement{}
	//	advertisement.Id = "id" + strconv.Itoa(i)
	//	advertisement.Name = "name" + strconv.Itoa(i)
	//	advertisement.Image = "Image" + strconv.Itoa(i)
	//	advertisement.Link = "link" + strconv.Itoa(i)
	//	advertisementList = append(advertisementList, advertisement)
	//}
	//c.JSON(http.StatusOK, util.BuildSuccessResult(advertisementList))


	argObj := &arg.Advertisement{}
	argObj.NotIncludeHide = true
	argObj.OrderBy = "priority desc"
	serviceObj := &service.Advertisement{}
	argObj.ShowColumnNames = []string{"image", "link"}
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

func AdvertisementUploadImage(c *gin.Context) {
	file, err := c.FormFile("theFile")
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(constant.ErrUpload.Error()))
		return
	}
	fileSuffix := filepath.Ext(file.Filename)
	if fileSuffix == "" {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(constant.ErrUploadUnknownType.Error()))
		return
	}
	fileSuffix = strings.ToLower(fileSuffix)
	support := false
	for _, v := range constant.UploadImageSupport {
		if v == fileSuffix {
			support = true
		}
	}
	if !support {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(constant.ErrUploadImageSupport.Error()))
		return
	}

	newName := util.BuildUUID() + fileSuffix

	directoryHead := config.Get("uploadPath")
	if directoryHead == "" {
		directoryHead = "upload"
	}
	directoryHead, err = filepath.Abs(directoryHead)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(constant.ErrUpload.Error()))
		return
	}
	directoryHead = filepath.Join(directoryHead, "/image")
	now := time.Now()
	directory := filepath.Join("/advertisement", "/" + now.Format("200601"), "/" + strconv.Itoa(now.Day()))

	imagePath := filepath.Join(directoryHead, directory, newName)

	path := filepath.Join(directoryHead, directory)
	err = os.MkdirAll(path, 0777)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(constant.ErrUpload.Error()))
		return
	}

	err = c.SaveUploadedFile(file, imagePath)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(constant.ErrUpload.Error()))
		return
	}

	result := util.BuildSuccessResult(filepath.Join("/upload/image", directory, newName))
	c.JSON(http.StatusOK, result)
}