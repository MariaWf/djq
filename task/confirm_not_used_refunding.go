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

func ConfirmNotUsedRefunding() {
	serviceRefund := &service.Refund{}
	argRefund := &arg.Refund{}
	argRefund.StatusEqual = strconv.Itoa(constant.RefundStatusNotUsedRefunding)
	argRefund.DisplayNames = []string{"id"}
	list, err := service.Find(serviceRefund, argRefund)
	if err != nil {
		log.Println(err)
		goto nextCycle
	}
	for _, obj := range list {
		err = serviceRefund.Confirm(obj.(*model.Refund).Id)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	nextCycle:
	time.Sleep(time.Minute * 30)
	go ConfirmNotUsedRefunding()
}