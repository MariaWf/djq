package dao

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/util"
)

type UserDao struct {
	Conn *sql.Tx
}

func (u *UserDao) SetConn(conn *sql.Tx) {
	u.Conn = conn
}

func (u *UserDao) Find(arg *arg.User) ([]*model.User, error) {
	users := make([]*model.User, 0, arg.PageSize)
	sql, params, columnNames := arg.BuildFindSql()
	stmt, err := u.Conn.Prepare(sql)
	if err != nil {
		return nil, errors.Wrap(err, "conn")
	}
	defer stmt.Close()
	rows, err := stmt.Query(params...)
	if err != nil {
		return nil, errors.Wrap(err, "stmt")
	}
	defer rows.Close()
	for rows.Next() {
		user := new(model.User)
		err = rows.Scan(user.GetPointers4DB(columnNames)...)
		if err != nil {
			return nil, errors.Wrap(err, "rows")
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserDao) Count(arg *arg.User) (int, error) {
	sql, params := arg.BuildCountSql()
	stmt, err := u.Conn.Prepare(sql)
	if err != nil {
		return 0, errors.Wrap(err, "conn")
	}
	defer stmt.Close()
	rows, err := stmt.Query(params...)
	if err != nil {
		return 0, errors.Wrap(err, "stmt")
	}
	defer rows.Close()
	for rows.Next() {
		total := 0
		err = rows.Scan(&total)
		if err != nil {
			return 0, errors.Wrap(err, "rows")
		}
		return total, nil
	}
	return 0, errors.Wrap(ErrUnknown, "dao:user:count")
}

func (u *UserDao) Get(id string) (*model.User, error) {
	if id == "" {
		return nil, ErrIdEmpty
	}
	arg := &arg.User{}
	arg.IdEqual = id
	arg.PageSize = 1
	arg.TargetPage = util.BeginPage
	sql, params, columnNames := arg.BuildFindSql()
	stmt, err := u.Conn.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := new(model.User)
		err = rows.Scan(user.GetPointers4DB(columnNames)...)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, ErrObjectNotFound
}

func (u *UserDao) Add(user *model.User) (*model.User, error) {
	if user == nil {
		return nil, ErrObjectEmpty
	}
	arg := &arg.User{}
	sql, columnNames := arg.BuildInsertSql()
	stmt, err := u.Conn.Prepare(sql)
	if err != nil {
		return nil, errors.Wrap(err, "conn")
	}
	defer stmt.Close()
	id := BuildId()
	params := user.GetValues4DB(columnNames)[:]
	params[0] = id
	_, err = stmt.Exec(params...)
	if err != nil {
		return nil, errors.Wrap(err, "stmt")
	}
	user.Id = id
	return user, nil
}

func (u *UserDao) Update(user *model.User, args ...string) (*model.User, error) {
	if user == nil {
		return nil, ErrObjectEmpty
	}
	if user.Id == "" {
		return nil, ErrIdEmpty
	}
	arg := &arg.User{}
	arg.IdEqual = user.Id
	arg.UpdateObject = user
	arg.UpdateColumnNames = args
	sql, params := arg.BuildUpdateSql()
	stmt, err := u.Conn.Prepare(sql)
	if err != nil {
		return user, errors.Wrap(err, "conn")
	}
	defer stmt.Close()
	_, err = stmt.Exec(params...)
	if err != nil {
		return user, errors.Wrap(err, "stmt")
	}
	return user, nil
}

func (u *UserDao) BatchUpdate(arg *arg.User) (int64, error) {
	sql, params := arg.BuildUpdateSql()
	stmt, err := u.Conn.Prepare(sql)
	if err != nil {
		return 0, errors.Wrap(err, "conn")
	}
	defer stmt.Close()
	result, err := stmt.Exec(params...)
	if err != nil {
		return 0, errors.Wrap(err, "stmt")
	}
	return result.RowsAffected()
}

func (u *UserDao) Delete(arg *arg.User) (int64, error) {
	sql, params := arg.BuildDeleteSql()
	stmt, err := u.Conn.Prepare(sql)
	if err != nil {
		return 0, errors.Wrap(err, "conn")
	}
	defer stmt.Close()
	result, err := stmt.Exec(params...)
	if err != nil {
		return 0, errors.Wrap(err, "stmt")
	}
	return result.RowsAffected()

}
