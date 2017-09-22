package task

import (
	"time"
	"mimi/djq/service"
	"mimi/djq/dao/arg"
	"strconv"
	"mimi/djq/constant"
	"log"
	"mimi/djq/model"
)

//定时自动同意未使用代金券订单的退款申请
func AgreeNotUsedRefunding() {
	serviceRefund := &service.Refund{}
	argRefund := &arg.Refund{}
	argRefund.StatusEqual = strconv.Itoa(constant.RefundStatusNotUsedRefunding)
	argRefund.DisplayNames = []string{"id"}
	argRefund.RefundOrderNumberEqual = ""
	list, err := service.Find(serviceRefund, argRefund)
	if err != nil {
		log.Println(err)
		goto nextCycle
	}
	for _, obj := range list {
		err = serviceRefund.Agree(obj.(*model.Refund).Id,"同意",obj.(*model.Refund).RefundAmount)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	nextCycle:
	time.Sleep(time.Minute * 30)
	go AgreeNotUsedRefunding()
}