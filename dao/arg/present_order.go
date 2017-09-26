package arg

import (
	"mimi/djq/model"
)

type PresentOrder struct {
	IdEqual        string
	IncludeDeleted bool
	NameLike       string
	NameEqual      string
	OrderBy        string
	IdsIn          []string

	UserIdEqual string
	NumberEqual string `form:"keyword" json:"keyword"`
	StatusEqual string `form:"status" json:"status"`

	PageSize   int `form:"pageSize" json:"pageSize"`
	TargetPage int `form:"targetPage" json:"targetPage"`

	DisplayNames []string

	UpdateObject interface{}
	UpdateNames  []string
}

func (arg *PresentOrder) GetDisplayNames() []string {
	return arg.DisplayNames
}

func (arg *PresentOrder) GetModelInstance() model.BaseModelInterface {
	return &model.PresentOrder{}
}

func (arg *PresentOrder) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *PresentOrder) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *PresentOrder) GetUpdateNames() []string {
	return arg.UpdateNames
}

func (arg *PresentOrder) SetUpdateNames(updateNames []string) {
	arg.UpdateNames = updateNames
}

func (arg *PresentOrder) GetOrderBy() string {
	return arg.OrderBy
}

func (arg *PresentOrder) SetOrderBy(orderBy string) {
	arg.OrderBy = orderBy
}

func (arg *PresentOrder) GetIncludeDeleted() bool {
	return arg.IncludeDeleted
}

func (arg *PresentOrder) SetIncludeDeleted(includeDeleted bool) {
	arg.IncludeDeleted = includeDeleted
}

func (arg *PresentOrder) GetIdEqual() string {
	return arg.IdEqual
}

func (arg *PresentOrder) SetIdEqual(idEqual string) {
	arg.IdEqual = idEqual
}

func (arg *PresentOrder) GetUpdateObject() interface{} {
	return arg.UpdateObject
}

func (arg *PresentOrder) SetUpdateObject(updateObject interface{}) {
	arg.UpdateObject = updateObject
}

func (arg *PresentOrder) GetTargetPage() int {
	return arg.TargetPage
}

func (arg *PresentOrder) SetTargetPage(targetPage int) {
	arg.TargetPage = targetPage
}

func (arg *PresentOrder) GetPageSize() int {
	return arg.PageSize
}

func (arg *PresentOrder) SetPageSize(pageSize int) {
	arg.PageSize = pageSize
}

func (arg *PresentOrder) getCountConditions() (string, []interface{}) {
	sql := ""
	params := make([]interface{}, 0, 9)
	if arg.NameLike != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " name like ?"
		params = append(params, "%"+arg.NameLike+"%")
	}
	if arg.NumberEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " number = ?"
		params = append(params, arg.NumberEqual)
	}
	if arg.NameEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " name = ?"
		params = append(params, arg.NameEqual)
	}
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
