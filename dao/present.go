package dao

import (
	"database/sql"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
)

type Present struct {
	Conn *sql.Tx
}

func (dao *Present) GetConn() *sql.Tx {
	return dao.Conn
}
func (dao *Present) GetArgInstance() arg.BaseArgInterface {
	return &arg.Present{}
}
func (dao *Present) GetModelInstance() model.BaseModelInterface {
	return &model.Present{}
}
