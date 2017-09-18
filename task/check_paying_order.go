package task

import (
	"time"
	"log"
	"mimi/djq/dao/arg"
	"mimi/djq/constant"
	"fmt"
	"strconv"
	"mimi/djq/service"
	"mimi/djq/util"
	"mimi/djq/model"
	"mimi/djq/wxpay"
)

func CheckPayingOrder() {
	serviceCashCouponOrder := &service.CashCouponOrder{}
	argCashCouponOrder := &arg.CashCouponOrder{}
	argCashCouponOrder.DisplayNames = []string{"payOrderNumber", "payBegin"}
	argCashCouponOrder.StatusEqual = strconv.Itoa(constant.CashCouponOrderStatusInCart)
	argCashCouponOrder.PayBeginGT = util.StringTime4DB(time.Now().Add(time.Second * 5))
	list, err := service.Find(serviceCashCouponOrder, argCashCouponOrder)
	if err != nil {
		log.Println(err)
		goto nextCycle
	}
	payOrderNumberList := make([]string, 0, len(list))
	for _, v := range list {
		payOrderNumber := v.(*model.CashCouponOrder).PayOrderNumber
		exist := false
		for _, p := range payOrderNumberList {
			if p == payOrderNumber {
				exist = true
				break;
			}
		}
		if !exist {
			payOrderNumberList = append(payOrderNumberList, payOrderNumber)
		}
		wxpay.OrderQuery(payOrderNumber)
	}

	nextCycle:
	time.Sleep(time.Second * 5)
	go CheckPayingOrder()
}
