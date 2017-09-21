package arg

import (
	"mimi/djq/model"
	"time"
	"mimi/djq/util"
)

type Present struct {
	IdEqual          string
	IncludeDeleted   bool
	NameLike         string `form:"keyword" json:"keyword"`
	NameEqual        string
	OrderBy          string
	IdsIn            []string

	Enough           bool
	NotIncludeHide   bool
	OverExpiryDate   bool
	BeforeExpiryDate bool

	PageSize         int `form:"pageSize" json:"pageSize"`
	TargetPage       int `form:"targetPage" json:"targetPage"`

	DisplayNames     []string

	UpdateObject     interface{}
	UpdateNames      []string
}

func (arg *Present) GetDisplayNames() []string {
	return arg.DisplayNames
}

func (arg *Present) GetModelInstance() model.BaseModelInterface {
	return &model.Present{}
}

func (arg *Present) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *Present) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *Present) GetUpdateNames() []string {
	return arg.UpdateNames
}

func (arg *Present) SetUpdateNames(updateNames []string) {
	arg.UpdateNames = updateNames
}

func (arg *Present) GetOrderBy() string {
	return arg.OrderBy
}

func (arg *Present) SetOrderBy(orderBy string) {
	arg.OrderBy = orderBy
}

func (arg *Present) GetIncludeDeleted() bool {
	return arg.IncludeDeleted
}

func (arg *Present) SetIncludeDeleted(includeDeleted bool) {
	arg.IncludeDeleted = includeDeleted
}

func (arg *Present) GetIdEqual() string {
	return arg.IdEqual
}

func (arg *Present) SetIdEqual(idEqual string) {
	arg.IdEqual = idEqual
}

func (arg *Present) GetUpdateObject() interface{} {
	return arg.UpdateObject
}

func (arg *Present) SetUpdateObject(updateObject interface{}) {
	arg.UpdateObject = updateObject
}

func (arg *Present) GetTargetPage() int {
	return arg.TargetPage
}

func (arg *Present) SetTargetPage(targetPage int) {
	arg.TargetPage = targetPage
}

func (arg *Present) GetPageSize() int {
	return arg.PageSize
}

func (arg *Present) SetPageSize(pageSize int) {
	arg.PageSize = pageSize
}

func (arg *Present) getCountConditions() (string, []interface{}) {
	sql := ""
	params := make([]interface{}, 0, 9)
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
	if arg.Enough{
		if sql != "" {
			sql += " and"
		}
		sql += " stock > requirement"
	}
	if len(params) != 0 {
		sql = " where" + sql
	}
	return sql, params
}
