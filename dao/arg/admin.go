package arg

import (
	"mimi/djq/model"
)

type Admin struct {
	IdEqual        string
	IncludeDeleted bool
	NameLike       string
	NameEqual      string
	OrderBy        string
	IdsIn          []string
	LockedOnly     bool
	UnlockedOnly   bool
	KeywordLike    string

	PasswordEqual string

	PageSize   int `form:"pageSize" json:"pageSize"`
	TargetPage int `form:"targetPage" json:"targetPage"`

	ShowColumnNames []string

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
		case "id":
			s = append(s, "id")
		case "name":
			s = append(s, "name")
		case "mobile":
			s = append(s, "mobile")
		case "password":
			s = append(s, "password")
		case "locked":
			s = append(s, "locked")
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
		case "name":
			s = append(s, "name = ?")
			params = append(params, arg.UpdateObject.(*model.Admin).Name)
		case "mobile":
			s = append(s, "mobile = ?")
			params = append(params, arg.UpdateObject.(*model.Admin).Mobile)
		case "password":
			s = append(s, "password = ?")
			params = append(params, arg.UpdateObject.(*model.Admin).Password)
		case "locked":
			s = append(s, "locked = ?")
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
		params = append(params, "%"+arg.NameLike+"%")
	}
	if arg.NameEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " name = ?"
		params = append(params, arg.NameEqual)
	}
	if arg.KeywordLike != "" {
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
