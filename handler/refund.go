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

func RefundList(c *gin.Context) {
	argObj := &arg.Refund{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	argObj.OrderBy = "status,cash_coupon_order_id,refund_amount"

	serviceObj := &service.Refund{}
	argObj.DisplayNames = []string{"id", "cashCouponOrderId", "evidence", "reason", "comment", "refundAmount", "status"}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func RefundGet(c *gin.Context) {
	serviceObj := &service.Refund{}
	obj, err := service.Get(serviceObj, c.Param("id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	serviceCashCouponOrder := &service.CashCouponOrder{}
	cashCouponOrder, err := service.Get(serviceCashCouponOrder, obj.(*model.Refund).CashCouponOrderId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	serviceCashCoupon := &service.CashCoupon{}
	cashCoupon, err := service.Get(serviceCashCoupon, cashCouponOrder.(*model.CashCouponOrder).CashCouponId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(map[string]interface{}{"refund":obj,"cashCouponOrder":cashCouponOrder,"cashCoupon":cashCoupon})
	//result := service.ResultGet(serviceObj, c.Param("id"))
	c.JSON(http.StatusOK, result)
}

func RefundPost(c *gin.Context) {
	obj := &model.Refund{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.Refund{}
	result := service.ResultAdd(serviceObj, obj)
	c.JSON(http.StatusOK, result)
}

func RefundPatch(c *gin.Context) {
	obj := &model.Refund{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.Refund{}
	result := service.ResultUpdate(serviceObj, obj, "evidence", "reason", "comment", "refundAmount", "status")
	c.JSON(http.StatusOK, result)
}

func RefundDelete(c *gin.Context) {
	ids := strings.Split(c.Query("ids"), constant.Split4Id)

	serviceObj := &service.Refund{}
	argObj := &arg.Refund{}
	argObj.IdsIn = ids
	result := service.ResultBatchDelete(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}