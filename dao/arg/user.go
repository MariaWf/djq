package arg

import (
	"mimi/djq/model"
)

type User struct {
	IdEqual                   string
	IncludeDeleted            bool
	PromotionalPartnerIdEqual string
	MobileEqual               string
	OrderBy                   string
	IdsIn                     []string
	LockedOnly                bool
	UnlockedOnly              bool

	PageSize                  int `form:"pageSize" json:"pageSize"`
	TargetPage                int `form:"targetPage" json:"targetPage"`

	DisplayNames              []string

	UpdateObject              interface{}
	UpdateNames               []string
}

func (arg *User) GetDisplayNames() []string {
	return arg.DisplayNames
}

func (arg *User) GetModelInstance() model.BaseModelInterface {
	return &model.User{}
}

func (arg *User) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *User) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *User) GetUpdateNames() []string {
	return arg.UpdateNames
}

func (arg *User) SetUpdateNames(updateNames []string) {
	arg.UpdateNames = updateNames
}

func (arg *User) GetOrderBy() string {
	return arg.OrderBy
}

func (arg *User) SetOrderBy(orderBy string) {
	arg.OrderBy = orderBy
}

func (arg *User) GetIncludeDeleted() bool {
	return arg.IncludeDeleted
}

func (arg *User) SetIncludeDeleted(includeDeleted bool) {
	arg.IncludeDeleted = includeDeleted
}

func (arg *User) GetIdEqual() string {
	return arg.IdEqual
}

func (arg *User) SetIdEqual(idEqual string) {
	arg.IdEqual = idEqual
}

func (arg *User) GetUpdateObject() interface{} {
	return arg.UpdateObject
}

func (arg *User) SetUpdateObject(updateObject interface{}) {
	arg.UpdateObject = updateObject
}

func (arg *User) GetTargetPage() int {
	return arg.TargetPage
}

func (arg *User) SetTargetPage(targetPage int) {
	arg.TargetPage = targetPage
}

func (arg *User) GetPageSize() int {
	return arg.PageSize
}

func (arg *User) SetPageSize(pageSize int) {
	arg.PageSize = pageSize
}

func (arg *User) getCountConditions() (string, []interface{}) {
	sql := ""
	params := make([]interface{}, 0, 9)
	if arg.MobileEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " name = ?"
		params = append(params, arg.MobileEqual)
	}
	if arg.PromotionalPartnerIdEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " password = ?"
		params = append(params, arg.PromotionalPartnerIdEqual)
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
