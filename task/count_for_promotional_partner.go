package task

import (
	"time"
	"fmt"
)

//每天凌晨2点统计合作伙伴数据
func CountForPromotionalPartner() {
	defaultHour := 2
	now := time.Now()
	if now.Hour() != defaultHour {
		goto nextCycle
	}
	fmt.Println("countPromotionalPartner")
	nextCycle:
	var sleepTime time.Duration
	if now.Hour() > defaultHour {
		sleepTime = time.Duration( 23 - now.Hour()) * time.Hour
	}
	time.Sleep(sleepTime)
	go CountForPromotionalPartner()
}
