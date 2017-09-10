package dao

import (
	"database/sql"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
)

type ShopIntroductionImage struct {
	Conn *sql.Tx
}

func (dao *ShopIntroductionImage) GetConn() *sql.Tx {
	return dao.Conn
}
func (dao *ShopIntroductionImage) GetArgInstance() arg.BaseArgInterface {
	return &arg.ShopIntroductionImage{}
}
func (dao *ShopIntroductionImage) GetModelInstance() model.BaseModelInterface {
	return &model.ShopIntroductionImage{}
}
