package task

import (
	"log"
	"mimi/djq/cache"
	"mimi/djq/constant"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/service"
	"mimi/djq/util"
	"mimi/djq/wxpay"
	"time"
)

//每1分钟检测一次，所有离开始退款时间超过1分钟的执行中退款，
// 如果支付结果为已退款(SUCCESS)，确认该订单，
//如果是退款处理中（PROCESSING），继续等待，
// 其他情况，关闭订单
func CheckRefundingOrder() {
	err := cache.Set(cache.CacheNameCheckingRefundingOrder, "false", 0)
	if err != nil {
		log.Println(err)
	}
	FixTimeIntervalCycle(CheckRefundingOrderAction, time.Minute*1)
}

func CheckRefundingOrderAction() {
	lock.Lock()
	counting, err := cache.Get(cache.CacheNameCheckingRefundingOrder)
	if err != nil {
		lock.Unlock()
		checkErr(err)
	}
	if counting == "true" {
		lock.Unlock()
		return
	}
	counting = "true"
	err = cache.Set(cache.CacheNameCheckingRefundingOrder, counting, 0)
	if err != nil {
		lock.Unlock()
		checkErr(err)
	}
	lock.Unlock()
	defer cache.Set(cache.CacheNameCheckingRefundingOrder, "false", 0)

	serviceRefund := &service.Refund{}
	argRefund := &arg.Refund{}
	argRefund.DisplayNames = []string{"refundOrderNumber", "refundBegin"}
	argRefund.StatusIn = []int{constant.RefundStatusNotUsedRefunding, constant.RefundStatusUsedRefunding}
	argRefund.RefundBeginLT = util.StringTime4DB(time.Now().Add(time.Minute * -1))
	list, err := service.Find(serviceRefund, argRefund)
	var refundOrderNumber string
	serviceObj := &service.Refund{}
	checkErr(err)
	for _, obj := range list {
		refundOrderNumber = obj.(*model.Refund).RefundOrderNumber
		refundState, _, err := wxpay.RefundQueryResult(refundOrderNumber)
		if err != nil {
			log.Println(err)
			continue
		}
		//var idListStr string
		switch refundState {
		case "SUCCESS":
			err = serviceObj.ConfirmByRefundOrderNumber(refundOrderNumber)
			if err != nil {
				log.Println(err)
				continue
			}
		case "PROCESSING":
			continue
		default:
			err = serviceObj.FailCloseByRefundOrderNumber(refundOrderNumber)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
}
