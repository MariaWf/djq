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

func CashCouponOrderList(c *gin.Context) {
	argObj := &arg.CashCouponOrder{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	argObj.OrderBy = "number"

	serviceObj := &service.CashCouponOrder{}
	argObj.DisplayNames = []string{"id", "userId", "cashCouponId", "price", "refundAmount", "payOrderNumber", "number", "status"}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func CashCouponOrderGet(c *gin.Context) {
	serviceObj := &service.CashCouponOrder{}
	result := service.ResultGet(serviceObj, c.Param("id"))
	c.JSON(http.StatusOK, result)
}

func CashCouponOrderPost(c *gin.Context) {
	obj := &model.CashCouponOrder{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.CashCouponOrder{}
	result := service.ResultAdd(serviceObj, obj)
	c.JSON(http.StatusOK, result)
}

func CashCouponOrderPatch(c *gin.Context) {
	obj := &model.CashCouponOrder{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.CashCouponOrder{}
	result := service.ResultUpdate(serviceObj, obj, "status")
	c.JSON(http.StatusOK, result)
}

func CashCouponOrderDelete(c *gin.Context) {
	ids := strings.Split(c.Query("ids"), constant.Split4Id)

	serviceObj := &service.CashCouponOrder{}
	argObj := &arg.CashCouponOrder{}
	argObj.IdsIn = ids
	result := service.ResultBatchDelete(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}