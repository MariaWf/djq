package dao

import (
	"database/sql"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
)

type CashCouponOrder struct {
	Conn *sql.Tx
}

func (dao *CashCouponOrder) GetConn() *sql.Tx {
	return dao.Conn
}
func (dao *CashCouponOrder) GetArgInstance() arg.BaseArgInterface {
	return &arg.CashCouponOrder{}
}
func (dao *CashCouponOrder) GetModelInstance() model.BaseModelInterface {
	return &model.CashCouponOrder{}
}
