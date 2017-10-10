package util

import (
	"github.com/satori/go.uuid"
	"strconv"
	"strings"
	"sync"
	"time"
	"mimi/djq/cache"
)

type orderNumber struct {
	Type  string
	Time  time.Time
	Count int
	mutex sync.Mutex
}

//func (on *orderNumber) build() string {
//	list := make([]string, 15, 15)
//	list[0] = on.Type
//
//	list[1] = strconv.Itoa((on.Time.Year() % 100) / 10)
//	list[2] = strconv.Itoa(int(on.Time.Month()) / 10)
//	list[3] = strconv.Itoa(on.Time.Day() / 10)
//	list[4] = strconv.Itoa(on.Time.Hour() / 10)
//	list[5] = strconv.Itoa(on.Time.Minute() / 10)
//	list[6] = strconv.Itoa(on.Time.Second() / 10)
//	list[7] = strconv.Itoa(on.Count / 10)
//
//	list[8] = strconv.Itoa(on.Count % 10)
//	list[9] = strconv.Itoa(on.Time.Second() % 10)
//	list[10] = strconv.Itoa(on.Time.Minute() % 10)
//	list[11] = strconv.Itoa(on.Time.Hour() % 10)
//	list[12] = strconv.Itoa(on.Time.Day() % 10)
//	list[13] = strconv.Itoa(int(on.Time.Month()) % 10)
//	list[14] = strconv.Itoa(on.Time.Year() % 10)
//
//	return strings.Join(list, "")
//}

func (on *orderNumber) build() string {
	list := make([]string, 15, 15)
	list[0] = on.Type

	list[1] = strconv.Itoa(on.Time.Year() % 10)
	list[2] = strconv.Itoa((on.Time.YearDay() % 100) / 10)
	list[3] = strconv.Itoa(on.Time.Hour() / 10)
	list[4] = strconv.Itoa(on.Time.YearDay() % 10)
	list[8] = strconv.Itoa(on.Count % 10)
	list[6] = strconv.Itoa((on.Time.YearDay()) / 100)
	list[7] = strconv.Itoa(on.Time.Hour() % 10)
	list[5] = strconv.Itoa(on.Count / 10)

	return strings.Join(list, "")
}
func (on *orderNumber) next() string {
	on.mutex.Lock()
	defer func() {
		if err := recover(); err != nil {
			on.mutex.Unlock()
			panic(err)
		}
		on.mutex.Unlock()
	}()
	var cnLastTime, cnLastCount string
	if on.Type == "P" {
		cnLastTime = cache.CacheNamePresentOrderNumberLastTime
		cnLastCount = cache.CacheNamePresentOrderNumberLastCount
	} else {
		cnLastTime = cache.CacheNameCashCouponOrderNumberLastTime
		cnLastCount = cache.CacheNameCashCouponOrderNumberLastCount
	}
	timeStr, err := cache.Get(cnLastTime)
	if err != nil {
		panic(err)
	}
	countStr, err := cache.Get(cnLastCount)
	if err != nil {
		panic(err)
	}
	now := time.Now()
	if countStr == "" {
		on.Count = 0
	} else {
		on.Count, err = strconv.Atoi(countStr)
		if err != nil {
			panic(err)
		}
	}
	if timeStr == "" {
		on.Time = now
		on.Count = 0
	} else {
		on.Time, err = ParseTimeFromDB(timeStr)
		if err != nil {
			panic(err)
		}
	}
	if now.Sub(on.Time) > time.Hour * 1 {
		on.Time = now
		on.Count = 0
	} else if on.Count > 99 {
		on.Time = on.Time.Add(time.Hour * 1)
		on.Count = 0
	}
	str := on.build()
	on.Count = on.Count + 1

	timeStr = StringTime4DB(on.Time)
	countStr = strconv.Itoa(on.Count)
	err = cache.Set(cnLastTime, timeStr, 0)
	if err != nil {
		panic(err)
	}
	err = cache.Set(cnLastCount, countStr, 0)
	if err != nil {
		panic(err)
	}
	return str
}

//func (on *orderNumber) unLock() {
//	if err := recover(); err != nil {
//		on.mutex.Unlock()
//		panic(err)
//	}
//	on.mutex.Unlock()
//}

var lastPresentOrderNumber orderNumber
var lastCashCouponOrderNumber orderNumber

func BuildUUID() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}

func BuildPresentOrderNumber() string {
	if lastPresentOrderNumber.Type == "" {
		lastPresentOrderNumber.Type = "P"
	}
	return lastPresentOrderNumber.next()
}

func BuildCashCouponOrderNumber() string {
	if lastCashCouponOrderNumber.Type == "" {
		lastCashCouponOrderNumber.Type = "C"
	}
	return lastCashCouponOrderNumber.next()
}
