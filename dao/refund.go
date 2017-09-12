package dao

import (
	"database/sql"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
)

type Refund struct {
	Conn *sql.Tx
}

func (dao *Refund) GetConn() *sql.Tx {
	return dao.Conn
}
func (dao *Refund) GetArgInstance() arg.BaseArgInterface {
	return &arg.Refund{}
}
func (dao *Refund) GetModelInstance() model.BaseModelInterface {
	return &model.Refund{}
}
