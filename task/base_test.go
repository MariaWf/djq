package task

import (
	"testing"
	"time"
)

func TestCheckPayingOrder(t *testing.T) {
	go CheckPayingOrder()
	time.Sleep(time.Hour)
}
