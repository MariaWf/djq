package arg

import (
	"mimi/djq/model"
)

type PromotionalPartner struct {
	IdEqual        string
	IncludeDeleted bool
	NameLike       string `form:"keyword" json:"keyword"`
	NameEqual      string
	OrderBy        string
	IdsIn          []string

	PageSize   int `form:"pageSize" json:"pageSize"`
	TargetPage int `form:"targetPage" json:"targetPage"`

	DisplayNames []string

	UpdateObject interface{}
	UpdateNames  []string
}

func (arg *PromotionalPartner) GetDisplayNames() []string {
	return arg.DisplayNames
}

func (arg *PromotionalPartner) GetModelInstance() model.BaseModelInterface {
	return &model.PromotionalPartner{}
}

func (arg *PromotionalPartner) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *PromotionalPartner) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *PromotionalPartner) GetUpdateNames() []string {
	return arg.UpdateNames
}

func (arg *PromotionalPartner) SetUpdateNames(updateNames []string) {
	arg.UpdateNames = updateNames
}

func (arg *PromotionalPartner) GetOrderBy() string {
	return arg.OrderBy
}

func (arg *PromotionalPartner) SetOrderBy(orderBy string) {
	arg.OrderBy = orderBy
}

func (arg *PromotionalPartner) GetIncludeDeleted() bool {
	return arg.IncludeDeleted
}

func (arg *PromotionalPartner) SetIncludeDeleted(includeDeleted bool) {
	arg.IncludeDeleted = includeDeleted
}

func (arg *PromotionalPartner) GetIdEqual() string {
	return arg.IdEqual
}

func (arg *PromotionalPartner) SetIdEqual(idEqual string) {
	arg.IdEqual = idEqual
}

func (arg *PromotionalPartner) GetUpdateObject() interface{} {
	return arg.UpdateObject
}

func (arg *PromotionalPartner) SetUpdateObject(updateObject interface{}) {
	arg.UpdateObject = updateObject
}

func (arg *PromotionalPartner) GetTargetPage() int {
	return arg.TargetPage
}

func (arg *PromotionalPartner) SetTargetPage(targetPage int) {
	arg.TargetPage = targetPage
}

func (arg *PromotionalPartner) GetPageSize() int {
	return arg.PageSize
}

func (arg *PromotionalPartner) SetPageSize(pageSize int) {
	arg.PageSize = pageSize
}

func (arg *PromotionalPartner) getCountConditions() (string, []interface{}) {
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
	if len(params) != 0 {
		sql = " where" + sql
	}
	return sql, params
}
