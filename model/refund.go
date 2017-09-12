package model

import "github.com/pkg/errors"

type Refund struct {
	Id                string `form:"id" json:"id" db:"id" desc:"id"`
	CashCouponOrderId string `form:"cashCouponOrderId" json:"cashCouponOrderId" db:"cash_coupon_order_id" desc:"代金券订单ID"`
	Evidence          string `form:"evidence" json:"evidence" db:"evidence" desc:"退款凭证"`
	Reason            string `form:"reason" json:"reason" db:"reason" desc:"退款理由"`
	Common            string `form:"comment" json:"comment" db:"comment" desc:"平台意见"`
	RefundAmount      int `form:"refundAmount" json:"refundAmount" db:"refund_amount" desc:"退款金额"`
	Status            int    `form:"status" json:"status" db:"status" desc:"状态"`
}

func (obj *Refund) GetId() string {
	return obj.Id
}

func (obj *Refund) SetId(id string) {
	obj.Id = id
}

func (obj *Refund) GetTableName() string {
	return "tbl_refund"
}

func (obj *Refund) GetDBNames() []string {
	return []string{"id", "cash_coupon_order_id", "evidence", "reason", "comment", "refund_amount", "status"}
}

func (obj *Refund) GetMapNames() []string {
	return []string{"id", "cashCouponOrderId", "evidence", "reason", "comment", "refundAmount", "status"}
}

func (obj *Refund) GetValue4Map(name string) interface{} {
	switch name {
	case "id": return obj.Id
	case "cashCouponOrderId": return obj.CashCouponOrderId
	case "evidence": return obj.Evidence
	case "reason": return obj.Reason
	case "comment": return obj.Common
	case "refundAmount": return obj.RefundAmount
	case "status": return obj.Status
	}
	panic(errors.New("对象refund属性[" + name + "]不存在"))
}

func (obj *Refund) GetDBFromMapName(name string) string {
	str := GetDBFromMapName(obj, name)
	if str != "" {
		return str
	}
	panic(errors.New("对象refund属性[" + name + "]不存在"))
}

func (obj *Refund) GetPointer4DB(name string) interface{} {
	switch name {
	case "id": return &obj.Id
	case "cash_coupon_order_id": return &obj.CashCouponOrderId
	case "evidence": return &obj.Evidence
	case "reason": return &obj.Reason
	case "comment": return &obj.Common
	case "refund_amount": return &obj.RefundAmount
	case "status": return &obj.Status
	}
	panic(errors.New("对象refund属性[" + name + "]不存在"))
}

func (obj *Refund) GetValue4DB(name string) interface{} {
	switch name {
	case "id": return obj.Id
	case "cash_coupon_order_id": return obj.CashCouponOrderId
	case "evidence": return obj.Evidence
	case "reason": return obj.Reason
	case "comment": return obj.Common
	case "refund_amount": return obj.RefundAmount
	case "status": return obj.Status
	}
	panic(errors.New("对象refund属性[" + name + "]不存在"))
}