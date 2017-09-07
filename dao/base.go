package dao

import (
	"mimi/djq/util"
	"database/sql"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"github.com/pkg/errors"
)

var ErrIdEmpty = errors.New("dao: id is empty")

var ErrObjectNotFound = errors.New("dao: object not found")

var ErrObjectEmpty = errors.New("dao: object is empty")

var ErrUnknown = errors.New("dao: unknown")

func BuildId() string {
	return util.BuildUUID()
}

type BaseDaoInterface interface {
	GetArgInstance() arg.BaseArgInterface
	GetModelInstance() model.BaseModelInterface
	GetConn() *sql.Tx
}

func Find(dao BaseDaoInterface,argObj arg.BaseArgInterface) ([]interface{}, error) {
	objList := make([]interface{}, 0, argObj.GetPageSize())
	sql, params, columnNames := arg.BuildFindSql(argObj)
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return nil, errors.Wrap(err, "conn:"+sql)
	}
	defer stmt.Close()
	rows, err := stmt.Query(params...)
	if err != nil {
		return nil, errors.Wrap(err, "stmt:"+sql)
	}
	defer rows.Close()
	for rows.Next() {
		obj := dao.GetModelInstance()
		err = rows.Scan(model.GetPointers4DB(columnNames, obj)...)
		if err != nil {
			return nil, errors.Wrap(err, "rows:"+sql)
		}
		objList = append(objList, obj)
	}
	return objList, nil
}

func Count(dao BaseDaoInterface,argObj arg.BaseArgInterface) (int, error) {
	sql, params := arg.BuildCountSql(argObj)
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return 0, errors.Wrap(err, "conn:"+sql)
	}
	defer stmt.Close()
	rows, err := stmt.Query(params...)
	if err != nil {
		return 0, errors.Wrap(err, "stmt:"+sql)
	}
	defer rows.Close()
	for rows.Next() {
		total := 0
		err = rows.Scan(&total)
		if err != nil {
			return 0, errors.Wrap(err, "rows:"+sql)
		}
		return total, nil
	}
	return 0, errors.Wrap(ErrUnknown, "dao:model:count:"+sql)
}

func Get(dao BaseDaoInterface,id string) (interface{}, error) {
	if id == "" {
		return nil, ErrIdEmpty
	}
	argObj := dao.GetArgInstance()
	argObj.SetIdEqual(id)
	argObj.SetPageSize(1)
	argObj.SetTargetPage(util.BeginPage)
	sql, params, columnNames := arg.BuildFindSql(argObj)
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return nil, errors.Wrap(err, "conn:"+sql)
	}
	defer stmt.Close()
	rows, err := stmt.Query(params...)
	if err != nil {
		return nil, errors.Wrap(err, "stmt:"+sql)
	}
	defer rows.Close()
	for rows.Next() {
		obj := dao.GetModelInstance()
		err = rows.Scan(model.GetPointers4DB(columnNames, obj)...)
		if err != nil {
			return nil, errors.Wrap(err, "rows:"+sql)
		}
		return obj, nil
	}
	return nil, ErrObjectNotFound
}

func Add(dao BaseDaoInterface,obj  model.BaseModelInterface) (interface{}, error) {
	if obj == nil {
		return nil, ErrObjectEmpty
	}
	argObj := dao.GetArgInstance()
	sql, columnNames := arg.BuildInsertSql(argObj)
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return nil, errors.Wrap(err, "conn:"+sql)
	}
	defer stmt.Close()
	id := BuildId()
	params := model.GetValues4DB(columnNames, obj)[:]
	params[0] = id
	_, err = stmt.Exec(params...)
	if err != nil {
		return nil, errors.Wrap(err, "stmt:"+sql)
	}
	obj.SetId(id)
	return obj, nil

}

func Update(dao BaseDaoInterface,obj model.BaseModelInterface, args ... string) (interface{}, error) {
	if obj == nil {
		return nil, ErrObjectEmpty
	}
	if obj.GetId() == "" {
		return nil, ErrIdEmpty
	}
	argObj := dao.GetArgInstance()
	argObj.SetIdEqual(obj.GetId())
	argObj.SetUpdateObject(obj)
	argObj.SetUpdateColumnNames(args)
	sql, params := arg.BuildUpdateSql(argObj)
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return obj, errors.Wrap(err, "conn:"+sql)
	}
	defer stmt.Close()
	_, err = stmt.Exec(params...)
	if err != nil {
		return obj, errors.Wrap(err, "stmt:"+sql)
	}
	return obj, nil
}

func BatchUpdate(dao BaseDaoInterface,argObj arg.BaseArgInterface) (int64, error) {
	sql, params := arg.BuildUpdateSql(argObj)
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return 0, errors.Wrap(err, "conn:"+sql)
	}
	defer stmt.Close()
	result, err := stmt.Exec(params...)
	if err != nil {
		return 0, errors.Wrap(err, "stmt:"+sql)
	}
	return result.RowsAffected()
}

func Delete(dao BaseDaoInterface,argObj arg.BaseArgInterface) (int64, error) {
	sql, params := arg.BuildDeleteSql(argObj)
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return 0, errors.Wrap(err, "conn:"+sql)
	}
	defer stmt.Close()
	result, err := stmt.Exec(params...)
	if err != nil {
		return 0, errors.Wrap(err, "stmt:"+sql)
	}
	return result.RowsAffected()

}

func LogicalDelete(dao BaseDaoInterface,argObj arg.BaseArgInterface) (int64, error) {
	sql, params := arg.BuildLogicalDeleteSql(argObj)
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return 0, errors.Wrap(err, "conn:"+sql)
	}
	defer stmt.Close()
	result, err := stmt.Exec(params...)
	if err != nil {
		return 0, errors.Wrap(err, "stmt:"+sql)
	}
	return result.RowsAffected()
}

//func SqlQuery(dao BaseDaoInterface,args ...interface{})