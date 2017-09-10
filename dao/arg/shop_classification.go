package arg

import (
	"mimi/djq/model"
)

type ShopClassification struct {
	IdEqual           string
	IncludeDeleted    bool
	NameLike          string
	NameEqual         string
	OrderBy           string
	IdsIn             []string

	NotIncludeHide    bool

	PageSize          int
	TargetPage        int

	DisplayNames   []string

	UpdateObject      interface{}
	UpdateNames []string
}

func (arg *ShopClassification) GetDisplayNames() []string {
	return arg.DisplayNames
}

func (arg *ShopClassification) GetModelInstance() model.BaseModelInterface {
	return &model.ShopClassification{}
}

func (arg *ShopClassification) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *ShopClassification) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *ShopClassification) GetUpdateNames() []string {
	return arg.UpdateNames
}

func (arg *ShopClassification) SetUpdateNames(updateNames []string) {
	arg.UpdateNames = updateNames
}

func (arg *ShopClassification) GetOrderBy() string {
	return arg.OrderBy
}

func (arg *ShopClassification) SetOrderBy(orderBy string) {
	arg.OrderBy = orderBy
}

func (arg *ShopClassification) GetIncludeDeleted() bool {
	return arg.IncludeDeleted
}

func (arg *ShopClassification) SetIncludeDeleted(includeDeleted bool) {
	arg.IncludeDeleted = includeDeleted
}

func (arg *ShopClassification) GetIdEqual() string {
	return arg.IdEqual
}

func (arg *ShopClassification) SetIdEqual(idEqual string) {
	arg.IdEqual = idEqual
}

func (arg *ShopClassification) GetUpdateObject() interface{} {
	return arg.UpdateObject
}

func (arg *ShopClassification) SetUpdateObject(updateObject interface{}) {
	arg.UpdateObject = updateObject
}

func (arg *ShopClassification) GetTargetPage() int {
	return arg.TargetPage
}

func (arg *ShopClassification) SetTargetPage(targetPage int) {
	arg.TargetPage = targetPage
}

func (arg *ShopClassification) GetPageSize() int {
	return arg.PageSize
}

func (arg *ShopClassification) SetPageSize(pageSize int) {
	arg.PageSize = pageSize
}

func (arg *ShopClassification) getCountConditions() (string, []interface{}) {
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
