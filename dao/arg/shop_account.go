package arg

import (
	"mimi/djq/model"
)

type ShopAccount struct {
	IdEqual        string
	IncludeDeleted bool
	NameLike       string `form:"keyword" json:"keyword"`
	NameEqual      string
	OrderBy        string
	IdsIn          []string
	LockedOnly     bool
	UnlockedOnly   bool
	LockedEqual    string `form:"locked" json:"locked"`

	PasswordEqual  string
	ShopIdEqual    string `form:"shopId" json:"shopId"`

	PageSize       int `form:"pageSize" json:"pageSize"`
	TargetPage     int `form:"targetPage" json:"targetPage"`

	DisplayNames   []string

	UpdateObject   interface{}
	UpdateNames    []string
}

func (arg *ShopAccount) GetDisplayNames() []string {
	return arg.DisplayNames
}

func (arg *ShopAccount) GetModelInstance() model.BaseModelInterface {
	return &model.ShopAccount{}
}

func (arg *ShopAccount) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *ShopAccount) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *ShopAccount) GetUpdateNames() []string {
	return arg.UpdateNames
}

func (arg *ShopAccount) SetUpdateNames(updateNames []string) {
	arg.UpdateNames = updateNames
}

func (arg *ShopAccount) GetOrderBy() string {
	return arg.OrderBy
}

func (arg *ShopAccount) SetOrderBy(orderBy string) {
	arg.OrderBy = orderBy
}

func (arg *ShopAccount) GetIncludeDeleted() bool {
	return arg.IncludeDeleted
}

func (arg *ShopAccount) SetIncludeDeleted(includeDeleted bool) {
	arg.IncludeDeleted = includeDeleted
}

func (arg *ShopAccount) GetIdEqual() string {
	return arg.IdEqual
}

func (arg *ShopAccount) SetIdEqual(idEqual string) {
	arg.IdEqual = idEqual
}

func (arg *ShopAccount) GetUpdateObject() interface{} {
	return arg.UpdateObject
}

func (arg *ShopAccount) SetUpdateObject(updateObject interface{}) {
	arg.UpdateObject = updateObject
}

func (arg *ShopAccount) GetTargetPage() int {
	return arg.TargetPage
}

func (arg *ShopAccount) SetTargetPage(targetPage int) {
	arg.TargetPage = targetPage
}

func (arg *ShopAccount) GetPageSize() int {
	return arg.PageSize
}

func (arg *ShopAccount) SetPageSize(pageSize int) {
	arg.PageSize = pageSize
}

func (arg *ShopAccount) getCountConditions() (string, []interface{}) {
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
	if arg.PasswordEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " password = ?"
		params = append(params, arg.PasswordEqual)
	}
	if arg.LockedEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " locked = ?"
		params = append(params, arg.LockedEqual == "true")
	}
	if arg.LockedOnly || arg.UnlockedOnly {
		if sql != "" {
			sql += " and"
		}
		sql += " locked = ?"
		params = append(params, arg.LockedOnly)
	}
	if len(params) != 0 {
		sql = " where" + sql
	}
	return sql, params
}
