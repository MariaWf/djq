package handler

import (
	"github.com/gin-gonic/gin"
	"mimi/djq/model"
	"mimi/djq/util"
	"net/http"
	"strings"
	"log"
	"mimi/djq/constant"
	"mimi/djq/service"
	"mimi/djq/dao/arg"
)

func ShopList4Open(c *gin.Context) {
	argObj := &arg.Shop{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	argObj.OrderBy = "priority desc"
	argObj.NotIncludeHide = true
	serviceObj := &service.Shop{}
	argObj.DisplayNames = []string{"id", "name", "preImage", "totalCashCouponNumber", "totalCashCouponPrice", "priority"}
	argObj.PageSize = 5
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func ShopGet4Open(c *gin.Context) {
	serviceObj := &service.Shop{}
	obj, err := serviceObj.Get(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	serviceShopIntroductionImage := &service.ShopIntroductionImage{}
	argShopIntroductionImage := &arg.ShopIntroductionImage{}
	argShopIntroductionImage.ShopIdEqual = obj.Id
	argShopIntroductionImage.NotIncludeHide = true
	argShopIntroductionImage.OrderBy = "priority desc"
	argShopIntroductionImage.DisplayNames = []string{"contentUrl"}
	shopIntroductionImageList, err := service.Find(serviceShopIntroductionImage, argShopIntroductionImage)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	obj.SetShopIntroductionImageListFromInterfaceArr(shopIntroductionImageList)

	serviceCashCoupon := &service.CashCoupon{}
	argCashCoupon := &arg.CashCoupon{}
	argCashCoupon.ShopIdEqual = obj.Id
	argCashCoupon.NotIncludeHide = true
	argCashCoupon.BeforeExpiryDate = true
	argCashCoupon.OrderBy = "priority desc"
	argCashCoupon.DisplayNames = []string{"id", "name", "preImage", "expiryDate"}
	cashCouponList, err := service.Find(serviceCashCoupon, argCashCoupon)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	obj.SetCashCouponListFromInterfaceArr(cashCouponList)
	result := util.BuildSuccessResult(obj)
	c.JSON(http.StatusOK, result)
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
	argObj.DisplayNames = []string{"id", "name", "logo", "preImage", "totalCashCouponNumber", "totalCashCouponPrice", "introduction", "address", "priority", "hide"}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func ShopGet(c *gin.Context) {
	serviceObj := &service.Shop{}
	obj, err := serviceObj.Get(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(obj)
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
	ids := strings.Split(c.Query("ids"), constant.Split4Id)

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

func ShopUploadPreImage(c *gin.Context) {
	commentUploadImage(c, "shop/preImage")
}

func ShopUploadLogo(c *gin.Context) {
	commentUploadImage(c, "shop/logo")
}