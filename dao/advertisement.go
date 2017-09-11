package dao

import (
	"database/sql"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
)

type Advertisement struct {
	Conn *sql.Tx
}

func (dao *Advertisement) GetConn() *sql.Tx {
	return dao.Conn
}
func (dao *Advertisement) GetArgInstance() arg.BaseArgInterface {
	return &arg.Advertisement{}
}
func (dao *Advertisement) GetModelInstance() model.BaseModelInterface {
	return &model.Advertisement{}
}
