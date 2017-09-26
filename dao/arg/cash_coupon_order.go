package arg

import (
	"mimi/djq/model"
	"mimi/djq/util"
)

type CashCouponOrder struct {
	IdEqual        string
	IncludeDeleted bool
	OrderBy        string
	IdsIn          []string
	UserIdEqual    string `form:"userId" json:"userId"`
	UserIdsIn      []string

	NumberEqual string `form:"keyword" json:"keyword"`
	StatusEqual string `form:"status" json:"status"`
	StatusIn    []int

	CashCouponIdsIn []string
	NotComplete     bool

	PayBeginGT string
	PayBeginLT string

	PayOrderNumberEqual string `form:"payOrderNumber" json:"payOrderNumber"`

	PageSize   int `form:"pageSize" json:"pageSize"`
	TargetPage int `form:"targetPage" json:"targetPage"`

	DisplayNames []string

	UpdateObject interface{}
	UpdateNames  []string
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
	if arg.NumberEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " number = ?"
		params = append(params, arg.NumberEqual)
	}
	if arg.PayOrderNumberEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " pay_order_number = ?"
		params = append(params, arg.PayOrderNumberEqual)
	}
	if arg.StatusEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " status = ?"
		params = append(params, arg.StatusEqual)
	}
	if arg.StatusIn != nil && len(arg.StatusIn) != 0 {
		if sql != "" {
			sql += " and"
		}
		sql += " status in ("
		for i, status := range arg.StatusIn {
			if i != 0 {
				sql += ","
			}
			sql += "?"
			params = append(params, status)
		}
		sql += ")"
	}
	if arg.NotComplete {
		if sql != "" {
			sql += " and"
		}
		sql += " price > refund_amount"
	}
	if arg.UserIdsIn != nil && len(arg.UserIdsIn) != 0 {
		if sql != "" {
			sql += " and"
		}
		sql += " user_id in ("
		for i, userId := range arg.UserIdsIn {
			if i != 0 {
				sql += ","
			}
			sql += "?"
			params = append(params, userId)
		}
		sql += ")"
	}
	if arg.CashCouponIdsIn != nil && len(arg.CashCouponIdsIn) != 0 {
		if sql != "" {
			sql += " and"
		}
		sql += " cash_coupon_id in ("
		for i, cashCouponId := range arg.CashCouponIdsIn {
			if i != 0 {
				sql += ","
			}
			sql += "?"
			params = append(params, cashCouponId)
		}
		sql += ")"
	}
	if arg.PayBeginLT != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " pay_begin < ?"
		params = append(params, arg.PayBeginLT)
		sql += " and pay_begin != ?"
		params = append(params, util.StringDefaultTime4DB())
	}
	if arg.PayBeginGT != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " pay_begin > ?"
		params = append(params, arg.PayBeginGT)
	}
	if len(params) != 0 {
		sql = " where" + sql
	}
	return sql, params
}
