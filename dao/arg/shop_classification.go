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

	ShowColumnNames   []string

	UpdateObject      interface{}
	UpdateColumnNames []string
}

func (arg *ShopClassification) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *ShopClassification) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *ShopClassification) GetUpdateColumnNames() []string {
	return arg.UpdateColumnNames
}

func (arg *ShopClassification) SetUpdateColumnNames(updateColumnNames []string) {
	arg.UpdateColumnNames = updateColumnNames
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

func (arg *ShopClassification) getBaseSql(sql string) string {
	return bindTableName(sql, tableNameShopClassification)
}

func (arg *ShopClassification) getAllColumnNames() []string {
	return ColumnNamesShopClassification
}

func (arg *ShopClassification) getShowColumnNames() []string {
	if arg.ShowColumnNames == nil || len(arg.ShowColumnNames) == 0 {
		return arg.getAllColumnNames()
	}
	s := make([]string, 0, len(arg.ShowColumnNames))
	for _, v := range arg.ShowColumnNames {
		switch v {
		case "id":
			s = append(s, "id")
		case "name":
			s = append(s, "name")
		case "description":
			s = append(s, "description")
		case "hide":
			s = append(s, "hide")
		case "priority":
			s = append(s, "priority")
		}
	}
	if len(s) == 0 {
		return arg.getAllColumnNames()
	}
	return s
}

func (arg *ShopClassification) getColumnNameValues() ([]string, []interface{}) {
	if arg.UpdateColumnNames == nil || len(arg.UpdateColumnNames) == 0 {
		arg.UpdateColumnNames = arg.getAllColumnNames()[1:]
	}
	s := make([]string, 0, len(arg.UpdateColumnNames))
	params := make([]interface{}, 0, 9)
	for _, v := range arg.UpdateColumnNames {
		//s = append(s,v+" = ?")
		//params = append(params,arg.UpdateObject.(*model.ShopClassification).GetValue4Map(v))
		switch v {
		case "name":
			s = append(s, "name = ?")
			params = append(params, arg.UpdateObject.(*model.ShopClassification).Name)
		case "description":
			s = append(s, "description = ?")
			params = append(params, arg.UpdateObject.(*model.ShopClassification).Description)
		case "hide":
			s = append(s, "hide")
			params = append(params, arg.UpdateObject.(*model.ShopClassification).Hide)
		case "priority":
			s = append(s, "priority")
			params = append(params, arg.UpdateObject.(*model.ShopClassification).Priority)

		}
	}
	return s, params
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
