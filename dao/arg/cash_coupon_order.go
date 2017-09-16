package arg

import (
	"mimi/djq/model"
)

type CashCouponOrder struct {
	IdEqual        string
	IncludeDeleted bool
	OrderBy        string
	IdsIn          []string
	UserIdEqual    string `form:"userId" json:"userId"`

	StatusEqual    string `form:"status" json:"status"`

	PageSize       int `form:"pageSize" json:"pageSize"`
	TargetPage     int `form:"targetPage" json:"targetPage"`

	DisplayNames   []string

	UpdateObject   interface{}
	UpdateNames    []string
}

func (arg *CashCouponOrder) GetDisplayNames() []string {
	return arg.DisplayNames
}

func (arg *CashCouponOrder) GetModelInstance() model.BaseModelInterface {
	return &model.CashCouponOrder{}
}

func (arg *CashCouponOrder) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *CashCouponOrder) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *CashCouponOrder) GetUpdateNames() []string {
	return arg.UpdateNames
}

func (arg *CashCouponOrder) SetUpdateNames(updateNames []string) {
	arg.UpdateNames = updateNames
}

func (arg *CashCouponOrder) GetOrderBy() string {
	return arg.OrderBy
}

func (arg *CashCouponOrder) SetOrderBy(orderBy string) {
	arg.OrderBy = orderBy
}

func (arg *CashCouponOrder) GetIncludeDeleted() bool {
	return arg.IncludeDeleted
}

func (arg *CashCouponOrder) SetIncludeDeleted(includeDeleted bool) {
	arg.IncludeDeleted = includeDeleted
}

func (arg *CashCouponOrder) GetIdEqual() string {
	return arg.IdEqual
}

func (arg *CashCouponOrder) SetIdEqual(idEqual string) {
	arg.IdEqual = idEqual
}

func (arg *CashCouponOrder) GetUpdateObject() interface{} {
	return arg.UpdateObject
}

func (arg *CashCouponOrder) SetUpdateObject(updateObject interface{}) {
	arg.UpdateObject = updateObject
}

func (arg *CashCouponOrder) GetTargetPage() int {
	return arg.TargetPage
}

func (arg *CashCouponOrder) SetTargetPage(targetPage int) {
	arg.TargetPage = targetPage
}

func (arg *CashCouponOrder) GetPageSize() int {
	return arg.PageSize
}

func (arg *CashCouponOrder) SetPageSize(pageSize int) {
	arg.PageSize = pageSize
}

func (arg *CashCouponOrder) getCountConditions() (string, []interface{}) {
	sql := ""
	params := make([]interface{}, 0, 9)
	if arg.UserIdEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " user_id = ?"
		params = append(params, arg.UserIdEqual)
	}
	if arg.StatusEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " status = ?"
		params = append(params, arg.StatusEqual)
	}
	if len(params) != 0 {
		sql = " where" + sql
	}
	return sql, params
}
