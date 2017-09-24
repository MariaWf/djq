package task

import (
	"mimi/djq/service"
	"mimi/djq/cache"
	"log"
)

//定时检查过期代金券，
func CheckExpiredCashCoupon() {
	err := cache.Set(cache.CacheNameCheckingExpiredCashCoupon, "false", 0)
	if err != nil {
		log.Println(err)
	}
	FixTimeCycle(CheckExpiredCashCouponAction, 4, 0, 0)
}

func CheckExpiredCashCouponAction() {
	lock.Lock()
	counting, err := cache.Get(cache.CacheNameCheckingExpiredCashCoupon)
	if err != nil {
		lock.Unlock()
		checkErr(err)
	}
	if counting == "true" {
		lock.Unlock()
		return
	}
	counting = "true"
	err = cache.Set(cache.CacheNameCheckingExpiredCashCoupon, counting, 0)
	if err != nil {
		lock.Unlock()
		checkErr(err)
	}
	lock.Unlock()
	defer cache.Set(cache.CacheNameCheckingExpiredCashCoupon, "false", 0)

	serviceCashCoupon := &service.CashCoupon{}
	err = serviceCashCoupon.RefreshExpired()
	checkErr(err)
}