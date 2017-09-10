package model

import (
	"github.com/pkg/errors"
)

type CashCoupon struct {
	Id             string `form:"id" json:"id" db:"id" desc:"id"`
	ShopId         string `form:"shopId" json:"shopId" db:"shop_id" desc:"shopId"`
	Name           string `form:"name" json:"name" db:"name" desc:"名称"`
	PreImage       string `form:"preImage" json:"preImage" db:"pre_image" desc:"图片"`
	DiscountAmount int `form:"discountAmount" json:"discountAmount" db:"discount_amount" desc:"优惠额度"`
	ExpiryDate     string   `form:"expiryDate" json:"expiryDate" db:"expiry_date" desc:"有效日期" time_format:"2006-01-02" time_utc:"1"`
	Expired        bool `form:"expired" json:"expired" db:"expired" desc:"已过期"`
	Hide           bool   `form:"hide" json:"hide" db:"hide" desc:"隐藏"`
	Priority       int    `form:"priority" json:"priority" db:"priority" desc:"优先权重"`
}

func (obj *CashCoupon) GetId() string {
	return obj.Id
}

func (obj *CashCoupon) SetId(id string) {
	obj.Id = id
}

func (obj *CashCoupon) GetTableName() string {
	return "tbl_cash_coupon"
}

func (obj *CashCoupon) GetDBNames() []string {
	return []string{"id", "shop_id", "name", "pre_image", "discount_amount", "expiry_date", "expired", "hide", "priority"}

}

func (obj *CashCoupon) GetMapNames() []string {
	return []string{"id", "shopId", "name", "preImage", "discountAmount", "expiryDate", "expired", "hide", "priority"}
}

func (obj *CashCoupon) GetValue4Map(name string) interface{} {
	switch name {
	case "id": return obj.Id
	case "shopId": return obj.ShopId
	case "name": return obj.Name
	case "preImage": return obj.PreImage
	case "discountAmount": return obj.DiscountAmount
	case "expiryDate": return obj.ExpiryDate
	case "expired": return obj.Expired
	case "hide": return obj.Hide
	case "priority": return obj.Priority
	}
	panic(errors.New("对象cashCoupon属性[" + name + "]不存在"))
}

func (obj *CashCoupon) GetDBFromMapName(name string) string {
	str := GetDBFromMapName(obj, name)
	if str != "" {
		return str
	}
	panic(errors.New("对象cashCoupon属性[" + name + "]不存在"))
}

func (obj *CashCoupon) GetPointer4DB(name string) interface{} {
	switch name {
	case "id": return &obj.Id
	case "shop_id": return &obj.ShopId
	case "name": return &obj.Name
	case "pre_image": return &obj.PreImage
	case "discount_amount": return &obj.DiscountAmount
	case "expiry_date": return &obj.ExpiryDate
	case "expired": return &obj.Expired
	case "hide": return &obj.Hide
	case "priority": return &obj.Priority
	}
	panic(errors.New("对象cashCoupon属性[" + name + "]不存在"))
}

func (obj *CashCoupon) GetValue4DB(name string) interface{} {
	switch name {
	case "id": return obj.Id
	case "shop_id": return obj.ShopId
	case "name": return obj.Name
	case "pre_image": return obj.PreImage
	case "discount_amount": return obj.DiscountAmount
	case "expiry_date": return obj.ExpiryDate
	case "expired": return obj.Expired
	case "hide": return obj.Hide
	case "priority": return obj.Priority
	}
	panic(errors.New("对象cashCoupon属性[" + name + "]不存在"))
}