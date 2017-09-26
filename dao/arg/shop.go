package arg

import (
	"mimi/djq/model"
)

type Shop struct {
	IdEqual        string
	IncludeDeleted bool
	NameLike       string `form:"keyword" json:"keyword"`
	NameEqual      string
	OrderBy        string
	IdsIn          []string

	NotIncludeHide bool
	HideEqual      string `form:"hide" json:"hide"`

	PageSize   int `form:"pageSize" json:"pageSize"`
	TargetPage int `form:"targetPage" json:"targetPage"`

	DisplayNames []string

	UpdateObject interface{}
	UpdateNames  []string
}

func (arg *Shop) GetDisplayNames() []string {
	return arg.DisplayNames
}

func (arg *Shop) GetModelInstance() model.BaseModelInterface {
	return &model.Shop{}
}

func (arg *Shop) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *Shop) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *Shop) GetUpdateNames() []string {
	return arg.UpdateNames
}

func (arg *Shop) SetUpdateNames(updateNames []string) {
	arg.UpdateNames = updateNames
}

func (arg *Shop) GetOrderBy() string {
	return arg.OrderBy
}

func (arg *Shop) SetOrderBy(orderBy string) {
	arg.OrderBy = orderBy
}

func (arg *Shop) GetIncludeDeleted() bool {
	return arg.IncludeDeleted
}

func (arg *Shop) SetIncludeDeleted(includeDeleted bool) {
	arg.IncludeDeleted = includeDeleted
}

func (arg *Shop) GetIdEqual() string {
	return arg.IdEqual
}

func (arg *Shop) SetIdEqual(idEqual string) {
	arg.IdEqual = idEqual
}

func (arg *Shop) GetUpdateObject() interface{} {
	return arg.UpdateObject
}

func (arg *Shop) SetUpdateObject(updateObject interface{}) {
	arg.UpdateObject = updateObject
}

func (arg *Shop) GetTargetPage() int {
	return arg.TargetPage
}

func (arg *Shop) SetTargetPage(targetPage int) {
	arg.TargetPage = targetPage
}

func (arg *Shop) GetPageSize() int {
	return arg.PageSize
}

func (arg *Shop) SetPageSize(pageSize int) {
	arg.PageSize = pageSize
}

func (arg *Shop) getCountConditions() (string, []interface{}) {
	sql := ""
	params := make([]interface{}, 0, 9)
	if arg.NameLike != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " name like ?"
		params = append(params, "%"+arg.NameLike+"%")
	}
	if arg.HideEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " hide = ?"
		params = append(params, arg.HideEqual == "true")
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
