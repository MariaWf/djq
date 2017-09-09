package arg

import (
	"mimi/djq/model"
)

type Role struct {
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

func (arg *Role) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *Role) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *Role) GetUpdateColumnNames() []string {
	return arg.UpdateColumnNames
}

func (arg *Role) SetUpdateColumnNames(updateColumnNames []string) {
	arg.UpdateColumnNames = updateColumnNames
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

func (arg *Role) getBaseSql(sql string) string {
	return bindTableName(sql, tableNameRole)
}

func (arg *Role) getAllColumnNames() []string {
	return ColumnNamesRole
}

func (arg *Role) getShowColumnNames() []string {
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
		case "permissionListStr":
			s = append(s, "permission_list_str")
		}
	}
	if len(s) == 0 {
		return arg.getAllColumnNames()
	}
	return s
}

func (arg *Role) getColumnNameValues() ([]string, []interface{}) {
	if arg.UpdateColumnNames == nil || len(arg.UpdateColumnNames) == 0 {
		arg.UpdateColumnNames = arg.getAllColumnNames()[1:]
	}
	s := make([]string, 0, len(arg.UpdateColumnNames))
	params := make([]interface{}, 0, 9)
	for _, v := range arg.UpdateColumnNames {
		switch v {
		case "name":
			s = append(s, "name = ?")
			params = append(params, arg.UpdateObject.(*model.Role).Name)
		case "description":
			s = append(s, "description = ?")
			params = append(params, arg.UpdateObject.(*model.Role).Description)
		case "permissionListStr":
			s = append(s, "permission_list_str = ?")
			params = append(params, arg.UpdateObject.(*model.Role).PermissionListStr)
		}
	}
	return s, params
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
	if len(params) != 0 {
		sql = " where" + sql
	}
	return sql, params
}
