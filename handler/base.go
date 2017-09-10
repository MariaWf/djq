package handler

import (
	"github.com/pkg/errors"
	"path/filepath"
	"net/http"
	"time"
	"strconv"
	"os"
	"strings"
	"github.com/gin-gonic/gin"
	"mimi/djq/util"
	"log"
	"mimi/djq/constant"
	"mimi/djq/config"
)

var ErrParamException = errors.New("参数异常")

func commonUploadImage(c *gin.Context, typeHead string) {
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
	directory := filepath.Join(typeHead, "/" + now.Format("200601"), "/" + strconv.Itoa(now.Day()))

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