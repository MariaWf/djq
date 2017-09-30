package task

import (
	"log"
	"mimi/djq/cache"
	"mimi/djq/constant"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/service"
	"strconv"
	"time"
)

//定时自动同意未使用代金券订单的退款申请
func AgreeNotUsedRefunding() {
	err := cache.Set(cache.CacheNameAgreeingNotUsedRefunding, "false", 0)
	if err != nil {
		log.Println(err)
	}
	FixTimeIntervalCycle(AgreeNotUsedRefundingAction, time.Minute*30)
}

func AgreeNotUsedRefundingAction() {
	lock.Lock()
	counting, err := cache.Get(cache.CacheNameAgreeingNotUsedRefunding)
	if err != nil {
		lock.Unlock()
		checkErr(err)
	}
	if counting == "true" {
		lock.Unlock()
		return
	}
	counting = "true"
	err = cache.Set(cache.CacheNameAgreeingNotUsedRefunding, counting, 0)
	if err != nil {
		lock.Unlock()
		checkErr(err)
	}
	lock.Unlock()
	defer cache.Set(cache.CacheNameAgreeingNotUsedRefunding, "false", 0)

	serviceRefund := &service.Refund{}
	argRefund := &arg.Refund{}
	argRefund.StatusEqual = strconv.Itoa(constant.RefundStatusNotUsedRefunding)
	argRefund.DisplayNames = []string{"id","refundAmount"}
	argRefund.RefundOrderNumberEqual = ""
	list, err := service.Find(serviceRefund, argRefund)
	checkErr(err)
	for _, obj := range list {
		err = serviceRefund.Agree(obj.(*model.Refund).Id, "同意", obj.(*model.Refund).RefundAmount)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
