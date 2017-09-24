package task

import (
	"mimi/djq/dao/arg"
	"mimi/djq/service"
	"mimi/djq/model"
	"mimi/djq/constant"
	"mimi/djq/cache"
	"strconv"
	"log"
)

//每天凌晨2点统计合作伙伴数据
func CountForPromotionalPartner() {
	err := cache.Set(cache.CacheNamePromotionalPartnerCounting, "false", 0)
	if err != nil {
		log.Println(err)
	}
	FixTimeCycle(CountForPromotionalPartnerAction, 2, 0, 0)
}

func CountForPromotionalPartnerAction() {
	lock.Lock()
	counting, err := cache.Get(cache.CacheNamePromotionalPartnerCounting)
	if err != nil {
		lock.Unlock()
		checkErr(err)
	}
	if counting == "true" {
		lock.Unlock()
		return
	}
	counting = "true"
	err = cache.Set(cache.CacheNamePromotionalPartnerCounting, counting, 0)
	if err != nil {
		lock.Unlock()
		checkErr(err)
	}
	lock.Unlock()
	defer cache.Set(cache.CacheNamePromotionalPartnerCounting, "false", 0)

	servicePromotionalPartner := &service.PromotionalPartner{}
	argPromotionalPartner := &arg.PromotionalPartner{}
	promotionalPartnerListO, err := service.Find(servicePromotionalPartner, argPromotionalPartner)
	checkErr(err)
	serviceUser := &service.User{}
	serviceCashCouponOrder := &service.CashCouponOrder{}
	rateStr, err := cache.Get(cache.CacheNamePromotionalPartnerRate)
	checkErr(err)
	if rateStr == "" {
		rateStr = "3"
	}
	rate, err := strconv.Atoi(rateStr)
	checkErr(err)
	for _, v := range promotionalPartnerListO {
		promotionalPartner := v.(*model.PromotionalPartner)
		argUser := &arg.User{}
		argUser.PromotionalPartnerIdEqual = promotionalPartner.Id
		userListO, err := service.Find(serviceUser, argUser)
		checkErr(err)
		totalUser := len(userListO)
		totalPrice := 0
		if totalUser > 0 {
			userIds := make([]string, totalUser, totalUser)
			for i, v2 := range userListO {
				user := v2.(*model.User)
				userIds[i] = user.Id
			}
			argCashCouponOrder := &arg.CashCouponOrder{}
			argCashCouponOrder.UserIdsIn = userIds
			argCashCouponOrder.StatusIn = []int{constant.CashCouponOrderStatusUsed, constant.CashCouponOrderStatusUsedRefunded}
			userOrderListO, err := service.Find(serviceCashCouponOrder, argCashCouponOrder)
			checkErr(err)
			for _, v3 := range userOrderListO {
				userOrder := v3.(*model.CashCouponOrder)
				money := userOrder.Price - userOrder.RefundAmount
				if money > 0 {
					totalPrice += money
				}
			}
		}
		totalPrice = int(float32(totalPrice) * float32(rate) / float32(100))
		if promotionalPartner.TotalUser != totalUser || promotionalPartner.TotalPrice != totalPrice {
			promotionalPartner.TotalUser = totalUser
			promotionalPartner.TotalPrice = totalPrice
			_, err = service.Update(servicePromotionalPartner, promotionalPartner, "totalUser", "totalPrice")
			checkErr(err)
		}
	}
}