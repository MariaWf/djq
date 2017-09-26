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
	"strconv"
	"time"
)

//每1分钟检测一次，所有离开始支付时间超过1分钟的购物车代金券，
// 如果支付结果为已支付，确认该订单，
// 其他情况若且离开始时间超过10分钟，关闭订单
func CheckPayingOrder() {
	err := cache.Set(cache.CacheNameCheckingPayingOrder, "false", 0)
	if err != nil {
		log.Println(err)
	}
	FixTimeIntervalCycle(CheckPayingOrderAction, time.Minute*1)
}

func CheckPayingOrderAction() {
	lock.Lock()
	counting, err := cache.Get(cache.CacheNameCheckingPayingOrder)
	if err != nil {
		lock.Unlock()
		checkErr(err)
	}
	if counting == "true" {
		lock.Unlock()
		return
	}
	counting = "true"
	err = cache.Set(cache.CacheNameCheckingPayingOrder, counting, 0)
	if err != nil {
		lock.Unlock()
		checkErr(err)
	}
	lock.Unlock()
	defer cache.Set(cache.CacheNameCheckingPayingOrder, "false", 0)

	serviceCashCouponOrder := &service.CashCouponOrder{}
	argCashCouponOrder := &arg.CashCouponOrder{}
	argCashCouponOrder.DisplayNames = []string{"payOrderNumber", "payBegin"}
	argCashCouponOrder.StatusEqual = strconv.Itoa(constant.CashCouponOrderStatusInCart)
	argCashCouponOrder.PayBeginLT = util.StringTime4DB(time.Now().Add(time.Minute * -1))
	list, err := service.Find(serviceCashCouponOrder, argCashCouponOrder)
	var payOrderNumber string
	serviceObj := &service.CashCouponOrder{}
	checkErr(err)
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
}
