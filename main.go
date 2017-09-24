package main

import (
	"log"
	"mimi/djq/config"
	"mimi/djq/constant"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/router"
	"mimi/djq/service"
	"mimi/djq/util"
	"os"
	"path/filepath"
	"time"
	"strconv"
	"math/rand"
	"mimi/djq/task"
	"mimi/djq/init"
)

func main() {
	log.Println("------start：" + time.Now().String())
	init.InitGlobalLog()
	log.Println("------start：" + time.Now().String())
	log.SetFlags(log.LstdFlags | log.Llongfile)
	initData()
	if config.Get("buildTestData") == "true" {
		initTestData()
	}
	if "true" == config.Get("task_run") {
		beginTask()
	}
	router.Begin()
}

func beginTask() {
	go task.CheckPayingOrder()
	go task.CheckRefundingOrder()
	go task.AgreeNotUsedRefunding()
	go task.CountCashCoupon()
	go task.CountForPromotionalPartner()
	go task.CheckExpiredCashCoupon()
	go task.CheckExpiredPresent()
}

func initData() {
	constant.AdminId = init.InitAdmin()
}

func initTestData() {
	//tempInitTestShop()
	init.InitTestAdmin()
	init.InitTestAdvertisement()
	init.InitTestShop()
	init.InitTestUser()
	init.InitTestCashCouponOrder()
}