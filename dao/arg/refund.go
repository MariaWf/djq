package arg

import (
	"mimi/djq/model"
)

type Refund struct {
	IdEqual              string
	IncludeDeleted       bool
	OrderBy              string
	IdsIn                []string
	CashCouponOrderIdsIn []string

	PageSize             int `form:"pageSize" json:"pageSize"`
	TargetPage           int `form:"targetPage" json:"targetPage"`

	DisplayNames         []string

	UpdateObject         interface{}
	UpdateNames          []string
}

func (arg *Refund) GetDisplayNames() []string {
	return arg.DisplayNames
}

func (arg *Refund) GetModelInstance() model.BaseModelInterface {
	return &model.Refund{}
}

func (arg *Refund) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *Refund) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *Refund) GetUpdateNames() []string {
	return arg.UpdateNames
}

func (arg *Refund) SetUpdateNames(updateNames []string) {
	arg.UpdateNames = updateNames
}

func (arg *Refund) GetOrderBy() string {
	return arg.OrderBy
}

func (arg *Refund) SetOrderBy(orderBy string) {
	arg.OrderBy = orderBy
}

func (arg *Refund) GetIncludeDeleted() bool {
	return arg.IncludeDeleted
}

func (arg *Refund) SetIncludeDeleted(includeDeleted bool) {
	arg.IncludeDeleted = includeDeleted
}

func (arg *Refund) GetIdEqual() string {
	return arg.IdEqual
}

func (arg *Refund) SetIdEqual(idEqual string) {
	arg.IdEqual = idEqual
}

func (arg *Refund) GetUpdateObject() interface{} {
	return arg.UpdateObject
}

func (arg *Refund) SetUpdateObject(updateObject interface{}) {
	arg.UpdateObject = updateObject
}

func (arg *Refund) GetTargetPage() int {
	return arg.TargetPage
}

func (arg *Refund) SetTargetPage(targetPage int) {
	arg.TargetPage = targetPage
}

func (arg *Refund) GetPageSize() int {
	return arg.PageSize
}

func (arg *Refund) SetPageSize(pageSize int) {
	arg.PageSize = pageSize
}

func (arg *Refund) getCountConditions() (string, []interface{}) {
	sql := ""
	params := make([]interface{}, 0, 9)
	if arg.CashCouponOrderIdsIn != nil && len(arg.CashCouponOrderIdsIn) != 0 {
		if sql != "" {
			sql += " and"
		}
		sql += " cash_coupon_order_id in ("
		for i, id := range arg.CashCouponOrderIdsIn {
			if i != 0 {
				sql += ","
			}
			sql += "?"
			params = append(params, id)
		}
		sql += ")"
	}
	if len(params) != 0 {
		sql = " where" + sql
	}
	return sql, params
}
