package dao

import (
	"database/sql"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
)

type User struct {
	Conn *sql.Tx
}

func (dao *User) GetConn() *sql.Tx {
	return dao.Conn
}
func (dao *User) GetArgInstance() arg.BaseArgInterface {
	return &arg.User{}
}
func (dao *User) GetModelInstance() model.BaseModelInterface {
	return &model.User{}
}