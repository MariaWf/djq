package arg

import (
	"mimi/djq/model"
)

type ShopIntroductionImage struct {
	IdEqual        string
	IncludeDeleted bool
	OrderBy        string
	IdsIn          []string

	ShopIdEqual    string `form:"shopId" json:"shopId"`
	NotIncludeHide bool
	HideEqual      string `form:"hide" json:"hide"`

	PageSize       int `form:"pageSize" json:"pageSize"`
	TargetPage     int `form:"targetPage" json:"targetPage"`

	DisplayNames   []string

	UpdateObject   interface{}
	UpdateNames    []string
}

func (arg *ShopIntroductionImage) GetDisplayNames() []string {
	return arg.DisplayNames
}

func (arg *ShopIntroductionImage) GetModelInstance() model.BaseModelInterface {
	return &model.ShopIntroductionImage{}
}

func (arg *ShopIntroductionImage) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *ShopIntroductionImage) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *ShopIntroductionImage) GetUpdateNames() []string {
	return arg.UpdateNames
}

func (arg *ShopIntroductionImage) SetUpdateNames(updateNames []string) {
	arg.UpdateNames = updateNames
}

func (arg *ShopIntroductionImage) GetOrderBy() string {
	return arg.OrderBy
}

func (arg *ShopIntroductionImage) SetOrderBy(orderBy string) {
	arg.OrderBy = orderBy
}

func (arg *ShopIntroductionImage) GetIncludeDeleted() bool {
	return arg.IncludeDeleted
}

func (arg *ShopIntroductionImage) SetIncludeDeleted(includeDeleted bool) {
	arg.IncludeDeleted = includeDeleted
}

func (arg *ShopIntroductionImage) GetIdEqual() string {
	return arg.IdEqual
}

func (arg *ShopIntroductionImage) SetIdEqual(idEqual string) {
	arg.IdEqual = idEqual
}

func (arg *ShopIntroductionImage) GetUpdateObject() interface{} {
	return arg.UpdateObject
}

func (arg *ShopIntroductionImage) SetUpdateObject(updateObject interface{}) {
	arg.UpdateObject = updateObject
}

func (arg *ShopIntroductionImage) GetTargetPage() int {
	return arg.TargetPage
}

func (arg *ShopIntroductionImage) SetTargetPage(targetPage int) {
	arg.TargetPage = targetPage
}

func (arg *ShopIntroductionImage) GetPageSize() int {
	return arg.PageSize
}

func (arg *ShopIntroductionImage) SetPageSize(pageSize int) {
	arg.PageSize = pageSize
}

func (arg *ShopIntroductionImage) getCountConditions() (string, []interface{}) {
	sql := ""
	params := make([]interface{}, 0, 9)
	if arg.ShopIdEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " shop_id = ?"
		params = append(params, arg.ShopIdEqual)
	}
	if arg.NotIncludeHide {
		if sql != "" {
			sql += " and"
		}
		sql += " hide = ?"
		params = append(params, false)
	}
	if arg.HideEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " hide = ?"
		params = append(params, arg.HideEqual == "true")
	}
	if len(params) != 0 {
		sql = " where" + sql
	}
	return sql, params
}
