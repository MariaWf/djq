package arg

import (
	"mimi/djq/model"
	"mimi/djq/util"
	"strings"
)

type User struct {
	NameLike      string
	NameEqual     string
	MobileLike    string
	MobileEqual   string
	NickNameLike  string
	NickNameEqual string
	IdEqual       string

	LoginNameLike  string
	LoginNameEqual string
	PasswordEqual  string

	PageSize   int
	TargetPage int

	ShowColumnNames []string

	UpdateObject      *model.User
	UpdateColumnNames []string
}

func (u *User) BuildFindSql() (string, []interface{}, []string) {
	sql := u.getBaseSql(SelectSql)

	columnNames := u.getShowColumnNames()
	sql = bindColumnNames(sql, strings.Join(columnNames, ","))

	conditionStr, params := u.getFindConditions()
	sql = bindConditions(sql, conditionStr)

	return sql, params, columnNames
}

func (u *User) BuildCountSql() (string, []interface{}) {
	sql := u.getBaseSql(CountSql)

	conditionStr, params := u.getCountConditions()
	sql = bindConditions(sql, conditionStr)

	return sql, params
}

func (u *User) BuildInsertSql() (string, []string) {
	sql := u.getBaseSql(InsertSql)

	columnNames := ColumnNamesUser
	sql = bindColumnNames(sql, strings.Join(columnNames, ","))

	columnValues := u.getColumnValuePlaceHolder(columnNames)
	sql = bindColumnValues(sql, strings.Join(columnValues, ","))

	return sql, columnNames
}

func (u *User) BuildUpdateSql() (string, []interface{}) {
	if u.UpdateObject == nil {
		panic(ErrUpdateObjectEmpty)
	}
	sql := u.getBaseSql(UpdateSql)

	columnNameValues, paramValues := u.getColumnNameValues()
	sql = bindColumnNameValues(sql, strings.Join(columnNameValues, ","))

	conditionStr, paramConditions := u.getCountConditions()
	sql = bindConditions(sql, conditionStr)

	return sql, append(paramValues, paramConditions...)
}

func (u *User) BuildDeleteSql() (string, []interface{}) {
	sql := u.getBaseSql(DeleteSql)

	conditionStr, params := u.getCountConditions()
	sql = bindConditions(sql, conditionStr)

	return sql, params
}

func (u *User) getBaseSql(sql string) string {
	return bindTableName(sql, tableNameUser)
}

func (u *User) getShowColumnNames() []string {
	if u.ShowColumnNames == nil || len(u.ShowColumnNames) == 0 {
		return []string{"*"}
	}
	s := make([]string, 0, len(u.ShowColumnNames))
	for _, v := range u.ShowColumnNames {
		switch v {
		case "id":
			s = append(s, "id")
		case "name":
			s = append(s, "name")
		case "nickName":
			s = append(s, "nick_name")
		case "mobile":
			s = append(s, "mobile")
		case "password":
			s = append(s, "password")
		}
	}
	if len(s) == 0 {
		s = append(s, "*")
	}
	return s
}

func (u *User) getColumnNameValues() ([]string, []interface{}) {
	if u.UpdateColumnNames == nil || len(u.UpdateColumnNames) == 0 {
		u.UpdateColumnNames = ColumnNamesUser[1:]
	}
	s := make([]string, 0, len(u.UpdateColumnNames))
	params := make([]interface{}, 0, 9)
	for _, v := range u.UpdateColumnNames {
		switch v {
		case "name":
			s = append(s, "name = ?")
			params = append(params, u.UpdateObject.Name)
		case "nickName":
			s = append(s, "nick_name = ?")
			params = append(params, u.UpdateObject.NickName)
		case "mobile":
			s = append(s, "mobile = ?")
			params = append(params, u.UpdateObject.Mobile)
		case "password":
			s = append(s, "password = ?")
			params = append(params, u.UpdateObject.Password)
		}
	}
	return s, params
}

func (u *User) getColumnValuePlaceHolder(columnNames []string) []string {
	len := len(columnNames)
	placeholder := make([]string, 0, len)
	for i := 0; i < len; i++ {
		placeholder = append(placeholder, "?")
	}
	return placeholder
}

func (u *User) getFindConditions() (string, []interface{}) {
	sql, params := u.getCountConditions()
	if u.PageSize > 0 {
		sql += " limit ?,?"
		start := util.ComputePageStart(u.TargetPage, u.PageSize)
		params = append(params, start, u.PageSize)
	}
	return sql, params
}

func (u *User) getCountConditions() (string, []interface{}) {
	sql := ""
	params := make([]interface{}, 0, 9)
	if u.NameLike != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " name like ?"
		params = append(params, "%"+u.NameLike+"%")
	}
	if u.NameEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " name = ?"
		params = append(params, u.NameEqual)
	}
	if u.NickNameLike != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " nick_name like ?"
		params = append(params, "%"+u.NickNameLike+"%")
	}
	if u.NickNameEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " nick_name = ?"
		params = append(params, u.NickNameEqual)
	}
	if u.MobileLike != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " mobile like ?"
		params = append(params, "%"+u.MobileLike+"%")
	}
	if u.MobileEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " mobile = ?"
		params = append(params, u.MobileEqual)
	}
	if u.LoginNameEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " (nick_name = ? or mobile = ?)"
		params = append(params, u.LoginNameEqual, u.LoginNameEqual)
	}
	if u.PasswordEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " password = ?"
		params = append(params, u.PasswordEqual)
	}
	if len(params) != 0 {
		sql = " where" + sql
	}
	return sql, params
}
