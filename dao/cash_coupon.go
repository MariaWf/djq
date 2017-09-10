package dao

import (
	"database/sql"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
)

type CashCoupon struct {
	Conn *sql.Tx
}

func (dao *CashCoupon) GetConn() *sql.Tx {
	return dao.Conn
}
func (dao *CashCoupon) GetArgInstance() arg.BaseArgInterface {
	return &arg.CashCoupon{}
}
func (dao *CashCoupon) GetModelInstance() model.BaseModelInterface {
	return &model.CashCoupon{}
}
