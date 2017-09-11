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

func PromotionalPartnerList(c *gin.Context) {
	argObj := &arg.PromotionalPartner{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	argObj.OrderBy = "name"

	serviceObj := &service.PromotionalPartner{}
	argObj.DisplayNames = []string{"id", "name", "description", "totalUser", "totalPrice", "totalPay"}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func PromotionalPartnerGet(c *gin.Context) {
	serviceObj := &service.PromotionalPartner{}
	result := service.ResultGet(serviceObj, c.Param("id"))
	c.JSON(http.StatusOK, result)
}

func PromotionalPartnerPost(c *gin.Context) {
	obj := &model.PromotionalPartner{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.PromotionalPartner{}
	result := service.ResultAdd(serviceObj, obj)
	c.JSON(http.StatusOK, result)
}

func PromotionalPartnerPatch(c *gin.Context) {
	obj := &model.PromotionalPartner{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.PromotionalPartner{}
	result := service.ResultUpdate(serviceObj, obj, "name", "description", "totalUser", "totalPrice", "totalPay")
	c.JSON(http.StatusOK, result)
}

func PromotionalPartnerDelete(c *gin.Context) {
	ids := strings.Split(c.Query("ids"), constant.Split4Id)

	serviceObj := &service.PromotionalPartner{}
	argObj := &arg.PromotionalPartner{}
	argObj.IdsIn = ids
	count, err := serviceObj.Delete(ids...)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(count)
	c.JSON(http.StatusOK, result)
}