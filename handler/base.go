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
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var ErrParamException = errors.New("参数异常")

var ErrUnknown = errors.New("操作失败，请稍后重试")

func commentUploadImage(c *gin.Context, typeHead string) {
	aliyunUpload(c, typeHead)
}

func localUpload(c *gin.Context, typeHead string) {
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

	err = os.MkdirAll(filepath.Join(directoryHead, directory), 0777)
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

	result := util.BuildSuccessResult(util.PathAppend(config.Get("server_root_url"), filepath.Join("/upload/image", directory, newName)))
	c.JSON(http.StatusOK, result)
}

func aliyunUpload(c *gin.Context, typeHead string) {
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

	src, err := file.Open()
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(constant.ErrUpload.Error()))
		return
	}
	defer src.Close()

	client, err := oss.New(config.Get("aliyun_oss_end_point"), config.Get("aliyun_access_key_id"), config.Get("aliyun_access_secret"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(constant.ErrUpload.Error()))
		return
		// HandleError(err)
	}

	bucket, err := client.Bucket(config.Get("aliyun_oss_bucket"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(constant.ErrUpload.Error()))
		return
		// HandleError(err)
	}
	objectKey := util.PathAppend(constant.AliyunOssUploadImagePath, typeHead, newName)
	err = bucket.PutObject(objectKey, src)
	//err = bucket.PutObjectFromFile("my-object", "LocalFile")
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(constant.ErrUpload.Error()))
		return
	}
	result := util.BuildSuccessResult(util.PathAppend("http://", config.Get("static_resource_domain"), objectKey))
	c.JSON(http.StatusOK, result)
}