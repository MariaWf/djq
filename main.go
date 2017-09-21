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
)

func main() {
	log.Println("------start：" + time.Now().String())
	initLog()
	log.Println("------start：" + time.Now().String())
	log.SetFlags(log.LstdFlags | log.Llongfile)
	initData()
	if config.Get("buildTestData") == "true" {
		initTestData()
	}
	//beginTask()
	router.Begin()
}

func beginTask(){
	go task.CheckPayingOrder()
	go task.CheckRefundingOrder()
}

func initLog() {
	if "false" == config.Get("output_log") {
		return
	}
	globalLogUrl := config.Get("global_log")
	if globalLogUrl == "" {
		globalLogUrl = "logs/global.log"
	}
	path := filepath.Dir(globalLogUrl)
	os.MkdirAll(path, 0777)
	logFile, err := os.OpenFile(globalLogUrl, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.Println()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func initData() {
	constant.AdminId = initAdmin()
}

func initTestData() {
	//tempInitTestShop()
	initTestAdmin()
	initTestAdvertisement()
	initTestShop()
	initTestUser()
	initTestCashCouponOrder()
}

func initRole() string {
	serviceRole := &service.Role{}
	argRole := &arg.Role{}
	argRole.NameEqual = config.Get("adminRole")
	roleList, err := service.Find(serviceRole, argRole)
	checkErr(err)
	if roleList == nil || len(roleList) == 0 {
		role := &model.Role{}
		role.PermissionList = model.GetPermissionList()
		role.Name = config.Get("adminRole")
		role.Description = "超级管理员不能删除，不能修改"
		role, err := serviceRole.Add(role)
		checkErr(err)
		return role.Id
	} else {
		role := roleList[0].(*model.Role)
		role.BindStr2PermissionList()
		if len(role.PermissionList) != len(model.GetPermissionList()) {
			role.PermissionList = model.GetPermissionList()
			role, err := serviceRole.Update(role)
			checkErr(err)
			return role.Id
		}
	}
	return roleList[0].(*model.Role).Id
}

func initAdmin() string {
	roleId := initRole()
	constant.AdminRoleId = roleId
	serviceAdmin := &service.Admin{}
	argAdmin := &arg.Admin{}
	argAdmin.NameEqual = config.Get("adminName")
	objList, err := service.Find(serviceAdmin, argAdmin)
	checkErr(err)
	if objList == nil || len(objList) == 0 {
		obj := &model.Admin{}
		obj.Name = config.Get("adminName")
		obj.Password = config.Get("adminPassword")
		obj.Password, err = util.EncryptPassword(obj.Password)
		checkErr(err)
		obj.RoleList = []*model.Role{{Id: roleId}}
		obj, err := serviceAdmin.Add(obj)
		checkErr(err)
		return obj.Id
	}
	return objList[0].(*model.Admin).Id
}

func initTestRole() {
	serviceRole := &service.Role{}
	argRole := &arg.Role{}
	count, err := service.Count(serviceRole, argRole)
	checkErr(err)
	if count < 5 {
		for i := 0; i < 50; i++ {
			obj := &model.Role{}
			pl := make([]*model.Permission, 0, 10)
			for _, v := range model.GetPermissionList() {
				if rand.Intn(2) < 1 {
					pl = append(pl, v)
				}
			}
			obj.PermissionList = pl
			obj.Name = "名称" + strconv.Itoa(i)
			obj.Description = "描述" + strconv.Itoa(i)
			obj, err := serviceRole.Add(obj)
			checkErr(err)
		}
	}
}

func initTestAdmin() {
	initTestRole()
	serviceAdmin := &service.Admin{}
	argAdmin := &arg.Admin{}
	count, err := service.Count(serviceAdmin, argAdmin)
	checkErr(err)
	if count < 5 {
		serviceRole := &service.Role{}
		argRole := &arg.Role{}
		roleList, err := service.Find(serviceRole, argRole)
		checkErr(err)
		for i := 0; i < 50; i++ {
			obj := &model.Admin{}
			rl := make([]*model.Role, 0, 10)
			for _, v := range roleList {
				if rand.Intn(10) < 2 {
					rl = append(rl, v.(*model.Role))
				}
			}
			obj.Mobile = "12345678910"
			obj.RoleList = rl
			obj.Name = "name" + strconv.Itoa(i)
			obj.Password = "123123"
			obj.Locked = rand.Intn(2) < 1
			obj.Password, err = util.EncryptPassword(obj.Password)
			checkErr(err)
			obj, err := serviceAdmin.Add(obj)
			checkErr(err)
		}
	}
}

func initTestAdvertisement() {
	serviceAdvertisement := &service.Advertisement{}
	argAdvertisement := &arg.Advertisement{}
	count, err := service.Count(serviceAdvertisement, argAdvertisement)
	checkErr(err)
	if count < 5 {
		for i := 0; i < 50; i++ {
			obj := &model.Advertisement{}
			obj.Name = "名称" + strconv.Itoa(i)
			obj.Image = "https://www.baidu.com/img/bd_logo1.png"
			obj.Link = "https://www.baidu.com"
			obj.Priority = rand.Intn(1000)
			obj.Hide = rand.Intn(2) < 1
			obj.Description = "描述" + strconv.Itoa(i)
			_, err := service.Add(serviceAdvertisement, obj)
			checkErr(err)
		}
	}
}

func initTestShopClassification() {
	serviceShopClassification := &service.ShopClassification{}
	argShopClassification := &arg.ShopClassification{}
	count, err := service.Count(serviceShopClassification, argShopClassification)
	checkErr(err)
	if count < 5 {
		for i := 0; i < 50; i++ {
			obj := &model.ShopClassification{}
			obj.Name = "名称" + strconv.Itoa(i)
			obj.Priority = rand.Intn(1000)
			obj.Hide = rand.Intn(2) < 1
			obj.Description = "描述" + strconv.Itoa(i)
			_, err := service.Add(serviceShopClassification, obj)
			checkErr(err)
		}
	}
}

func initTestShop() {
	initTestShopClassification()
	serviceShop := &service.Shop{}
	argShop := &arg.Shop{}
	count, err := service.Count(serviceShop, argShop)
	checkErr(err)
	if count < 5 {
		serviceShopClassification := &service.ShopClassification{}
		argShopClassification := &arg.ShopClassification{}
		roleList, err := service.Find(serviceShopClassification, argShopClassification)
		checkErr(err)
		for i := 0; i < 50; i++ {
			obj := &model.Shop{}
			rl := make([]*model.ShopClassification, 0, 10)
			for _, v := range roleList {
				if rand.Intn(10) < 2 {
					rl = append(rl, v.(*model.ShopClassification))
				}
			}
			obj.ShopClassificationList = rl
			obj.Name = "名称" + strconv.Itoa(i)
			obj.Logo = "https://www.baidu.com/img/bd_logo1.png"
			obj.PreImage = "https://www.baidu.com/img/bd_logo1.png"
			obj.TotalCashCouponNumber = rand.Intn(1000)
			obj.TotalCashCouponPrice = rand.Intn(1000) * obj.TotalCashCouponNumber
			obj.Introduction = "介绍" + strconv.Itoa(i)
			obj.Address = "地址" + strconv.Itoa(i)
			obj.Priority = rand.Intn(1000)
			obj.Hide = rand.Intn(2) < 1
			obj, err := serviceShop.Add(obj)
			checkErr(err)
			initTestShopAccount(obj.Id, i)
			initTestShopIntroductionImage(obj.Id)
			initTestCashCoupon(obj.Id)
		}
	}
}
func tempInitTestShop() {
	serviceShop := &service.Shop{}
	argShop := &arg.Shop{}
	list, err := service.Find(serviceShop, argShop)
	checkErr(err)
	for _, obj := range list {
		//initTestShopAccount(obj.(*model.Shop).Id, i)
		//initTestShopIntroductionImage(obj.(*model.Shop).Id)
		initTestCashCoupon(obj.(*model.Shop).Id)
	}
}

func initTestShopAccount(shopId string, index int) {
	serviceShopAccount := &service.ShopAccount{}
	total := rand.Intn(5)
	var err error
	for i := 0; i < total; i++ {
		obj := &model.ShopAccount{}
		obj.ShopId = shopId
		obj.Name = "名称" + strconv.Itoa(index) + "_" + strconv.Itoa(i)
		obj.Password = "123123"
		obj.Password, err = util.EncryptPassword(obj.Password)
		checkErr(err)
		obj.MoneyChance = rand.Intn(20)
		obj.TotalMoney = rand.Intn(1000)
		obj.Locked = rand.Intn(2) < 1
		obj.Description = "描述" + strconv.Itoa(i)
		_, err := serviceShopAccount.Add(obj)
		checkErr(err)
	}
}

func initTestShopIntroductionImage(shopId string) {
	serviceShopIntroductionImage := &service.ShopIntroductionImage{}
	total := rand.Intn(10)
	for i := 0; i < total; i++ {
		obj := &model.ShopIntroductionImage{}
		obj.ShopId = shopId
		obj.Priority = rand.Intn(1000)
		obj.Hide = rand.Intn(2) < 1
		obj.ContentUrl = "https://www.baidu.com/img/bd_logo1.png"
		_, err := service.Add(serviceShopIntroductionImage, obj)
		checkErr(err)
	}
}

func initTestCashCoupon(shopId string) {
	serviceCashCoupon := &service.CashCoupon{}
	total := rand.Intn(10)
	for i := 0; i < total; i++ {
		obj := &model.CashCoupon{}
		obj.ShopId = shopId
		obj.Name = "名称" + strconv.Itoa(i)
		obj.PreImage = "https://www.baidu.com/img/bd_logo1.png"
		obj.DiscountAmount = 10 + rand.Intn(2000)
		obj.Price = obj.DiscountAmount + rand.Intn(2000)
		t := time.Now()
		if rand.Intn(2) < 1 {
			t = t.Add(time.Hour * time.Duration(rand.Int63n(1000)))
		} else {
			obj.Expired = true
			t = t.Add(-time.Hour * time.Duration(rand.Int63n(1000)))
		}
		obj.ExpiryDate = util.StringTime4DB(t)
		obj.Hide = rand.Intn(2) < 1
		obj.Priority = rand.Intn(1000)
		_, err := service.Add(serviceCashCoupon, obj)
		checkErr(err)
	}
}

func initTestPromotionalPartner() {
	servicePromotionalPartner := &service.PromotionalPartner{}
	argPromotionalPartner := &arg.PromotionalPartner{}
	count, err := service.Count(servicePromotionalPartner, argPromotionalPartner)
	checkErr(err)
	if count < 5 {
		checkErr(err)
		for i := 0; i < 50; i++ {
			obj := &model.PromotionalPartner{}
			obj.Name = "名称" + strconv.Itoa(i)
			obj.Description = "描述" + strconv.Itoa(i)
			_, err := service.Add(servicePromotionalPartner, obj)
			checkErr(err)
		}
	}
}

func initTestUser() {
	initTestPromotionalPartner()
	servicePromotionalPartner := &service.PromotionalPartner{}
	argPromotionalPartner := &arg.PromotionalPartner{}
	promotionalPartnerList, err := service.Find(servicePromotionalPartner, argPromotionalPartner)
	checkErr(err)

	initTestPresent()
	servicePresent := &service.Present{}
	argPresent := &arg.Present{}
	presentList, err := service.Find(servicePresent, argPresent)
	checkErr(err)

	serviceUser := &service.User{}
	argUser := &arg.User{}
	count, err := service.Count(serviceUser, argUser)
	checkErr(err)
	if count < 5 {
		checkErr(err)
		for i := 10; i < 80; i++ {
			obj := &model.User{}
			obj.Mobile = "123456789" + strconv.Itoa(i)
			obj.PresentChance = rand.Intn(10)
			obj.Locked = rand.Intn(2) < 1
			obj.Shared = rand.Intn(2) < 1
			if rand.Intn(2) < 1 {
				obj.PromotionalPartnerId = promotionalPartnerList[rand.Intn(len(promotionalPartnerList))].(*model.PromotionalPartner).Id
			}
			_, err := service.Add(serviceUser, obj)
			for j := 0; j < 10; j++ {
				if rand.Intn(2) < 1 {
					initTestPresentOrder(obj.Id, presentList[rand.Intn(len(presentList))].(*model.Present).Id)
				}
			}
			checkErr(err)

		}
	}
}

func initTestPresent() {
	servicePresent := &service.Present{}
	argPresent := &arg.Present{}
	count, err := service.Count(servicePresent, argPresent)
	checkErr(err)
	if count < 5 {
		checkErr(err)
		for i := 0; i < 50; i++ {
			obj := &model.Present{}
			obj.Name = "名称" + strconv.Itoa(i)
			obj.Weight = rand.Intn(1000)
			obj.Address = "地址" + strconv.Itoa(i)
			t := time.Now()
			if rand.Intn(2) < 1 {
				t = t.Add(time.Hour * time.Duration(rand.Int63n(1000)))
			} else {
				t = t.Add(-time.Hour * time.Duration(rand.Int63n(1000)))
			}
			obj.ExpiryDate = util.StringTime4DB(t)
			obj.Hide = rand.Intn(2) < 1
			obj.Image = "https://www.baidu.com/img/bd_logo1.png"
			obj.Requirement = 1000 + rand.Intn(2000)
			obj.Stock = 1000 + rand.Intn(2000) + obj.Requirement
			_, err := service.Add(servicePresent, obj)
			checkErr(err)
		}
	}
}

func initTestPresentOrder(userId, presentId string) {
	servicePresentOrder := &service.PresentOrder{}
	total := rand.Intn(10)
	for i := 0; i < total; i++ {
		obj := &model.PresentOrder{}
		obj.PresentId = presentId
		obj.UserId = userId
		obj.Number = util.BuildPresentOrderNumber()
		if rand.Intn(2) < 1 {
			obj.Status = constant.PresentOrderStatusReceived
		} else {
			obj.Status = constant.PresentOrderStatusWaiting2Receive
		}
		_, err := service.Add(servicePresentOrder, obj)
		checkErr(err)
	}
}

func initTestRefundReason() {
	serviceRefundReason := &service.RefundReason{}
	argRefundReason := &arg.RefundReason{}
	count, err := service.Count(serviceRefundReason, argRefundReason)
	checkErr(err)
	if count < 5 {
		checkErr(err)
		for i := 0; i < 50; i++ {
			obj := &model.RefundReason{}
			obj.Description = "退款理由" + strconv.Itoa(i)
			obj.Priority = rand.Intn(1000)
			obj.Hide = rand.Intn(2) < 1
			_, err := service.Add(serviceRefundReason, obj)
			checkErr(err)
		}
	}
}

func initTestCashCouponOrder() {
	initTestRefundReason()
	serviceCashCoupon := &service.CashCoupon{}
	argCashCoupon := &arg.CashCoupon{}
	cashCouponList, err := service.Find(serviceCashCoupon, argCashCoupon)
	checkErr(err)

	serviceUser := &service.User{}
	argUser := &arg.User{}
	userList, err := service.Find(serviceUser, argUser)
	checkErr(err)

	serviceCashCouponOrder := &service.CashCouponOrder{}
	argCashCouponOrder := &arg.CashCouponOrder{}
	count, err := service.Count(serviceCashCouponOrder, argCashCouponOrder)
	checkErr(err)
	if count < 5 {
		checkErr(err)
		status := []int{constant.CashCouponOrderStatusInCart,
			constant.CashCouponOrderStatusPaidNotUsed,
			constant.CashCouponOrderStatusUsed,
			constant.CashCouponOrderStatusNotUsedRefunding,
			constant.CashCouponOrderStatusUsedRefunding,
			constant.CashCouponOrderStatusNotUsedRefunded,
			constant.CashCouponOrderStatusUsedRefunded}
		for i := 0; i < 50; i++ {
			cashCoupon := cashCouponList[rand.Intn(len(cashCouponList))].(*model.CashCoupon)
			obj := &model.CashCouponOrder{}
			obj.UserId = userList[rand.Intn(len(userList))].(*model.User).GetId()
			obj.CashCouponId = cashCoupon.GetId()
			obj.Number = util.BuildCashCouponOrderNumber()
			obj.Status = status[rand.Intn(len(status))]
			obj.Price = cashCoupon.Price
			obj.PayOrderNumber = util.BuildUUID()
			if obj.Status == constant.CashCouponOrderStatusUsedRefunded || obj.Status == constant.CashCouponOrderStatusUsedRefunding {
				obj.RefundAmount = rand.Intn(obj.Price)
			}
			_, err := service.Add(serviceCashCouponOrder, obj)
			initTestRefund(obj)
			checkErr(err)
		}
	}
}

func initTestRefund(cashCouponOrder *model.CashCouponOrder) {
	serviceRefundReason := &service.RefundReason{}
	argRefundReason := &arg.RefundReason{}
	refundReasonList, err := service.Find(serviceRefundReason, argRefundReason)
	checkErr(err)

	serviceRefund := &service.Refund{}
	//status := []int{constant.RefundStatusNotUsedRefunding,
	//	constant.RefundStatusNotUsedRefundSuccess,
	//	constant.RefundStatusNotUsedRefundCancel,
	//	constant.RefundStatusUsedRefunding,
	//	constant.RefundStatusUsedRefundSuccess,
	//	constant.RefundStatusUsedRefundFail,
	//	constant.RefundStatusUsedRefundCancel}

	switch cashCouponOrder.Status{
	case constant.CashCouponOrderStatusInCart:
		return
	case constant.CashCouponOrderStatusPaidNotUsed:
		goto PaidNotUsed
	case constant.CashCouponOrderStatusUsed:
		goto Used
	case constant.CashCouponOrderStatusNotUsedRefunding:
		obj := buildTestRefundModel(refundReasonList, cashCouponOrder)
		obj.Status = constant.RefundStatusNotUsedRefunding
		_, err := service.Add(serviceRefund, obj)
		checkErr(err)
		goto PaidNotUsed
	case constant.CashCouponOrderStatusUsedRefunding:
		if cashCouponOrder.RefundAmount > 0 {
			obj := buildTestRefundModel(refundReasonList, cashCouponOrder)
			obj.Status = constant.RefundStatusUsedRefundSuccess
			obj.RefundAmount = cashCouponOrder.RefundAmount
			_, err := service.Add(serviceRefund, obj)
			checkErr(err)
		}
		obj := buildTestRefundModel(refundReasonList, cashCouponOrder)
		obj.Status = constant.RefundStatusUsedRefunding
		obj.RefundAmount = rand.Intn(cashCouponOrder.Price - cashCouponOrder.RefundAmount) + 1
		_, err := service.Add(serviceRefund, obj)
		checkErr(err)
		goto Used
	case constant.CashCouponOrderStatusNotUsedRefunded:
		obj := buildTestRefundModel(refundReasonList, cashCouponOrder)
		obj.Status = constant.RefundStatusNotUsedRefundSuccess
		_, err := service.Add(serviceRefund, obj)
		checkErr(err)
		goto PaidNotUsed
	case constant.CashCouponOrderStatusUsedRefunded:
		obj := buildTestRefundModel(refundReasonList, cashCouponOrder)
		obj.Status = constant.RefundStatusUsedRefundSuccess
		obj.RefundAmount = cashCouponOrder.RefundAmount
		_, err := service.Add(serviceRefund, obj)
		checkErr(err)
		goto Used
	}

	Used:
	if rand.Intn(2) < 1 {
		obj := buildTestRefundModel(refundReasonList, cashCouponOrder)
		obj.Status = constant.RefundStatusUsedRefundCancel
		_, err := service.Add(serviceRefund, obj)
		checkErr(err)
	}
	if rand.Intn(2) < 1 {
		obj := buildTestRefundModel(refundReasonList, cashCouponOrder)
		obj.Status = constant.RefundStatusUsedRefundFail
		_, err := service.Add(serviceRefund, obj)
		checkErr(err)
	}

	PaidNotUsed:
	if rand.Intn(2) < 1 {
		obj := buildTestRefundModel(refundReasonList, cashCouponOrder)
		obj.Status = constant.RefundStatusNotUsedRefundCancel
		_, err := service.Add(serviceRefund, obj)
		checkErr(err)
	}

	//for i := 0; i < 3; i++ {
	//	obj := &model.Refund{}
	//	obj.Common = "平台意见" + strconv.Itoa(rand.Intn(1000))
	//	obj.CashCouponOrderId = cashCouponOrder.GetId()
	//	obj.Evidence = "https://www.baidu.com/img/bd_logo1.png"
	//	obj.Reason = refundReasonList[rand.Intn(len(refundReasonList))].(*model.RefundReason).Description
	//	if rand.Intn(10) < 1 {
	//		obj.Reason = "其他退款理由" + strconv.Itoa(rand.Intn(1000))
	//	}
	//	obj.RefundAmount = cashCouponOrder.Price
	//	obj.Status = status[rand.Intn(len(status))]
	//	_, err := service.Add(serviceRefund, obj)
	//	checkErr(err)
	//}
}

func buildTestRefundModel(refundReasonList []interface{}, cashCouponOrder *model.CashCouponOrder) *model.Refund {
	obj := &model.Refund{}
	obj.Common = "平台意见" + strconv.Itoa(rand.Intn(1000))
	obj.CashCouponOrderId = cashCouponOrder.GetId()
	obj.Evidence = "https://www.baidu.com/img/bd_logo1.png"
	obj.Reason = refundReasonList[rand.Intn(len(refundReasonList))].(*model.RefundReason).Description
	if rand.Intn(10) < 1 {
		obj.Reason = "其他退款理由" + strconv.Itoa(rand.Intn(1000))
	}
	obj.RefundAmount = cashCouponOrder.Price
	//obj.Status = status[rand.Intn(len(status))]
	return obj
}