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

func CashCouponList(c *gin.Context) {
	argObj := &arg.CashCoupon{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	argObj.OrderBy = "priority desc,name"

	serviceObj := &service.CashCoupon{}
	argObj.DisplayNames = []string{"id", "shopId", "name", "preImage", "price", "discountAmount", "expiryDate", "expired", "hide", "priority"}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func CashCouponGet(c *gin.Context) {
	serviceObj := &service.CashCoupon{}
	result := service.ResultGet(serviceObj, c.Param("id"))
	c.JSON(http.StatusOK, result)
}

func CashCouponPost(c *gin.Context) {
	obj := &model.CashCoupon{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.CashCoupon{}
	result := service.ResultAdd(serviceObj, obj)
	c.JSON(http.StatusOK, result)
}

func CashCouponPatch(c *gin.Context) {
	obj := &model.CashCoupon{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.CashCoupon{}
	result := service.ResultUpdate(serviceObj, obj, "name", "preImage", "price", "discountAmount", "expiryDate", "expired", "hide", "priority")
	c.JSON(http.StatusOK, result)
}

func CashCouponDelete(c *gin.Context) {
	ids := strings.Split(c.Query("ids"), constant.Split4Id)

	serviceObj := &service.CashCoupon{}
	argObj := &arg.CashCoupon{}
	argObj.IdsIn = ids
	result := service.ResultBatchDelete(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func CashCouponUploadImage(c *gin.Context) {
	commentUploadImage(c, "cashCoupon")
}
