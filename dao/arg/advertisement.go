package arg

import (
	"mimi/djq/model"
)

type Advertisement struct {
	IdEqual        string
	IncludeDeleted bool
	NameLike       string
	NameEqual      string
	OrderBy        string
	IdsIn          []string

	NotIncludeHide bool

	PageSize       int `form:"pageSize" json:"pageSize"`
	TargetPage     int `form:"targetPage" json:"targetPage"`

	DisplayNames   []string

	UpdateObject   interface{}
	UpdateNames    []string
}

func (arg *Advertisement) GetDisplayNames() []string {
	return arg.DisplayNames
}

func (arg *Advertisement) GetModelInstance() model.BaseModelInterface {
	return &model.Advertisement{}
}

func (arg *Advertisement) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *Advertisement) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *Advertisement) GetUpdateNames() []string {
	return arg.UpdateNames
}

func (arg *Advertisement) SetUpdateNames(updateNames []string) {
	arg.UpdateNames = updateNames
}

func (arg *Advertisement) GetOrderBy() string {
	return arg.OrderBy
}

func (arg *Advertisement) SetOrderBy(orderBy string) {
	arg.OrderBy = orderBy
}

func (arg *Advertisement) GetIncludeDeleted() bool {
	return arg.IncludeDeleted
}

func (arg *Advertisement) SetIncludeDeleted(includeDeleted bool) {
	arg.IncludeDeleted = includeDeleted
}

func (arg *Advertisement) GetIdEqual() string {
	return arg.IdEqual
}

func (arg *Advertisement) SetIdEqual(idEqual string) {
	arg.IdEqual = idEqual
}

func (arg *Advertisement) GetUpdateObject() interface{} {
	return arg.UpdateObject
}

func (arg *Advertisement) SetUpdateObject(updateObject interface{}) {
	arg.UpdateObject = updateObject
}

func (arg *Advertisement) GetTargetPage() int {
	return arg.TargetPage
}

func (arg *Advertisement) SetTargetPage(targetPage int) {
	arg.TargetPage = targetPage
}

func (arg *Advertisement) GetPageSize() int {
	return arg.PageSize
}

func (arg *Advertisement) SetPageSize(pageSize int) {
	arg.PageSize = pageSize
}

func (arg *Advertisement) getCountConditions() (string, []interface{}) {
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
