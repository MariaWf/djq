package task

import (
	"log"
	"sync"
	"time"
)

var lock = sync.Mutex{}

// hour [0,23] -1 代表不使用
//minute[0,59]
//second[0,59]
func FixTimeCycle(f interface{}, hour int, minute int, second int) {
	var timeSleep time.Duration
	var d time.Duration
	now := time.Now()
	h := now.Hour()
	m := now.Minute()
	s := now.Second()

	if hour > -1 {
		timeSleep += time.Duration(hour-h) * time.Hour
		timeSleep += time.Duration(minute-m) * time.Minute
		timeSleep += time.Duration(second-s) * time.Second
		if timeSleep < 0 {
			timeSleep += 24 * time.Hour
		}
		d = 24 * time.Hour
	} else if minute > -1 {
		timeSleep += time.Duration(minute-m) * time.Minute
		timeSleep += time.Duration(second-s) * time.Second
		if timeSleep < 0 {
			timeSleep += 60 * time.Minute
		}
		d = 60 * time.Minute
	} else if second > -1 {
		timeSleep += time.Duration(second-s) * time.Second
		if timeSleep < 0 {
			timeSleep += 60 * time.Second
		}
		d = 60 * time.Second
	} else {
		timeSleep = 0
		d = time.Second
	}
	time.Sleep(timeSleep)
	FixTimeIntervalCycle(f, d)
}

func FixTimeIntervalCycle(f interface{}, d time.Duration) {
	//for {
	//	f.(func())()
	//	time.Sleep(d)
	//}
	go f.(func())()
	ticker := time.NewTicker(d)
	for range ticker.C {
		go f.(func())()
	}
}

func goFunc(f interface{}, args ...interface{}) {
	if len(args) > 1 {
		go f.(func(...interface{}))(args)
	} else if len(args) == 1 {
		go f.(func(interface{}))(args[0])
	} else {
		go f.(func())()
	}
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
