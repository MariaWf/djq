package dao

import (
	"database/sql"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
)

type ShopAccount struct {
	Conn *sql.Tx
}

func (dao *ShopAccount) GetConn() *sql.Tx {
	return dao.Conn
}
func (dao *ShopAccount) GetArgInstance() arg.BaseArgInterface {
	return &arg.ShopAccount{}
}
func (dao *ShopAccount) GetModelInstance() model.BaseModelInterface {
	return &model.ShopAccount{}
}
