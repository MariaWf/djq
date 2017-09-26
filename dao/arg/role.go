package arg

import (
	"mimi/djq/model"
)

type Role struct {
	IdEqual        string
	IncludeDeleted bool
	NameLike       string `form:"keyword" json:"keyword"`
	NameEqual      string
	OrderBy        string
	IdsIn          []string
	IdsNotIn       []string

	PageSize   int `form:"pageSize" json:"pageSize"`
	TargetPage int `form:"targetPage" json:"targetPage"`

	DisplayNames []string

	UpdateObject interface{}
	UpdateNames  []string
}

func (arg *Role) GetDisplayNames() []string {
	return arg.DisplayNames
}

func (arg *Role) GetModelInstance() model.BaseModelInterface {
	return &model.Role{}
}

func (arg *Role) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *Role) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *Role) GetUpdateNames() []string {
	return arg.UpdateNames
}

func (arg *Role) SetUpdateNames(updateNames []string) {
	arg.UpdateNames = updateNames
}

func (arg *Role) GetOrderBy() string {
	return arg.OrderBy
}

func (arg *Role) SetOrderBy(orderBy string) {
	arg.OrderBy = orderBy
}

func (arg *Role) GetIncludeDeleted() bool {
	return arg.IncludeDeleted
}

func (arg *Role) SetIncludeDeleted(includeDeleted bool) {
	arg.IncludeDeleted = includeDeleted
}

func (arg *Role) GetIdEqual() string {
	return arg.IdEqual
}

func (arg *Role) SetIdEqual(idEqual string) {
	arg.IdEqual = idEqual
}

func (arg *Role) GetUpdateObject() interface{} {
	return arg.UpdateObject
}

func (arg *Role) SetUpdateObject(updateObject interface{}) {
	arg.UpdateObject = updateObject
}

func (arg *Role) GetTargetPage() int {
	return arg.TargetPage
}

func (arg *Role) SetTargetPage(targetPage int) {
	arg.TargetPage = targetPage
}

func (arg *Role) GetPageSize() int {
	return arg.PageSize
}

func (arg *Role) SetPageSize(pageSize int) {
	arg.PageSize = pageSize
}

func (arg *Role) getCountConditions() (string, []interface{}) {
	sql := ""
	params := make([]interface{}, 0, 9)
	if arg.NameLike != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " name like ?"
		params = append(params, "%"+arg.NameLike+"%")
	}
	if arg.NameEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " name = ?"
		params = append(params, arg.NameEqual)
	}
	if arg.IdsNotIn != nil && len(arg.IdsNotIn) != 0 {
		if sql != "" {
			sql += " and"
		}
		sql += " id not in ("
		for i, id := range arg.IdsNotIn {
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
