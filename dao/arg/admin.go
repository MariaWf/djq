package arg

import (
	"mimi/djq/model"
)

type Admin struct {
	IdEqual           string
	IncludeDeleted    bool
	NameLike          string
	NameEqual         string
	OrderBy           string
	IdsIn             []string
	LockedOnly        bool
	UnlockedOnly      bool
	KeywordLike       string

	PasswordEqual     string

	PageSize          int `form:"pageSize" json:"pageSize"`
	TargetPage        int `form:"targetPage" json:"targetPage"`

	ShowColumnNames   []string

	UpdateObject      interface{}
	UpdateColumnNames []string
}

func (arg *Admin) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *Admin) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *Admin) GetUpdateColumnNames() []string {
	return arg.UpdateColumnNames
}

func (arg *Admin) SetUpdateColumnNames(updateColumnNames []string) {
	arg.UpdateColumnNames = updateColumnNames
}

func (arg *Admin) GetOrderBy() string {
	return arg.OrderBy
}

func (arg *Admin) SetOrderBy(orderBy string) {
	arg.OrderBy = orderBy
}

func (arg *Admin) GetIncludeDeleted() bool {
	return arg.IncludeDeleted
}

func (arg *Admin) SetIncludeDeleted(includeDeleted bool) {
	arg.IncludeDeleted = includeDeleted
}

func (arg *Admin) GetIdEqual() string {
	return arg.IdEqual
}

func (arg *Admin) SetIdEqual(idEqual string) {
	arg.IdEqual = idEqual
}

func (arg *Admin) GetUpdateObject() interface{} {
	return arg.UpdateObject
}

func (arg *Admin) SetUpdateObject(updateObject interface{}) {
	arg.UpdateObject = updateObject
}

func (arg *Admin) GetTargetPage() int {
	return arg.TargetPage
}

func (arg *Admin) SetTargetPage(targetPage int) {
	arg.TargetPage = targetPage
}

func (arg *Admin) GetPageSize() int {
	return arg.PageSize
}

func (arg *Admin) SetPageSize(pageSize int) {
	arg.PageSize = pageSize
}

func (arg *Admin) getBaseSql(sql string) string {
	return bindTableName(sql, tableNameAdmin)
}

func (arg *Admin) getAllColumnNames() []string {
	return ColumnNamesAdmin
}

func (arg *Admin) getShowColumnNames() []string {
	if arg.ShowColumnNames == nil || len(arg.ShowColumnNames) == 0 {
		return arg.getAllColumnNames()
	}
	s := make([]string, 0, len(arg.ShowColumnNames))
	for _, v := range arg.ShowColumnNames {
		switch v {
		case "id":s = append(s, "id")
		case "name":s = append(s, "name")
		case "mobile":s = append(s, "mobile")
		case "password":s = append(s, "password")
		case "locked":s = append(s, "locked")
		}
	}
	if len(s) == 0 {
		return arg.getAllColumnNames()
	}
	return s
}

func (arg *Admin) getColumnNameValues() ([]string, []interface{}) {
	if arg.UpdateColumnNames == nil || len(arg.UpdateColumnNames) == 0 {
		arg.UpdateColumnNames = arg.getAllColumnNames()[1:]
	}
	s := make([]string, 0, len(arg.UpdateColumnNames))
	params := make([]interface{}, 0, 9)
	for _, v := range arg.UpdateColumnNames {
		switch v {
		case "name":s = append(s, "name = ?")
			params = append(params, arg.UpdateObject.(*model.Admin).Name)
		case "mobile":s = append(s, "mobile = ?")
			params = append(params, arg.UpdateObject.(*model.Admin).Mobile)
		case "password":s = append(s, "password = ?")
			params = append(params, arg.UpdateObject.(*model.Admin).Password)
		case "locked":s = append(s, "locked = ?")
			params = append(params, arg.UpdateObject.(*model.Admin).Locked)
		}
	}
	return s, params
}

func (arg *Admin) getCountConditions() (string, []interface{}) {
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
	if arg.KeywordLike!=""{
		if sql != "" {
			sql += " and"
		}
		sql += " (name like ? or mobile like ?)"
		params = append(params, arg.KeywordLike)
		params = append(params, arg.KeywordLike)
	}
	if arg.PasswordEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " password = ?"
		params = append(params, arg.PasswordEqual)
	}
	if arg.LockedOnly || arg.UnlockedOnly {
		if sql != "" {
			sql += " and"
		}
		sql += " locked = ?"
		params = append(params, arg.LockedOnly)
	}
	if len(params) != 0 {
		sql = " where" + sql
	}
	return sql, params
}



//type Admin struct {
//	NameLike          string
//	NameEqual         string
//	IdEqual           string
//	LockedOnly        bool
//	UnlockedOnly      bool
//
//	PasswordEqual     string
//
//	PageSize          int
//	TargetPage        int
//
//	ShowColumnNames   []string
//
//	UpdateObject      *model.Admin
//	UpdateColumnNames []string
//}
//
//func (arg *Admin) BuildFindSql() (string, []interface{}, []string) {
//	sql := arg.getBaseSql(SelectSql)
//
//	columnNames := arg.getShowColumnNames()
//	sql = bindColumnNames(sql, strings.Join(columnNames, ","))
//
//	conditionStr, params := arg.getFindConditions()
//	sql = bindConditions(sql, conditionStr)
//
//	return sql, params, columnNames
//}
//
//func (arg *Admin) BuildCountSql() (string, []interface{}) {
//	sql := arg.getBaseSql(CountSql)
//
//	conditionStr, params := arg.getCountConditions()
//	sql = bindConditions(sql, conditionStr)
//
//	return sql, params
//}
//
//func (arg *Admin) BuildInsertSql() (string, []string) {
//	sql := arg.getBaseSql(InsertSql)
//
//	columnNames := ColumnNamesAdmin
//	sql = bindColumnNames(sql, strings.Join(columnNames, ","))
//
//	columnValues := arg.getColumnValuePlaceHolder(columnNames)
//	sql = bindColumnValues(sql, strings.Join(columnValues, ","))
//
//	return sql, columnNames
//}
//
//func (arg *Admin) BuildUpdateSql() (string, []interface{}) {
//	if arg.UpdateObject == nil {
//		panic(ErrUpdateObjectEmpty)
//	}
//	sql := arg.getBaseSql(UpdateSql)
//
//	columnNameValues, paramValues := arg.getColumnNameValues()
//	sql = bindColumnNameValues(sql, strings.Join(columnNameValues, ","))
//
//	conditionStr, paramConditions := arg.getCountConditions()
//	sql = bindConditions(sql, conditionStr)
//
//	return sql, append(paramValues, paramConditions...)
//}
//
//func (arg *Admin) BuildDeleteSql() (string, []interface{}) {
//	sql := arg.getBaseSql(DeleteSql)
//
//	conditionStr, params := arg.getCountConditions()
//	sql = bindConditions(sql, conditionStr)
//
//	return sql, params
//}
//
//func (arg *Admin) BuildLogicalDeleteSql() (string, []interface{}) {
//	sql := arg.getBaseSql(UpdateSql)
//
//	sql = bindColumnNameValues(sql, " del_flag = true")
//
//	conditionStr, paramConditions := arg.getCountConditions()
//	sql = bindConditions(sql, conditionStr)
//
//	return sql, paramConditions
//}
//
//func (arg *Admin) getBaseSql(sql string) string {
//	return bindTableName(sql, tableNameAdmin)
//}
//
//func (arg *Admin) getShowColumnNames() []string {
//	if arg.ShowColumnNames == nil || len(arg.ShowColumnNames) == 0 {
//		return ColumnNamesAdmin
//	}
//	s := make([]string, 0, len(arg.ShowColumnNames))
//	for _, v := range arg.ShowColumnNames {
//		switch v {
//		case "id":s = append(s, "id")
//		case "name":s = append(s, "name")
//		case "mobile":s = append(s, "mobile")
//		case "password":s = append(s, "password")
//		case "locked":s = append(s, "locked")
//		}
//	}
//	if len(s) == 0 {
//		return ColumnNamesAdmin
//	}
//	return s
//}
//
//func (arg *Admin) getColumnNameValues() ([]string, []interface{}) {
//	if arg.UpdateColumnNames == nil || len(arg.UpdateColumnNames) == 0 {
//		arg.UpdateColumnNames = ColumnNamesAdmin[1:]
//	}
//	s := make([]string, 0, len(arg.UpdateColumnNames))
//	params := make([]interface{}, 0, 9)
//	for _, v := range arg.UpdateColumnNames {
//		switch v {
//		case "name":s = append(s, "name = ?")
//			params = append(params, arg.UpdateObject.Name)
//		case "mobile":s = append(s, "mobile = ?")
//			params = append(params, arg.UpdateObject.Mobile)
//		case "password":s = append(s, "password = ?")
//			params = append(params, arg.UpdateObject.Password)
//		case "locked":s = append(s, "locked = ?")
//			params = append(params, arg.UpdateObject.Locked)
//		}
//	}
//	return s, params
//}
//
//func (arg *Admin) getColumnValuePlaceHolder(columnNames []string) []string {
//	len := len(columnNames)
//	placeholder := make([]string, 0, len)
//	for i := 0; i < len; i++ {
//		placeholder = append(placeholder, "?")
//	}
//	return placeholder
//}
//
//func (arg *Admin) getFindConditions() (string, []interface{}) {
//	sql, params := arg.getCountConditions()
//	if arg.PageSize > 0 {
//		sql += " limit ?,?"
//		start := util.ComputePageStart(arg.TargetPage, arg.PageSize)
//		params = append(params, start, arg.PageSize)
//	}
//	return sql, params
//}
//
//func (arg *Admin) getCountConditions() (string, []interface{}) {
//	sql := ""
//	params := make([]interface{}, 0, 9)
//	if arg.NameLike != "" {
//		if sql != "" {
//			sql += " and"
//		}
//		sql += " name like ?"
//		params = append(params, "%" + arg.NameLike + "%")
//	}
//	if arg.NameEqual != "" {
//		if sql != "" {
//			sql += " and"
//		}
//		sql += " name = ?"
//		params = append(params, arg.NameEqual)
//	}
//	if arg.PasswordEqual != "" {
//		if sql != "" {
//			sql += " and"
//		}
//		sql += " password = ?"
//		params = append(params, arg.PasswordEqual)
//	}
//	if arg.LockedOnly || arg.UnlockedOnly {
//		if sql != "" {
//			sql += " and"
//		}
//		sql += " locked = ?"
//		params = append(params, arg.LockedOnly)
//	}
//	if len(params) != 0 {
//		sql = " where" + sql
//	}
//	return sql, params
//}