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
	"mimi/djq/session"
	"strconv"
)

func RefundList4Ui(c *gin.Context) {
	sn, err := session.GetUi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	userId, err := sn.Get(session.SessionNameUiUserId)
	if err != nil || userId == ""{
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	serviceCashCouponOrder := &service.CashCouponOrder{}
	argCashCouponOrder := &arg.CashCouponOrder{}
	argCashCouponOrder.UserIdEqual = userId
	argCashCouponOrder.DisplayNames = []string{"id"}
	cashCouponOrderList, err := service.Find(serviceCashCouponOrder, argCashCouponOrder)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	var result *util.ResultVO
	targetPageStr := c.Query("targetPage")
	targetPage, err := strconv.Atoi(targetPageStr)
	if err != nil {
		targetPage = util.BeginPage
	}
	pageSizeStr := c.Query("pageSize")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 0
	}
	if len(cashCouponOrderList) > 0 {
		cashCouponOrderIds := make([]string, len(cashCouponOrderList), len(cashCouponOrderList))
		for i, v := range cashCouponOrderList {
			cashCouponOrderIds[i] = v.(*model.CashCouponOrder).Id
		}
		argObj := &arg.Refund{}
		argObj.TargetPage = targetPage
		argObj.PageSize = pageSize
		argObj.CashCouponOrderIdsIn = cashCouponOrderIds
		argObj.OrderBy = "status,cash_coupon_order_id,refund_amount"

		serviceObj := &service.Refund{}
		argObj.DisplayNames = []string{"id", "cashCouponOrderId", "evidence", "reason", "comment", "refundAmount", "status"}
		result = service.ResultList(serviceObj, argObj)
	} else {
		result = util.BuildSuccessResult(util.BuildPageVO(targetPage, pageSize, 0, nil))
	}
	c.JSON(http.StatusOK, result)

}

func RefundPost4Ui(c *gin.Context) {
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
func RefundCancel4Ui(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("未知退款申请"))
		return
	}

	serviceObj := &service.Refund{}
	status, err := serviceObj.Cancel(id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	result := util.BuildSuccessResult(status)
	c.JSON(http.StatusOK, result)
}

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
	result := util.BuildSuccessResult(map[string]interface{}{"refund":obj, "cashCouponOrder":cashCouponOrder, "cashCoupon":cashCoupon})
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



func RefundUploadEvidence(c *gin.Context) {
	commonUploadImage(c, "evidence")
}