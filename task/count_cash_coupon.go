package task

import (
	"fmt"
	"time"
)

func CountCashCoupon() {
	FixTimeCycle(func() {
		fmt.Println(time.Now())
	}, 2, 0, 0)
}