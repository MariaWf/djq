package task

import (
	"time"
	"log"
	"mimi/djq/dao/arg"
	"mimi/djq/constant"
	"strconv"
	"mimi/djq/service"
	"mimi/djq/util"
	"mimi/djq/model"
	"mimi/djq/wxpay"
)

//每1分钟检测一次，所有离开始支付时间超过1分钟的购物车代金券，
// 如果支付结果为已支付，确认该订单，
// 其他情况若且离开始时间超过10分钟，关闭订单
func CheckPayingOrder() {
	serviceCashCouponOrder := &service.CashCouponOrder{}
	argCashCouponOrder := &arg.CashCouponOrder{}
	argCashCouponOrder.DisplayNames = []string{"payOrderNumber", "payBegin"}
	argCashCouponOrder.StatusEqual = strconv.Itoa(constant.CashCouponOrderStatusInCart)
	argCashCouponOrder.PayBeginLT = util.StringTime4DB(time.Now().Add(time.Minute * -1))
	list, err := service.Find(serviceCashCouponOrder, argCashCouponOrder)
	var payOrderNumber string
	serviceObj := &service.CashCouponOrder{}
	if err != nil {
		log.Println(err)
		goto nextCycle
	}
	for _, obj := range list {
		payOrderNumber = obj.(*model.CashCouponOrder).PayOrderNumber
		tradeState, totalFee, err := wxpay.OrderQueryResult(payOrderNumber)
		if err != nil {
			log.Println(err)
			continue
		}
		//var idListStr string
		switch tradeState {
		case "SUCCESS":
			_, err = serviceObj.ConfirmOrder(payOrderNumber, totalFee)
			if err != nil {
				log.Println(err)
				continue
			}
		default:
			t, err := util.ParseTimeFromDB(obj.(*model.CashCouponOrder).PayBegin)
			if err != nil {
				log.Println(err)
				continue
			}
			if t.Add(time.Minute * 10).Before(time.Now()) {
				_, err := wxpay.CloseOrderResult(payOrderNumber)
				if err != nil {
					log.Println(err)
					continue
				}
				_, err = serviceObj.CancelOrder(payOrderNumber)
				//cache.Set(cache.CacheNameWxpayPayOrderNumberCancel + payOrderNumber, idListStr, time.Hour * 24 * 7)
			}
		}
	}
	nextCycle:
	time.Sleep(time.Minute * 1)
	go CheckPayingOrder()
}
