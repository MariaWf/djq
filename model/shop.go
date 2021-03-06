package model

import (
	"github.com/pkg/errors"
)

type Shop struct {
	Id                    string `form:"id" json:"id" db:"id" desc:"id"`
	Name                  string `form:"name" json:"name" db:"name" desc:"名称"`
	Logo                  string `form:"logo" json:"logo" db:"logo" desc:"商标"`
	PreImage              string `form:"preImage" json:"preImage" db:"pre_image" desc:"预览图"`
	TotalCashCouponNumber int    `form:"totalCashCouponNumber" json:"totalCashCouponNumber" db:"total_cash_coupon_number" desc:"累计优惠次数"`
	TotalCashCouponPrice  int    `form:"totalCashCouponPrice" json:"totalCashCouponPrice" db:"total_cash_coupon_price" desc:"累计优惠金额"`
	Introduction          string `form:"introduction" json:"introduction" db:"introduction" desc:"介绍"`
	Address               string `form:"address" json:"address" db:"address" desc:"地址"`
	Priority              int    `form:"priority" json:"priority" db:"priority" desc:"优先权重"`
	Hide                  bool   `form:"hide" json:"hide" db:"hide" desc:"隐藏"`
	TitleFirst                  string `form:"titleFirst" json:"titleFirst" db:"title_first" desc:"一级标题"`
	TitleSecond                  string `form:"titleSecond" json:"titleSecond" db:"title_second" desc:"二级标题"`
	PhoneNumber                  string `form:"phoneNumber" json:"phoneNumber" db:"phone_number" desc:"电话"`

	ShopClassificationList    []*ShopClassification    `form:"shopClassificationList" json:"shopClassificationList" desc:"商家分类"`
	ShopIntroductionImageList []*ShopIntroductionImage `form:"shopIntroductionImageList" json:"shopIntroductionImageList" desc:"商店介绍图"`
	CashCouponList            []*CashCoupon            `form:"cashCouponList" json:"cashCouponList" desc:"代金券"`
	ShopAccountList           []*ShopAccount           `form:"shopAccountList" json:"shopAccountList" desc:"商家账户"`
}

func (obj *Shop) GetId() string {
	return obj.Id
}

func (obj *Shop) SetId(id string) {
	obj.Id = id
}

func (obj *Shop) GetTableName() string {
	return "tbl_shop"
}

func (obj *Shop) GetDBNames() []string {
	return []string{"id", "name", "logo", "pre_image", "total_cash_coupon_number", "total_cash_coupon_price", "introduction", "address", "priority", "hide","title_first","title_second","phone_number"}
}

func (obj *Shop) GetMapNames() []string {
	return []string{"id", "name", "logo", "preImage", "totalCashCouponNumber", "totalCashCouponPrice", "introduction", "address", "priority", "hide","titleFirst","titleSecond","phoneNumber"}
}

func (obj *Shop) GetValue4Map(name string) interface{} {
	switch name {
	case "id":
		return obj.Id
	case "name":
		return obj.Name
	case "logo":
		return obj.Logo
	case "preImage":
		return obj.PreImage
	case "totalCashCouponNumber":
		return obj.TotalCashCouponNumber
	case "totalCashCouponPrice":
		return obj.TotalCashCouponPrice
	case "introduction":
		return obj.Introduction
	case "address":
		return obj.Address
	case "priority":
		return obj.Priority
	case "hide":
		return obj.Hide
	case "titleFirst":
		return obj.TitleFirst
	case "titleSecond":
		return obj.TitleSecond
	case "phoneNumber":
		return obj.PhoneNumber
	}
	panic(errors.New("对象shop属性[" + name + "]不存在"))
}

func (obj *Shop) GetDBFromMapName(name string) string {
	str := GetDBFromMapName(obj, name)
	if str != "" {
		return str
	}
	panic(errors.New("对象shop属性[" + name + "]不存在"))
}

func (obj *Shop) GetPointer4DB(name string) interface{} {
	switch name {
	case "id":
		return &obj.Id
	case "name":
		return &obj.Name
	case "logo":
		return &obj.Logo
	case "pre_image":
		return &obj.PreImage
	case "total_cash_coupon_number":
		return &obj.TotalCashCouponNumber
	case "total_cash_coupon_price":
		return &obj.TotalCashCouponPrice
	case "introduction":
		return &obj.Introduction
	case "address":
		return &obj.Address
	case "priority":
		return &obj.Priority
	case "hide":
		return &obj.Hide
	case "title_first":
		return &obj.TitleFirst
	case "title_second":
		return &obj.TitleSecond
	case "phone_number":
		return &obj.PhoneNumber
	}
	panic(errors.New("对象shop属性[" + name + "]不存在"))
}

func (obj *Shop) GetValue4DB(name string) interface{} {
	switch name {
	case "id":
		return obj.Id
	case "name":
		return obj.Name
	case "logo":
		return obj.Logo
	case "pre_image":
		return obj.PreImage
	case "total_cash_coupon_number":
		return obj.TotalCashCouponNumber
	case "total_cash_coupon_price":
		return obj.TotalCashCouponPrice
	case "introduction":
		return obj.Introduction
	case "address":
		return obj.Address
	case "priority":
		return obj.Priority
	case "hide":
		return obj.Hide
	case "title_first":
		return obj.TitleFirst
	case "title_second":
		return obj.TitleSecond
	case "phone_number":
		return obj.PhoneNumber
	}
	panic(errors.New("对象shop属性[" + name + "]不存在"))
}

func (obj *Shop) SetShopClassificationListFromInterfaceArr(list []interface{}) {
	if list != nil && len(list) != 0 {
		obj.ShopClassificationList = make([]*ShopClassification, len(list), len(list))
		for i, shopClassification := range list {
			obj.ShopClassificationList[i] = shopClassification.(*ShopClassification)
		}
	} else {
		obj.ShopClassificationList = make([]*ShopClassification, 0, 1)
	}
}

func (obj *Shop) SetShopIntroductionImageListFromInterfaceArr(list []interface{}) {
	if list != nil && len(list) != 0 {
		obj.ShopIntroductionImageList = make([]*ShopIntroductionImage, len(list), len(list))
		for i, shopIntroductionImage := range list {
			obj.ShopIntroductionImageList[i] = shopIntroductionImage.(*ShopIntroductionImage)
		}
	} else {
		obj.ShopIntroductionImageList = make([]*ShopIntroductionImage, 0, 1)
	}
}

func (obj *Shop) SetCashCouponListFromInterfaceArr(list []interface{}) {
	if list != nil && len(list) != 0 {
		obj.CashCouponList = make([]*CashCoupon, len(list), len(list))
		for i, cashCoupon := range list {
			obj.CashCouponList[i] = cashCoupon.(*CashCoupon)
		}
	} else {
		obj.CashCouponList = make([]*CashCoupon, 0, 1)
	}
}
