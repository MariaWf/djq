package task

import (
	"mimi/djq/service"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/cache"
	"mimi/djq/constant"
	"log"
)

//每天凌晨3点统计商家代金券消费数据
func CountCashCoupon() {
	err := cache.Set(cache.CacheNameCashCouponOrderCounting, "false", 0)
	if err != nil {
		log.Println(err)
	}
	FixTimeCycle(CountCashCouponAction, 3, 0, 0)
}

func CountCashCouponAction() {
	lock.Lock()
	counting, err := cache.Get(cache.CacheNameCashCouponOrderCounting)
	if err != nil {
		lock.Unlock()
		checkErr(err)
	}
	if counting == "true" {
		lock.Unlock()
		return
	}
	counting = "true"
	err = cache.Set(cache.CacheNameCashCouponOrderCounting, counting, 0)
	if err != nil {
		lock.Unlock()
		checkErr(err)
	}
	lock.Unlock()
	defer cache.Set(cache.CacheNameCashCouponOrderCounting, "false", 0)

	serviceShop := &service.Shop{}
	argShop := &arg.Shop{}
	shopListO, err := service.Find(serviceShop, argShop)
	checkErr(err)
	serviceCashCoupon := &service.CashCoupon{}
	serviceCashCouponOrder := &service.CashCouponOrder{}
	globalTotalCashCouponNumber := 0
	globalTotalCashCouponPrice := 0
	for _, v := range shopListO {
		shop := v.(*model.Shop)
		argCashCoupon := &arg.CashCoupon{}
		argCashCoupon.ShopIdEqual = shop.Id
		cashCouponListO, err := service.Find(serviceCashCoupon, argCashCoupon)
		checkErr(err)
		size := len(cashCouponListO)
		totalCashCouponNumber := 0
		totalCashCouponPrice := 0
		if size > 0 {
			cashCouponIds := make([]string, size, size)
			for i, v2 := range cashCouponListO {
				cashCoupon := v2.(*model.CashCoupon)
				cashCouponIds[i] = cashCoupon.Id
			}
			argCashCouponOrder := &arg.CashCouponOrder{}
			argCashCouponOrder.CashCouponIdsIn = cashCouponIds
			argCashCouponOrder.StatusIn = []int{constant.CashCouponOrderStatusUsed, constant.CashCouponOrderStatusUsedRefunded}
			cashCouponOrderListO, err := service.Find(serviceCashCouponOrder, argCashCouponOrder)
			checkErr(err)
			for _, v3 := range cashCouponOrderListO {
				cashCouponOrder := v3.(*model.CashCouponOrder)
				money := cashCouponOrder.Price - cashCouponOrder.RefundAmount
				if money > 0 {
					totalCashCouponNumber++
					for _, v2 := range cashCouponListO {
						cashCoupon := v2.(*model.CashCoupon)
						if cashCoupon.Id == cashCouponOrder.CashCouponId {
							totalCashCouponPrice += int(float32(cashCoupon.DiscountAmount) * float32(money) / float32(cashCouponOrder.Price))
						}
					}
					totalCashCouponPrice += money
				}
			}
		}
		if shop.TotalCashCouponPrice != totalCashCouponPrice || shop.TotalCashCouponNumber != totalCashCouponNumber {
			shop.TotalCashCouponPrice = totalCashCouponPrice
			shop.TotalCashCouponNumber = totalCashCouponNumber
			_, err = service.Update(serviceShop, shop, "totalCashCouponPrice", "totalCashCouponNumber")
			checkErr(err)
		}
		globalTotalCashCouponNumber += shop.TotalCashCouponNumber
		globalTotalCashCouponPrice += shop.TotalCashCouponPrice
	}
	cache.Set(cache.CacheNameGlobalTotalCashCouponNumber, globalTotalCashCouponNumber, 0)
	cache.Set(cache.CacheNameGlobalTotalCashCouponPrice, globalTotalCashCouponPrice, 0)
}