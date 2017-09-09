package arg

import (
	"errors"
	"mimi/djq/util"
	"strings"
)

var ErrUpdateObjectEmpty = errors.New("dao: updateObject is empty")

const (
	//SELECT 列名称 FROM 表名称
	SelectSql = "select {columnNames} from {tableName} {conditions};"
	CountSql = "select count(*) from {tableName} {conditions};"
	//DELETE FROM 表名称 WHERE 列名称 = 值
	DeleteSql = "delete from {tableName} {conditions};"
	//INSERT INTO table_name (列1, 列2,...) VALUES (值1, 值2,....)
	InsertSql = "insert into {tableName} ({columnNames}) values ({columnValues});"
	//"UPDATE 表名称 SET 列名称 = 新值 WHERE 列名称 = 某值"
	UpdateSql = "update {tableName} set {collumnNameValues} {conditions};"
)

const (
	tableNameAdmin = "tbl_admin"
	tableNameRole = "tbl_role"
	tableNameUser = "tbl_user"
	tableNameAdvertisement = "tbl_advertisement"
)

var (
	ColumnNamesAdmin = []string{"id", "name", "mobile", "password", "locked"}
	ColumnNamesRole = []string{"id", "name", "description", "permission_list_str"}
	ColumnNamesUser = []string{"id", "name", "nick_name", "mobile", "password"}
	ColumnNamesAdvertisement = []string{"id", "name", "image", "link", "priority","hide","description"}
)

func bindColumnNames(sql, columnNames string) string {
	return strings.Replace(sql, "{columnNames}", columnNames, -1)
}

func bindColumnValues(sql, columnValues string) string {
	return strings.Replace(sql, "{columnValues}", columnValues, -1)
}

func bindColumnNameValues(sql, collumnNameValues string) string {
	return strings.Replace(sql, "{collumnNameValues}", collumnNameValues, -1)
}

func bindTableName(sql, tableName string) string {
	return strings.Replace(sql, "{tableName}", tableName, -1)
}

func bindConditions(sql, conditions string) string {
	return strings.Replace(sql, "{conditions}", conditions, -1)
}

type BaseArgInterface interface {
	getBaseSql(sql string) string
	getAllColumnNames() []string
	getShowColumnNames() []string
	getColumnNameValues() ([]string, []interface{})
	getCountConditions() (string, []interface{})
	GetTargetPage() int
	SetTargetPage(int)
	GetPageSize() int
	SetPageSize(int)
	GetIdEqual() string
	SetIdEqual(idEqual string)
	GetIdsIn() []string
	SetIdsIn(idsIn []string)
	GetIncludeDeleted() bool
	SetIncludeDeleted(includeDeleted bool)
	GetUpdateObject() interface{}
	SetUpdateObject(updateObject interface{})
	GetUpdateColumnNames() []string
	SetUpdateColumnNames(updateColumnNames []string)
	GetOrderBy() string
}

func BuildFindSql(arg BaseArgInterface) (string, []interface{}, []string) {
	sql := arg.getBaseSql(SelectSql)

	columnNames := arg.getShowColumnNames()
	sql = bindColumnNames(sql, strings.Join(columnNames, ","))

	conditionStr, params := getFindConditions(arg)
	sql = bindConditions(sql, conditionStr)

	return sql, params, columnNames
}

func BuildCountSql(arg BaseArgInterface) (string, []interface{}) {
	sql := arg.getBaseSql(CountSql)

	conditionStr, params := getCountConditions(arg)
	sql = bindConditions(sql, conditionStr)

	return sql, params
}

func BuildInsertSql(arg BaseArgInterface) (string, []string) {
	sql := arg.getBaseSql(InsertSql)

	columnNames := arg.getAllColumnNames()
	sql = bindColumnNames(sql, strings.Join(columnNames, ","))

	columnValues := getColumnValuePlaceHolder(columnNames)
	sql = bindColumnValues(sql, strings.Join(columnValues, ","))

	return sql, columnNames
}

func BuildUpdateSql(arg BaseArgInterface) (string, []interface{}) {
	if arg.GetUpdateObject() == nil {
		panic(ErrUpdateObjectEmpty)
	}
	sql := arg.getBaseSql(UpdateSql)

	columnNameValues, paramValues := arg.getColumnNameValues()
	sql = bindColumnNameValues(sql, strings.Join(columnNameValues, ","))

	conditionStr, paramConditions := getCountConditions(arg)
	sql = bindConditions(sql, conditionStr)

	return sql, append(paramValues, paramConditions...)
}

func BuildDeleteSql(arg BaseArgInterface) (string, []interface{}) {
	sql := arg.getBaseSql(DeleteSql)

	conditionStr, params := getCountConditions(arg)
	sql = bindConditions(sql, conditionStr)

	return sql, params
}

func BuildLogicalDeleteSql(arg BaseArgInterface) (string, []interface{}) {
	sql := arg.getBaseSql(UpdateSql)

	sql = bindColumnNameValues(sql, " del_flag = true")

	conditionStr, paramConditions := getCountConditions(arg)
	sql = bindConditions(sql, conditionStr)

	return sql, paramConditions
}

func getColumnValuePlaceHolder(columnNames []string) []string {
	len := len(columnNames)
	placeholder := make([]string, 0, len)
	for i := 0; i < len; i++ {
		placeholder = append(placeholder, "?")
	}
	return placeholder
}

func getFindConditions(arg BaseArgInterface) (string, []interface{}) {
	sql, params := getCountConditions(arg)
	if arg.GetPageSize() > 0 {
		sql += " limit ?,?"
		start := util.ComputePageStart(arg.GetTargetPage(), arg.GetPageSize())
		params = append(params, start, arg.GetPageSize())
	}
	if arg.GetOrderBy() != "" {
		sql += " order by " + arg.GetOrderBy()
	}
	return sql, params
}

func getCountConditions(arg BaseArgInterface) (string, []interface{}) {
	sql, params := arg.getCountConditions()
	if arg.GetIdEqual() != "" {
		if sql != "" {
			sql += " and"
		} else {
			sql += " where"
		}
		sql += " id = ?"
		params = append(params, arg.GetIdEqual())
	}
	if arg.GetIdsIn() != nil && len(arg.GetIdsIn()) != 0 {
		if sql != "" {
			sql += " and"
		} else {
			sql += " where"
		}
		sql += " id in ("
		for i, id := range arg.GetIdsIn() {
			if i != 0 {
				sql += ","
			}
			sql += "?"
			params = append(params, id)
		}
		sql += ")"
	}
	if !arg.GetIncludeDeleted() {
		if sql != "" {
			sql += " and"
		} else {
			sql += " where"
		}
		sql += " del_flag = false"
	}
	return sql, params
}
