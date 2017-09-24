package task

import (
	"mimi/djq/service"
	"mimi/djq/cache"
	"log"
)

//定时检查过期礼品，
func CheckExpiredPresent() {
	err := cache.Set(cache.CacheNameCheckingExpiredPresent, "false", 0)
	if err != nil {
		log.Println(err)
	}
	FixTimeCycle(CheckExpiredPresentAction, 4, 0, 0)
}

func CheckExpiredPresentAction() {
	lock.Lock()
	counting, err := cache.Get(cache.CacheNameCheckingExpiredPresent)
	if err != nil {
		lock.Unlock()
		checkErr(err)
	}
	if counting == "true" {
		lock.Unlock()
		return
	}
	counting = "true"
	err = cache.Set(cache.CacheNameCheckingExpiredPresent, counting, 0)
	if err != nil {
		lock.Unlock()
		checkErr(err)
	}
	lock.Unlock()
	defer cache.Set(cache.CacheNameCheckingExpiredPresent, "false", 0)

	servicePresent := &service.Present{}
	err = servicePresent.RefreshExpired()
	checkErr(err)
}