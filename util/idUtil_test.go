package util

import (
	"fmt"
	"testing"
)

func TestBuildPresentOrderNumber(t *testing.T) {
	list := make([]string, 0, 100)
	c := make(chan []string, 5)
	for i := 0; i < 5; i++ {
		go testBuildPresentOrderNumber(c)
		list = append(<-c, list...)
	}
	same := 0
	for i, _ := range list {
		fmt.Println(list[i])
		for j := i + 1; j < len(list); j++ {
			if list[i] == list[j] {
				same++
			}
		}
	}
	if same > 0 {
		t.Error(len(list), same)
		//}else{
		//	for _,v:=range list{
		//		t.Log(v)
		//	}
		//	t.Log(list)
	}
}
func TestBuildCashCouponOrderNumber(t *testing.T) {
	list := make([]string, 0, 100)
	c := make(chan []string, 5)
	for i := 0; i < 5; i++ {
		go testBuildCashCouponOrderNumber(c)
		list = append(<-c, list...)
	}
	same := 0
	for i, _ := range list {
		for j := i + 1; j < len(list); j++ {
			if list[i] == list[j] {
				same++
			}
		}
	}
	if same > 0 {
		t.Error(len(list), same)
		//}else{
		//	for _,v:=range list{
		//		t.Log(v)
		//	}
		//	t.Log(list)
	}
}

func testBuildPresentOrderNumber(c chan []string) {
	total := 1000
	list := make([]string, total, total)
	for i := 0; i < total; i++ {
		list[i] = BuildPresentOrderNumber()
	}
	c <- list
}

func testBuildCashCouponOrderNumber(c chan []string) {
	total := 1000
	list := make([]string, total, total)
	for i := 0; i < total; i++ {
		list[i] = BuildCashCouponOrderNumber()
	}
	c <- list
}

func TestBuildUUID(t *testing.T) {
	fmt.Println(BuildUUID())
}

func TestBuildOrderNumber(t *testing.T) {
	fmt.Println(BuildCashCouponOrderNumber())
}
