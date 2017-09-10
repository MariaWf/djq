package arg

import (
	"mimi/djq/model"
	"time"
	"mimi/djq/util"
)

type CashCoupon struct {
	IdEqual          string
	IncludeDeleted   bool
	NameLike         string
	NameEqual        string
	OrderBy          string
	IdsIn            []string

	ExpiredOnly      bool
	UnexpiredOnly    bool
	NotIncludeHide   bool
	ShopIdEqual      string `form:"shopId" json:"shopId"`

	OverExpiryDate   bool
	BeforeExpiryDate bool

	PageSize         int `form:"pageSize" json:"pageSize"`
	TargetPage       int `form:"targetPage" json:"targetPage"`

	DisplayNames     []string

	UpdateObject     interface{}
	UpdateNames      []string
}

func (arg *CashCoupon) GetDisplayNames() []string {
	return arg.DisplayNames
}

func (arg *CashCoupon) GetModelInstance() model.BaseModelInterface {
	return &model.CashCoupon{}
}

func (arg *CashCoupon) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *CashCoupon) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *CashCoupon) GetUpdateNames() []string {
	return arg.UpdateNames
}

func (arg *CashCoupon) SetUpdateNames(updateNames []string) {
	arg.UpdateNames = updateNames
}

func (arg *CashCoupon) GetOrderBy() string {
	return arg.OrderBy
}

func (arg *CashCoupon) SetOrderBy(orderBy string) {
	arg.OrderBy = orderBy
}

func (arg *CashCoupon) GetIncludeDeleted() bool {
	return arg.IncludeDeleted
}

func (arg *CashCoupon) SetIncludeDeleted(includeDeleted bool) {
	arg.IncludeDeleted = includeDeleted
}

func (arg *CashCoupon) GetIdEqual() string {
	return arg.IdEqual
}

func (arg *CashCoupon) SetIdEqual(idEqual string) {
	arg.IdEqual = idEqual
}

func (arg *CashCoupon) GetUpdateObject() interface{} {
	return arg.UpdateObject
}

func (arg *CashCoupon) SetUpdateObject(updateObject interface{}) {
	arg.UpdateObject = updateObject
}

func (arg *CashCoupon) GetTargetPage() int {
	return arg.TargetPage
}

func (arg *CashCoupon) SetTargetPage(targetPage int) {
	arg.TargetPage = targetPage
}

func (arg *CashCoupon) GetPageSize() int {
	return arg.PageSize
}

func (arg *CashCoupon) SetPageSize(pageSize int) {
	arg.PageSize = pageSize
}

func (arg *CashCoupon) getCountConditions() (string, []interface{}) {
	sql := ""
	params := make([]interface{}, 0, 9)
	if arg.ShopIdEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " shop_id = ?"
		params = append(params, arg.ShopIdEqual)
	}
	if arg.NameLike != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " name like ?"
		params = append(params, "%" + arg.NameLike + "%")
	}
	if arg.NameEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " name = ?"
		params = append(params, arg.NameEqual)
	}
	if arg.ExpiredOnly || arg.UnexpiredOnly {
		if sql != "" {
			sql += " and"
		}
		sql += " expired = ?"
		params = append(params, arg.ExpiredOnly)
	}
	if arg.OverExpiryDate {
		if sql != "" {
			sql += " and"
		}
		sql += " expiry_date <= ?"
		params = append(params, util.StringTime4DB(time.Now()))
	}
	if arg.BeforeExpiryDate {
		if sql != "" {
			sql += " and"
		}
		sql += " expiry_date > ?"
		params = append(params, util.StringTime4DB(time.Now()))
	}
	if arg.NotIncludeHide {
		if sql != "" {
			sql += " and"
		}
		sql += " hide = ?"
		params = append(params, false)
	}
	if len(params) != 0 {
		sql = " where" + sql
	}
	return sql, params
}
