package arg

import (
	"mimi/djq/model"
)

type RefundReason struct {
	IdEqual        string
	IncludeDeleted bool
	OrderBy        string
	IdsIn          []string

	DescriptionLike string `form:"keyword" json:"keyword"`

	NotIncludeHide bool

	PageSize   int `form:"pageSize" json:"pageSize"`
	TargetPage int `form:"targetPage" json:"targetPage"`

	DisplayNames []string

	UpdateObject interface{}
	UpdateNames  []string
}

func (arg *RefundReason) GetDisplayNames() []string {
	return arg.DisplayNames
}

func (arg *RefundReason) GetModelInstance() model.BaseModelInterface {
	return &model.RefundReason{}
}

func (arg *RefundReason) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *RefundReason) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *RefundReason) GetUpdateNames() []string {
	return arg.UpdateNames
}

func (arg *RefundReason) SetUpdateNames(updateNames []string) {
	arg.UpdateNames = updateNames
}

func (arg *RefundReason) GetOrderBy() string {
	return arg.OrderBy
}

func (arg *RefundReason) SetOrderBy(orderBy string) {
	arg.OrderBy = orderBy
}

func (arg *RefundReason) GetIncludeDeleted() bool {
	return arg.IncludeDeleted
}

func (arg *RefundReason) SetIncludeDeleted(includeDeleted bool) {
	arg.IncludeDeleted = includeDeleted
}

func (arg *RefundReason) GetIdEqual() string {
	return arg.IdEqual
}

func (arg *RefundReason) SetIdEqual(idEqual string) {
	arg.IdEqual = idEqual
}

func (arg *RefundReason) GetUpdateObject() interface{} {
	return arg.UpdateObject
}

func (arg *RefundReason) SetUpdateObject(updateObject interface{}) {
	arg.UpdateObject = updateObject
}

func (arg *RefundReason) GetTargetPage() int {
	return arg.TargetPage
}

func (arg *RefundReason) SetTargetPage(targetPage int) {
	arg.TargetPage = targetPage
}

func (arg *RefundReason) GetPageSize() int {
	return arg.PageSize
}

func (arg *RefundReason) SetPageSize(pageSize int) {
	arg.PageSize = pageSize
}

func (arg *RefundReason) getCountConditions() (string, []interface{}) {
	sql := ""
	params := make([]interface{}, 0, 9)
	if arg.DescriptionLike != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " description like ?"
		params = append(params, "%"+arg.DescriptionLike+"%")
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
