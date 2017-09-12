package dao

import (
	"database/sql"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
)

type RefundReason struct {
	Conn *sql.Tx
}

func (dao *RefundReason) GetConn() *sql.Tx {
	return dao.Conn
}
func (dao *RefundReason) GetArgInstance() arg.BaseArgInterface {
	return &arg.RefundReason{}
}
func (dao *RefundReason) GetModelInstance() model.BaseModelInterface {
	return &model.RefundReason{}
}
