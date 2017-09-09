package handler

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"mimi/djq/model"
	"mimi/djq/util"
	"net/http"
	"strconv"
	"time"
)

func Shop4IndexList(c *gin.Context) {
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
		shop.TotalCashCouponPrice = shop.TotalCashCouponNumber + rand.Intn(10000)*rand.Intn(5)
		shop.PreImage = "preImage"
		shopList = append(shopList, shop)
	}
	c.JSON(http.StatusOK, util.BuildSuccessPageResult(targetPage, pageSize, total, shopList))
}

func Shop4Index(c *gin.Context) {
	id := c.Param("id")
	shop := &model.Shop{}
	shop.Id = id
	shop.Name = "name"
	shop.TotalCashCouponNumber = rand.Intn(100000)
	shop.TotalCashCouponPrice = shop.TotalCashCouponNumber + rand.Intn(10000)*rand.Intn(5)
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
