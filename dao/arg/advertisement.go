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

	PageSize   int
	TargetPage int

	ShowColumnNames []string

	UpdateObject      interface{}
	UpdateColumnNames []string
}

func (arg *Advertisement) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *Advertisement) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *Advertisement) GetUpdateColumnNames() []string {
	return arg.UpdateColumnNames
}

func (arg *Advertisement) SetUpdateColumnNames(updateColumnNames []string) {
	arg.UpdateColumnNames = updateColumnNames
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

func (arg *Advertisement) getBaseSql(sql string) string {
	return bindTableName(sql, tableNameAdvertisement)
}

func (arg *Advertisement) getAllColumnNames() []string {
	return ColumnNamesAdvertisement
}

func (arg *Advertisement) getShowColumnNames() []string {
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
		case "image":
			s = append(s, "image")
		case "link":
			s = append(s, "link")
		case "priority":
			s = append(s, "priority")
		case "hide":
			s = append(s, "hide")
		case "description":
			s = append(s, "description")
		}
	}
	if len(s) == 0 {
		return arg.getAllColumnNames()
	}
	return s
}

func (arg *Advertisement) getColumnNameValues() ([]string, []interface{}) {
	if arg.UpdateColumnNames == nil || len(arg.UpdateColumnNames) == 0 {
		arg.UpdateColumnNames = arg.getAllColumnNames()[1:]
	}
	s := make([]string, 0, len(arg.UpdateColumnNames))
	params := make([]interface{}, 0, 9)
	for _, v := range arg.UpdateColumnNames {
		ColumnNamesAdvertisement = []string{"id", "name", "image", "link", "priority","hide","description"}
		switch v {
		case "name":
			s = append(s, "name = ?")
			params = append(params, arg.UpdateObject.(*model.Advertisement).Name)
		case "image":
			s = append(s, "image = ?")
			params = append(params, arg.UpdateObject.(*model.Advertisement).Image)
		case "link":
			s = append(s, "link = ?")
			params = append(params, arg.UpdateObject.(*model.Advertisement).Link)
		case "priority":
			s = append(s, "priority = ?")
			params = append(params, arg.UpdateObject.(*model.Advertisement).Priority)
		case "hide":
			s = append(s, "hide = ?")
			params = append(params, arg.UpdateObject.(*model.Advertisement).Hide)
		case "description":
			s = append(s, "description = ?")
			params = append(params, arg.UpdateObject.(*model.Advertisement).Description)
		}
	}
	return s, params
}

func (arg *Advertisement) getCountConditions() (string, []interface{}) {
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
