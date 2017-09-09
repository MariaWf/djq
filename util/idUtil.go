package util

import (
	"fmt"
	"github.com/satori/go.uuid"
	"strconv"
	"strings"
	"sync"
	"time"
)

type orderNumber struct {
	Type  string
	Time  time.Time
	Count int
	mutex sync.Mutex
}

func (on *orderNumber) build() string {
	list := make([]string, 15, 15)
	list[0] = on.Type

	list[1] = strconv.Itoa((on.Time.Year() % 100) / 10)
	list[2] = strconv.Itoa(int(on.Time.Month()) / 10)
	list[3] = strconv.Itoa(on.Time.Day() / 10)
	list[4] = strconv.Itoa(on.Time.Hour() / 10)
	list[5] = strconv.Itoa(on.Time.Minute() / 10)
	list[6] = strconv.Itoa(on.Time.Second() / 10)
	list[7] = strconv.Itoa(on.Count / 10)

	list[8] = strconv.Itoa(on.Count % 10)
	list[9] = strconv.Itoa(on.Time.Second() % 10)
	list[10] = strconv.Itoa(on.Time.Minute() % 10)
	list[11] = strconv.Itoa(on.Time.Hour() % 10)
	list[12] = strconv.Itoa(on.Time.Day() % 10)
	list[13] = strconv.Itoa(int(on.Time.Month()) % 10)
	list[14] = strconv.Itoa(on.Time.Year() % 10)

	return strings.Join(list, "")
}
func (on *orderNumber) next() string {
	on.mutex.Lock()
	defer func() {
		if err := recover(); err != nil {
			on.mutex.Unlock()
			fmt.Println(err)
			panic(err)
		}
		on.mutex.Unlock()
	}()
	now := time.Now()
	if now.Sub(on.Time) > time.Second*1 {
		on.Time = now
		on.Count = 0
	} else if on.Count == 99 {
		on.Time = on.Time.Add(time.Second * 1)
		on.Count = 0
	}
	str := on.build()
	on.Count = on.Count + 1
	//str:="ss"
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
