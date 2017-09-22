package model

import "github.com/pkg/errors"

type CashCouponOrder struct {
	Id             string `form:"id" json:"id" db:"id" desc:"id"`
	UserId         string `form:"userId" json:"userId" db:"user_id" desc:"用户ID"`
	CashCouponId   string `form:"cashCouponId" json:"cashCouponId" db:"cash_coupon_id" desc:"代金券ID"`
	Price          int `form:"price" json:"price" db:"price" desc:"价格"`
	RefundAmount   int `form:"refundAmount" json:"refundAmount" db:"refund_amount" desc:"累计退款金额"`
	PayOrderNumber string `form:"payOrderNumber" json:"payOrderNumber" db:"pay_order_number" desc:"支付订单编码"`
	PrepayId       string `form:"prepayId" json:"prepayId" db:"prepay_id" desc:"微信预支付ID"`
	Number         string `form:"number" json:"number" db:"name" desc:"编码"`
	Status         int    `form:"status" json:"status" db:"status" desc:"状态"`
	PayBegin       string   `form:"payBegin" json:"payBegin" db:"pay_begin" desc:"支付开始日期" time_format:"2006-01-02" time_utc:"1"`
	PayEnd         string   `form:"payEnd" json:"payEnd" db:"pay_end" desc:"支付结束日期" time_format:"2006-01-02" time_utc:"1"`

	CashCoupon     *CashCoupon `form:"cashCoupon" json:"cashCoupon"  desc:"代金券"`
	User           *User `form:"user" json:"user"  desc:"用户"`
}

func (obj *CashCouponOrder) GetId() string {
	return obj.Id
}

func (obj *CashCouponOrder) SetId(id string) {
	obj.Id = id
}

func (obj *CashCouponOrder) GetTableName() string {
	return "tbl_cash_coupon_order"
}

func (obj *CashCouponOrder) GetDBNames() []string {
	return []string{"id", "user_id", "cash_coupon_id", "price", "refund_amount", "pay_order_number", "prepay_id", "number", "status", "pay_begin", "pay_end"}
}

func (obj *CashCouponOrder) GetMapNames() []string {
	return []string{"id", "userId", "cashCouponId", "price", "refundAmount", "payOrderNumber", "prepayId", "number", "status", "payBegin", "payEnd"}
}

func (obj *CashCouponOrder) GetValue4Map(name string) interface{} {
	switch name {
	case "id": return obj.Id
	case "userId": return obj.UserId
	case "cashCouponId": return obj.CashCouponId
	case "price": return obj.Price
	case "refundAmount": return obj.RefundAmount
	case "payOrderNumber": return obj.PayOrderNumber
	case "prepayId": return obj.PrepayId
	case "number": return obj.Number
	case "status": return obj.Status
	case "payBegin": return obj.PayBegin
	case "payEnd": return obj.PayEnd
	}
	panic(errors.New("对象cashCouponOrder属性[" + name + "]不存在"))
}

func (obj *CashCouponOrder) GetDBFromMapName(name string) string {
	str := GetDBFromMapName(obj, name)
	if str != "" {
		return str
	}
	panic(errors.New("对象cashCouponOrder属性[" + name + "]不存在"))
}

func (obj *CashCouponOrder) GetPointer4DB(name string) interface{} {
	switch name {
	case "id": return &obj.Id
	case "user_id": return &obj.UserId
	case "cash_coupon_id": return &obj.CashCouponId
	case "price": return &obj.Price
	case "refund_amount": return &obj.RefundAmount
	case "pay_order_number": return &obj.PayOrderNumber
	case "prepay_id": return &obj.PrepayId
	case "number": return &obj.Number
	case "status": return &obj.Status
	case "pay_begin": return &obj.PayBegin
	case "pay_end": return &obj.PayEnd
	}
	panic(errors.New("对象cashCouponOrder属性[" + name + "]不存在"))
}

func (obj *CashCouponOrder) GetValue4DB(name string) interface{} {
	switch name {
	case "id": return obj.Id
	case "user_id": return obj.UserId
	case "cash_coupon_id": return obj.CashCouponId
	case "price": return obj.Price
	case "refund_amount": return obj.RefundAmount
	case "pay_order_number": return obj.PayOrderNumber
	case "prepay_id": return obj.PrepayId
	case "number": return obj.Number
	case "status": return obj.Status
	case "pay_begin": return obj.PayBegin
	case "pay_end": return obj.PayEnd
	}
	panic(errors.New("对象cashCouponOrder属性[" + name + "]不存在"))
}