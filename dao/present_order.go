package dao

import (
	"database/sql"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
)

type PresentOrder struct {
	Conn *sql.Tx
}

func (dao *PresentOrder) GetConn() *sql.Tx {
	return dao.Conn
}
func (dao *PresentOrder) GetArgInstance() arg.BaseArgInterface {
	return &arg.PresentOrder{}
}
func (dao *PresentOrder) GetModelInstance() model.BaseModelInterface {
	return &model.PresentOrder{}
}
