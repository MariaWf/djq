package handler

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"mimi/djq/model"
	"mimi/djq/util"
	"net/http"
	"strconv"
	"time"
	"path/filepath"
	"os"
	"strings"
	"log"
	"mimi/djq/config"
	"mimi/djq/constant"
	"mimi/djq/service"
	"mimi/djq/dao/arg"
)

func ShopList4Open(c *gin.Context) {
	targetPageStr := c.DefaultQuery("targetPage", strconv.Itoa(util.BeginPage))
	targetPage, err := strconv.Atoi(targetPageStr)
	if err != nil {
		panic(err)
	}
	pageSize := 5
	total := 123

	shopList := make([]*model.Shop, 0, pageSize)
	start := util.ComputePageStart(targetPage, pageSize)
	end := start + pageSize
	if end > total {
		end = total
	}
	for i := start; i < end; i++ {
		shop := &model.Shop{}
		shop.Id = "id" + strconv.Itoa(i)
		shop.Name = "name" + strconv.Itoa(i)
		shop.TotalCashCouponNumber = rand.Intn(100000)
		shop.TotalCashCouponPrice = shop.TotalCashCouponNumber + rand.Intn(10000) * rand.Intn(5)
		shop.PreImage = "preImage"
		shopList = append(shopList, shop)
	}
	c.JSON(http.StatusOK, util.BuildSuccessPageResult(targetPage, pageSize, total, shopList))
}

func ShopGet4Open(c *gin.Context) {
	id := c.Param("id")
	shop := &model.Shop{}
	shop.Id = id
	shop.Name = "name"
	shop.TotalCashCouponNumber = rand.Intn(100000)
	shop.TotalCashCouponPrice = shop.TotalCashCouponNumber + rand.Intn(10000) * rand.Intn(5)
	shop.PreImage = "preImage"
	shopIntroductionImageList := make([]*model.ShopIntroductionImage, 0, 5)
	for i := 0; i < 5; i++ {
		shopIntroductionImage := &model.ShopIntroductionImage{}
		shopIntroductionImage.Id = "id" + strconv.Itoa(i)
		shopIntroductionImage.ContentUrl = "contentUrl" + strconv.Itoa(i)
		shopIntroductionImageList = append(shopIntroductionImageList, shopIntroductionImage)
	}
	shop.ShopIntroductionImageList = shopIntroductionImageList

	cashCouponList := make([]*model.CashCoupon, 0, 5)
	for i := 0; i < 5; i++ {
		cashCoupon := &model.CashCoupon{}
		cashCoupon.Id = "id" + strconv.Itoa(i)
		cashCoupon.ExpiryDate = util.JSONTime(time.Now())
		cashCoupon.Name = "name" + strconv.Itoa(i)
		cashCoupon.PreImage = "preImage"
		cashCouponList = append(cashCouponList, cashCoupon)
	}
	shop.CashCouponList = cashCouponList
	c.JSON(http.StatusOK, util.BuildSuccessResult(shop))
}

func ShopList(c *gin.Context) {
	argObj := &arg.Shop{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	argObj.OrderBy = "priority desc"

	serviceObj := &service.Shop{}
	argObj.ShowColumnNames = []string{"id", "name", "logo", "preImage", "totalCashCouponNumber", "totalCashCouponPrice", "introduction", "address", "priority", "hide"}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func ShopGet(c *gin.Context) {
	serviceObj := &service.Shop{}
	result := service.ResultGet(serviceObj, c.Param("id"))
	c.JSON(http.StatusOK, result)
}

func ShopPost(c *gin.Context) {
	obj := &model.Shop{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	shopClassificationIds := c.PostForm("shopClassificationIds")
	if shopClassificationIds != "" {
		shopClassificationIdList := strings.Split(shopClassificationIds, ",")
		if shopClassificationIdList != nil && len(shopClassificationIdList) != 0 {
			obj.ShopClassificationList = make([]*model.ShopClassification, len(shopClassificationIdList), len(shopClassificationIdList))
			for i, shopClassificationId := range shopClassificationIdList {
				obj.ShopClassificationList[i] = &model.ShopClassification{Id: shopClassificationId}
			}
		}
	}

	serviceObj := &service.Shop{}
	obj, err = serviceObj.Add(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(obj.Id)
	c.JSON(http.StatusOK, result)
}

func ShopPatch(c *gin.Context) {
	obj := &model.Shop{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	shopClassificationIds := c.PostForm("shopClassificationIds")
	if shopClassificationIds != "" {
		shopClassificationIdList := strings.Split(shopClassificationIds, ",")
		if shopClassificationIdList != nil && len(shopClassificationIdList) != 0 {
			obj.ShopClassificationList = make([]*model.ShopClassification, len(shopClassificationIdList), len(shopClassificationIdList))
			for i, shopClassificationId := range shopClassificationIdList {
				obj.ShopClassificationList[i] = &model.ShopClassification{Id: shopClassificationId}
			}
		}
	}

	serviceObj := &service.Shop{}
	obj, err = serviceObj.Update(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(obj.Id)
	c.JSON(http.StatusOK, result)
	//result := service.ResultUpdate(serviceObj, obj, "name", "logo", "preImage", "totalCashCouponNumber", "totalCashCouponPrice", "introduction", "address", "priority", "hide")
	//c.JSON(http.StatusOK, result)
}

func ShopDelete(c *gin.Context) {
	ids := strings.Split(c.PostForm("ids"), constant.Split4Id)

	serviceObj := &service.Shop{}
	count, err := serviceObj.Delete(ids...)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(count)
	c.JSON(http.StatusOK, result)
}

func ShopUploadImage(c *gin.Context) {
	commonUploadImage(c, "shop")
}