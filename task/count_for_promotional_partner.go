package task

import (
	"mimi/djq/dao/arg"
	"mimi/djq/service"
	"mimi/djq/model"
	"mimi/djq/constant"
)

//每天凌晨2点统计合作伙伴数据
func CountForPromotionalPartner() {
	FixTimeCycle(CountForPromotionalPartnerAction, 2, 0, 0)
}

func CountForPromotionalPartnerAction() {
	servicePromotionalPartner := &service.PromotionalPartner{}
	argPromotionalPartner := &arg.PromotionalPartner{}
	promotionalPartnerListO, err := service.Find(servicePromotionalPartner, argPromotionalPartner)
	checkErr(err)
	serviceUser := &service.User{}
	serviceCashCouponOrder := &service.CashCouponOrder{}
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
		if promotionalPartner.TotalUser != totalUser || promotionalPartner.TotalPrice != totalPrice {
			promotionalPartner.TotalUser = totalUser
			promotionalPartner.TotalPrice = totalPrice
			_, err = service.Update(servicePromotionalPartner, promotionalPartner, "totalUser", "totalPrice")
			checkErr(err)
		}
	}
}