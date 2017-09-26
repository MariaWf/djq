package initialization

import (
	"math/rand"
	"mimi/djq/constant"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/service"
	"mimi/djq/util"
	"strconv"
	"strings"
	"time"
)

const specialMobile string = "15017528974"

const testImageHead string = "http://static.51zxiu.cn/app/djq/test/"

func getRandomPresentImage() string {
	return util.FormatImagePresent(testImageHead + "present" + strconv.Itoa(rand.Intn(10)+1) + ".jpg")
}

func getRandomShopLogo() string {
	return util.FormatImageShopLogo(testImageHead + "logo" + strconv.Itoa(rand.Intn(6)+1) + ".png")
}

func getRandomCashCoupon() string {
	return util.FormatImageCashCoupon(testImageHead + "djq" + strconv.Itoa(rand.Intn(7)+1) + ".png")
}

func getRandomShopPreImage() string {
	return util.FormatImageShopPreImage(testImageHead + "shopPreImage" + strconv.Itoa(rand.Intn(6)+1) + ".jpg")
}

func getRandomShopIntroductionImage() string {
	if rand.Intn(10) < 1 {
		return util.FormatImageShopIntroductionImage(testImageHead + "intro0.gif")
	}
	return util.FormatImageShopIntroductionImage(testImageHead + "intro" + strconv.Itoa(rand.Intn(10)+1) + ".jpg")
}

func getRandomAdvertisement() string {
	var url string
	if rand.Intn(2) < 1 {
		url = getRandomCashCoupon()
	} else {
		url = getRandomShopPreImage()
	}
	return util.FormatImageAdvertisement(url)
}

func getRandomEvidence() string {
	return testImageHead + "evidence.jpg"
}

func InitTestData() {
	InitTestAdmin()
	InitTestAdvertisement()
	shopList := InitTestShop()
	userList := InitTestUser()
	InitTestPresent()
	InitTestRefundReason()
	InitTestCashCouponOrderInCart(userList, shopList)
	InitTestPresentOrder(userList)

	//商家分类
	//商家
	//代金券
	//礼品
	//推广伙伴
	//用户
	//退款理由
	//
	//代金券订单1
	//礼品订单
	//退款申请
}

func InitTestRole() {
	serviceRole := &service.Role{}
	argRole := &arg.Role{}
	count, err := service.Count(serviceRole, argRole)
	checkErr(err)
	if count < 10 {
		for i := 0; i < 50; i++ {
			obj := &model.Role{}
			pl := make([]*model.Permission, 0, 10)
			for _, v := range model.GetPermissionList() {
				if rand.Intn(2) < 1 {
					pl = append(pl, v)
				}
			}
			obj.PermissionList = pl
			obj.Name = "角色名称" + strconv.Itoa(i)
			obj.Description = "角色描述描述描述描述描述描述描述描述描述描述描述描述描述描述描述" + strconv.Itoa(i)
			obj, err := serviceRole.Add(obj)
			checkErr(err)
		}
	}
}

func InitTestAdmin() {
	InitTestRole()
	serviceAdmin := &service.Admin{}
	argAdmin := &arg.Admin{}
	count, err := service.Count(serviceAdmin, argAdmin)
	checkErr(err)
	if count < 10 {
		serviceRole := &service.Role{}
		argRole := &arg.Role{}
		roleList, err := service.Find(serviceRole, argRole)
		checkErr(err)
		for i := 10; i < 50; i++ {
			obj := &model.Admin{}
			rl := make([]*model.Role, 0, 10)
			for _, v := range roleList {
				if rand.Intn(10) < 2 {
					rl = append(rl, v.(*model.Role))
				}
			}
			obj.Mobile = "123456789" + strconv.Itoa(i)
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

func InitTestAdvertisement() {
	serviceAdvertisement := &service.Advertisement{}
	argAdvertisement := &arg.Advertisement{}
	count, err := service.Count(serviceAdvertisement, argAdvertisement)
	checkErr(err)
	if count < 5 {
		for i := 0; i < 10; i++ {
			obj := &model.Advertisement{}
			obj.Name = "广告名称" + strconv.Itoa(i)
			obj.Image = getRandomAdvertisement()
			obj.Link = "https://www.baidu.com"
			obj.Priority = rand.Intn(1000)
			obj.Hide = rand.Intn(2) < 1
			obj.Description = "广告描述" + strconv.Itoa(i)
			_, err := service.Add(serviceAdvertisement, obj)
			checkErr(err)
		}
	}
}

func InitTestShopClassification() {
	serviceShopClassification := &service.ShopClassification{}
	argShopClassification := &arg.ShopClassification{}
	count, err := service.Count(serviceShopClassification, argShopClassification)
	checkErr(err)
	if count < 5 {
		for i := 0; i < 30; i++ {
			obj := &model.ShopClassification{}
			obj.Name = "分类名称" + strconv.Itoa(i)
			obj.Priority = rand.Intn(1000)
			obj.Hide = rand.Intn(2) < 1
			obj.Description = "分类描述" + strconv.Itoa(i)
			_, err := service.Add(serviceShopClassification, obj)
			checkErr(err)
		}
	}
}

func InitTestShop() []*model.Shop {
	InitTestShopClassification()
	serviceShop := &service.Shop{}
	argShop := &arg.Shop{}
	count, err := service.Count(serviceShop, argShop)
	checkErr(err)
	list := make([]*model.Shop, 0, 80)
	if count < 5 {
		serviceShopClassification := &service.ShopClassification{}
		argShopClassification := &arg.ShopClassification{}
		roleList, err := service.Find(serviceShopClassification, argShopClassification)
		checkErr(err)
		for i := 10; i < 99; i++ {
			obj := &model.Shop{}
			rl := make([]*model.ShopClassification, 0, 10)
			for _, v := range roleList {
				if rand.Intn(10) < 2 {
					rl = append(rl, v.(*model.ShopClassification))
				}
			}
			obj.ShopClassificationList = rl
			obj.Name = "名称" + strconv.Itoa(i)
			obj.Logo = getRandomShopLogo()
			obj.PreImage = getRandomShopPreImage()
			//obj.TotalCashCouponNumber = rand.Intn(1000)
			//obj.TotalCashCouponPrice = rand.Intn(1000) * obj.TotalCashCouponNumber
			obj.Introduction = "介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍介绍" + strconv.Itoa(i)
			obj.Address = "地址地址地址地址地址地址地址地址地址" + strconv.Itoa(i)
			obj.Priority = rand.Intn(1000)
			obj.Hide = rand.Intn(2) < 1
			obj, err := serviceShop.Add(obj)
			checkErr(err)
			obj.ShopAccountList = InitTestShopAccount(obj.Id, i)
			obj.ShopIntroductionImageList = InitTestShopIntroductionImage(obj.Id)
			obj.CashCouponList = InitTestCashCoupon(obj.Id)
			list = append(list, obj)
		}
	}
	return list
}

//func tempInitTestShop() {
//	serviceShop := &service.Shop{}
//	argShop := &arg.Shop{}
//	list, err := service.Find(serviceShop, argShop)
//	checkErr(err)
//	for _, obj := range list {
//		//InitTestShopAccount(obj.(*model.Shop).Id, i)
//		//InitTestShopIntroductionImage(obj.(*model.Shop).Id)
//		InitTestCashCoupon(obj.(*model.Shop).Id)
//	}
//}

func InitTestShopAccount(shopId string, index int) []*model.ShopAccount {
	serviceShopAccount := &service.ShopAccount{}
	total := rand.Intn(5) + 1
	list := make([]*model.ShopAccount, total, total)
	var err error
	for i := 0; i < total; i++ {
		obj := &model.ShopAccount{}
		obj.ShopId = shopId
		obj.Name = "name" + strconv.Itoa(index) + "_" + strconv.Itoa(i)
		obj.Password = "123123"
		obj.Password, err = util.EncryptPassword(obj.Password)
		checkErr(err)
		//obj.MoneyChance = rand.Intn(20)
		//obj.TotalMoney = rand.Intn(1000)
		obj.Locked = rand.Intn(2) < 1
		obj.Description = "描述" + strconv.Itoa(i)
		_, err := serviceShopAccount.Add(obj)
		checkErr(err)
		list[i] = obj
	}
	return list
}

func InitTestShopIntroductionImage(shopId string) []*model.ShopIntroductionImage {
	serviceShopIntroductionImage := &service.ShopIntroductionImage{}
	total := rand.Intn(10) + 1
	list := make([]*model.ShopIntroductionImage, total, total)
	for i := 0; i < total; i++ {
		obj := &model.ShopIntroductionImage{}
		obj.ShopId = shopId
		obj.Priority = rand.Intn(1000)
		obj.Hide = rand.Intn(2) < 1
		obj.ContentUrl = getRandomShopIntroductionImage()
		_, err := service.Add(serviceShopIntroductionImage, obj)
		checkErr(err)
		list[i] = obj
	}
	return list
}

func InitTestCashCoupon(shopId string) []*model.CashCoupon {
	serviceCashCoupon := &service.CashCoupon{}
	total := rand.Intn(10) + 1
	list := make([]*model.CashCoupon, total, total)
	for i := 0; i < total; i++ {
		obj := &model.CashCoupon{}
		obj.ShopId = shopId
		obj.Name = "代金券名称" + strconv.Itoa(i)
		obj.PreImage = getRandomCashCoupon()
		obj.DiscountAmount = 10 + rand.Intn(2000)
		obj.Price = 1 + rand.Intn(3)
		t := time.Now()
		if rand.Intn(2) < 1 {
			t = t.Add(time.Hour * time.Duration(rand.Int63n(10000)))
		} else {
			obj.Expired = true
			t = t.Add(-time.Hour * time.Duration(rand.Int63n(10000)))
		}
		obj.ExpiryDate = util.StringTime4DB(t)
		obj.Hide = rand.Intn(2) < 1
		obj.Priority = rand.Intn(1000)
		_, err := service.Add(serviceCashCoupon, obj)
		checkErr(err)
		list[i] = obj
	}
	return list
}

func InitTestPromotionalPartner() {
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

func InitTestUser() []*model.User {
	InitTestPromotionalPartner()
	servicePromotionalPartner := &service.PromotionalPartner{}
	argPromotionalPartner := &arg.PromotionalPartner{}
	promotionalPartnerList, err := service.Find(servicePromotionalPartner, argPromotionalPartner)
	checkErr(err)

	//InitTestPresent()
	//servicePresent := &service.Present{}
	//argPresent := &arg.Present{}
	//presentList, err := service.Find(servicePresent, argPresent)
	//checkErr(err)

	serviceUser := &service.User{}
	argUser := &arg.User{}
	count, err := service.Count(serviceUser, argUser)
	var list []*model.User
	checkErr(err)
	if count < 5 {
		checkErr(err)
		list = make([]*model.User, 300-100, 300-100)
		for i := 100; i < 300; i++ {
			obj := &model.User{}
			obj.Mobile = "12345678" + strconv.Itoa(i)
			obj.PresentChance = rand.Intn(10)
			obj.Locked = rand.Intn(2) < 1
			obj.Shared = rand.Intn(2) < 1
			if obj.Shared {
				obj.PresentChance = 1
			}
			if rand.Intn(2) < 1 {
				obj.PromotionalPartnerId = promotionalPartnerList[rand.Intn(len(promotionalPartnerList))].(*model.PromotionalPartner).Id
			}
			if i == 100 {
				obj.Mobile = specialMobile
				obj.PresentChance = 1000
				obj.Locked = false
				obj.Shared = false
			}
			_, err := service.Add(serviceUser, obj)
			checkErr(err)
			list[i-100] = obj
		}
	}
	return list
}

func InitTestPresent() []*model.Present {
	servicePresent := &service.Present{}
	argPresent := &arg.Present{}
	count, err := service.Count(servicePresent, argPresent)
	checkErr(err)
	var presentList []*model.Present
	if count < 5 {
		checkErr(err)
		presentList = make([]*model.Present, 100, 100)
		for i := 0; i < 100; i++ {
			obj := &model.Present{}
			obj.Name = "名称" + strconv.Itoa(i)
			obj.Weight = rand.Intn(1000)
			obj.Address = "地址" + strconv.Itoa(i)
			t := time.Now()
			if rand.Intn(2) < 1 {
				t = t.Add(time.Hour * time.Duration(rand.Int63n(10000)))
			} else {
				obj.Expired = true
				t = t.Add(-time.Hour * time.Duration(rand.Int63n(10000)))
			}
			obj.ExpiryDate = util.StringTime4DB(t)
			obj.Hide = rand.Intn(5) < 1
			obj.Image = getRandomPresentImage()
			//obj.Requirement = 10 + rand.Intn(2000)
			obj.Stock = 1000000
			_, err := service.Add(servicePresent, obj)
			checkErr(err)
			presentList[i] = obj
		}
	}
	return presentList
}

func InitTestRefundReason() {
	serviceRefundReason := &service.RefundReason{}
	argRefundReason := &arg.RefundReason{}
	count, err := service.Count(serviceRefundReason, argRefundReason)
	checkErr(err)
	if count < 5 {
		checkErr(err)
		for i := 0; i < 10; i++ {
			obj := &model.RefundReason{}
			obj.Description = "退款理由" + strconv.Itoa(i)
			obj.Priority = rand.Intn(1000)
			obj.Hide = rand.Intn(2) < 1
			_, err := service.Add(serviceRefundReason, obj)
			checkErr(err)
		}
	}
}

func InitTestCashCouponOrderInCart(userList []*model.User, shopList []*model.Shop) {
	serviceCashCoupon := &service.CashCoupon{}
	argCashCoupon := &arg.CashCoupon{}
	cashCouponList, err := service.Find(serviceCashCoupon, argCashCoupon)
	checkErr(err)

	serviceCashCouponOrder := &service.CashCouponOrder{}
	argCashCouponOrder := &arg.CashCouponOrder{}
	count, err := service.Count(serviceCashCouponOrder, argCashCouponOrder)
	checkErr(err)
	if count < 5 {
		for _, user := range userList {
			total := 0
			for i := 0; i < 10; i++ {
				if rand.Intn(3) < 1 {
					continue
				}
				total++
			}
			if user.Mobile == specialMobile {
				total = 1000
			}
			if total == 0 {
				continue
			}
			ids := make([]string, total, total)
			for i := 0; i < total; i++ {
				cashCoupon := cashCouponList[rand.Intn(len(cashCouponList))].(*model.CashCoupon)
				ids[i] = cashCoupon.Id
			}
			list, err := serviceCashCouponOrder.BatchAddInCart(user.Id, ids...)
			checkErr(err)
			for _, v2 := range list {
				if rand.Intn(5) < 4 {
					for _, v3 := range shopList {
						if v2.CashCoupon.ShopId == v3.Id {
							v2.CashCoupon.Shop = v3
							v2.User = user
							InitTestCashCouponOrderPayNotUsed(v2)
						}
					}
				}
			}
		}
	}
}

func InitTestCashCouponOrderPayNotUsed(cashCouponOrder *model.CashCouponOrder) {
	beginTime := time.Now().Add(-time.Duration(rand.Intn(1000))*time.Hour - time.Duration(rand.Intn(1000))*time.Minute - time.Duration(rand.Intn(1000))*time.Second)
	expiryDate, err := util.ParseTimeFromDB(cashCouponOrder.CashCoupon.ExpiryDate)
	checkErr(err)
	for beginTime.After(expiryDate) {
		beginTime = beginTime.Add(-time.Duration(rand.Intn(1000))*time.Hour - time.Duration(rand.Intn(1000))*time.Minute - time.Duration(rand.Intn(1000))*time.Second)
	}
	cashCouponOrder.PayOrderNumber = util.BuildUUID()
	cashCouponOrder.PayBegin = util.StringTime4DB(beginTime)
	cashCouponOrder.PrepayId = "test_prepay_id"
	cashCouponOrder.Status = constant.CashCouponOrderStatusPaidNotUsed
	cashCouponOrder.PayEnd = util.StringTime4DB(beginTime.Add(time.Duration(rand.Intn(100))*time.Second + 1))

	serviceCashCouponOrder := &service.CashCouponOrder{}
	_, err = service.Update(serviceCashCouponOrder, cashCouponOrder, "payOrderNumber", "payBegin", "payEnd", "prepayId", "status")
	checkErr(err)
	rm := rand.Intn(10)
	if rm < 5 {
		InitTestCashCouponOrderUsed(cashCouponOrder)
	} else if rm < 6 {
		InitTestCashCouponOrderRefunding(cashCouponOrder)
	} else if rm < 10 {
		InitTestCashCouponOrderRefunded(cashCouponOrder)
	}
}

func InitTestCashCouponOrderUsed(cashCouponOrder *model.CashCouponOrder) {
	serviceCashCouponOrder := &service.CashCouponOrder{}
	shopAccountId := cashCouponOrder.CashCoupon.Shop.ShopAccountList[rand.Intn(len(cashCouponOrder.CashCoupon.Shop.ShopAccountList))].Id
	if cashCouponOrder.CashCoupon.Expired {
		return
	}
	err := serviceCashCouponOrder.Complete(shopAccountId, cashCouponOrder.Id)
	checkErr(err)
	cashCouponOrder.Status = constant.CashCouponOrderStatusUsed
	if rand.Intn(2) < 1 {
		InitTestCashCouponOrderRefunded(cashCouponOrder)
	} else {
		InitTestCashCouponOrderRefunding(cashCouponOrder)
	}
}

func InitTestCashCouponOrderRefunding(cashCouponOrder *model.CashCouponOrder) {

}

func InitTestCashCouponOrderRefunded(cashCouponOrder *model.CashCouponOrder) {
	serviceRefund := &service.Refund{}
	serviceCashCouponOrder := &service.CashCouponOrder{}
	payEnd, err := util.ParseTimeFromDB(cashCouponOrder.PayEnd)
	checkErr(err)
	if cashCouponOrder.Status == constant.CashCouponOrderStatusPaidNotUsed {
		for i := 0; i < 2; i++ {
			if rand.Intn(5) < 1 {
				refund := buildSimpleTestRefundModel(cashCouponOrder.Price, constant.RefundStatusNotUsedRefundCancel, "", cashCouponOrder.Id)
				refund.RefundBegin = util.StringDefaultTime4DB()
				refund.RefundEnd = util.StringDefaultTime4DB()
				refund.RefundOrderNumber = ""
				_, err = service.Add(serviceRefund, refund)
				checkErr(err)
			}
		}
		for i := 0; i < 2; i++ {
			if rand.Intn(5) < 1 {
				refund := buildSimpleTestRefundModel(cashCouponOrder.Price, constant.RefundStatusNotUsedRefundFail, "", cashCouponOrder.Id)
				refund.RefundBegin = util.StringDefaultTime4DB()
				refund.RefundEnd = util.StringDefaultTime4DB()
				refund.RefundOrderNumber = ""
				_, err = service.Add(serviceRefund, refund)
				checkErr(err)
			}
		}
		if rand.Intn(2) < 1 {
			refund := buildSimpleTestRefundModel(cashCouponOrder.Price, constant.RefundStatusNotUsedRefundSuccess, "", cashCouponOrder.Id)
			refund.RefundOrderNumber = util.BuildUUID()
			beginTime := payEnd.Add(time.Duration(rand.Intn(10))*time.Hour + time.Duration(rand.Intn(10))*time.Minute + time.Duration(rand.Intn(100))*time.Second)
			refund.RefundBegin = util.StringTime4DB(beginTime)
			refund.RefundEnd = util.StringTime4DB(beginTime.Add(time.Duration(rand.Intn(100))*time.Second + 1))
			_, err = service.Add(serviceRefund, refund)
			checkErr(err)

			cashCouponOrder.Status = constant.CashCouponOrderStatusNotUsedRefunded
			cashCouponOrder.RefundAmount = refund.RefundAmount
			_, err = service.Update(serviceCashCouponOrder, cashCouponOrder, "status", "refundAmount")
			checkErr(err)
		}
	}
	if cashCouponOrder.Status == constant.CashCouponOrderStatusUsed {
		for i := 0; i < 2; i++ {
			if rand.Intn(5) < 1 {
				refund := buildSimpleTestRefundModel(cashCouponOrder.Price, constant.RefundStatusUsedRefundCancel, "", cashCouponOrder.Id)
				refund.RefundBegin = util.StringDefaultTime4DB()
				refund.RefundEnd = util.StringDefaultTime4DB()
				refund.RefundOrderNumber = ""
				refund.Evidence = getRandomEvidence()
				_, err = service.Add(serviceRefund, refund)
				checkErr(err)
			}
		}
		for i := 0; i < 2; i++ {
			if rand.Intn(5) < 1 {
				refund := buildSimpleTestRefundModel(cashCouponOrder.Price, constant.RefundStatusUsedRefundFail, "", cashCouponOrder.Id)
				refund.RefundBegin = util.StringDefaultTime4DB()
				refund.RefundEnd = util.StringDefaultTime4DB()
				refund.RefundOrderNumber = ""
				refund.Evidence = getRandomEvidence()
				_, err = service.Add(serviceRefund, refund)
				checkErr(err)
			}
		}
		for i := 0; i < 3; i++ {
			if cashCouponOrder.Price-cashCouponOrder.RefundAmount > 0 {
				r := rand.Intn(cashCouponOrder.Price - cashCouponOrder.RefundAmount)
				if r > 0 {
					refund := buildSimpleTestRefundModel(cashCouponOrder.Price, constant.RefundStatusUsedRefundSuccess, "", cashCouponOrder.Id)
					refund.RefundOrderNumber = util.BuildUUID()
					beginTime := payEnd.Add(time.Duration(rand.Intn(10))*time.Hour + time.Duration(rand.Intn(10))*time.Minute + time.Duration(rand.Intn(100))*time.Second)
					refund.RefundBegin = util.StringTime4DB(beginTime)
					refund.RefundEnd = util.StringTime4DB(beginTime.Add(time.Duration(rand.Intn(100))*time.Second + 1))
					refund.Evidence = getRandomEvidence()
					_, err = service.Add(serviceRefund, refund)
					checkErr(err)

					cashCouponOrder.RefundAmount = cashCouponOrder.RefundAmount + r
				}
			}
		}
		if cashCouponOrder.RefundAmount > 0 {
			cashCouponOrder.Status = constant.CashCouponOrderStatusUsedRefunded
			_, err = service.Update(serviceCashCouponOrder, cashCouponOrder, "status", "refundAmount")
			checkErr(err)
		}
	}

}
func buildSimpleTestRefundModel(refundAmount, status int, evidence, cashCouponOrderId string) *model.Refund {
	obj := &model.Refund{}
	obj.Comment = "平台意见" + strconv.Itoa(rand.Intn(1000))
	obj.CashCouponOrderId = cashCouponOrderId
	obj.Evidence = evidence
	obj.Reason = "其他退款理由" + strconv.Itoa(rand.Intn(1000))
	obj.RefundAmount = refundAmount
	obj.Status = status
	return obj
}

func InitTestPresentOrder(userList []*model.User) {
	servicePresent := &service.Present{}
	argPresent := &arg.Present{}
	argPresent.NotIncludeHide = true
	argPresent.BeforeExpiryDate = true
	presentList, err := service.Find(servicePresent, argPresent)
	checkErr(err)
	servicePresentOrder := &service.PresentOrder{}
	for _, user := range userList {
		if user.PresentChance == 0 || user.Locked {
			continue
		}
		total := rand.Intn(user.PresentChance)
		for i := 0; i < total; i++ {
			ids := make([]string, 12, 12)
			for j := 0; j < 12; j++ {
				ids[j] = presentList[rand.Intn(len(presentList))].(*model.Present).Id

			}
			presentOrder, err := servicePresentOrder.Random(user.Id, strings.Join(ids, constant.Split4Id))
			checkErr(err)
			if presentOrder != nil && rand.Intn(2) < 1 {
				err = servicePresentOrder.Complete(presentOrder.Id)
				checkErr(err)
			}
		}
	}
}

func InitTestCashCouponOrder() {
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
			InitTestRefund(obj)
			checkErr(err)
		}
	}
}

//func InitTestPresentOrder(userId, presentId string) {
//	servicePresentOrder := &service.PresentOrder{}
//	total := rand.Intn(10)
//	for i := 0; i < total; i++ {
//		obj := &model.PresentOrder{}
//		obj.PresentId = presentId
//		obj.UserId = userId
//		obj.Number = util.BuildPresentOrderNumber()
//		if rand.Intn(2) < 1 {
//			obj.Status = constant.PresentOrderStatusReceived
//		} else {
//			obj.Status = constant.PresentOrderStatusWaiting2Receive
//		}
//		_, err := service.Add(servicePresentOrder, obj)
//		checkErr(err)
//	}
//}

func InitTestRefund(cashCouponOrder *model.CashCouponOrder) {
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

	switch cashCouponOrder.Status {
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
		obj.RefundAmount = rand.Intn(cashCouponOrder.Price-cashCouponOrder.RefundAmount) + 1
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
	//	obj.Comment = "平台意见" + strconv.Itoa(rand.Intn(1000))
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
	obj.Comment = "平台意见" + strconv.Itoa(rand.Intn(1000))
	obj.CashCouponOrderId = cashCouponOrder.GetId()
	obj.Evidence = getRandomEvidence()
	obj.Reason = refundReasonList[rand.Intn(len(refundReasonList))].(*model.RefundReason).Description
	if rand.Intn(10) < 1 {
		obj.Reason = "其他退款理由" + strconv.Itoa(rand.Intn(1000))
	}
	obj.RefundAmount = cashCouponOrder.Price
	//obj.Status = status[rand.Intn(len(status))]
	return obj
}
