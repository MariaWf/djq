package dao

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/util"
)

type ShopClassification struct {
	Conn *sql.Tx
}

func (dao *ShopClassification) GetConn() *sql.Tx {
	return dao.Conn
}
func (dao *ShopClassification) GetArgInstance() arg.BaseArgInterface {
	return &arg.ShopClassification{}
}
func (dao *ShopClassification) GetModelInstance() model.BaseModelInterface {
	return &model.ShopClassification{}
}

func (dao *ShopClassification) DeleteRelationshipWithShop(shopClassificationIds ...string) (int64, error) {
	if shopClassificationIds == nil || len(shopClassificationIds) == 0 {
		return 0, ErrIdEmpty
	}
	sql := "update tr_shop_classification_shop set del_flag = true where shop_classification_id in ("
	for i, _ := range shopClassificationIds {
		if i != 0 {
			sql += ","
		}
		sql += "?"
	}
	sql += ");"
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return 0, errors.Wrap(err, "conn:"+sql)
	}
	defer stmt.Close()
	result, err := stmt.Exec(util.StringArrConvert2InterfaceArr(shopClassificationIds)...)
	if err != nil {
		return 0, errors.Wrap(err, "stmt:"+sql)
	}
	return result.RowsAffected()
}

func (dao *ShopClassification) ListShopClassificationIdsByShopId(shopId string) ([]string, error) {
	if shopId == "" {
		return nil, ErrIdEmpty
	}
	sql := "select distinct shop_classification_id from tr_shop_classification_shop where shop_id = ? and del_flag = false;"
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return nil, errors.Wrap(err, "conn:"+sql)
	}
	defer stmt.Close()
	rows, err := stmt.Query(shopId)
	if err != nil {
		return nil, errors.Wrap(err, "stmt:"+sql)
	}
	defer rows.Close()
	shopClassificationIds := make([]string, 0, 10)
	for rows.Next() {
		var shopClassificationId string
		err = rows.Scan(&shopClassificationId)
		if err != nil {
			return nil, errors.Wrap(err, "rows:"+sql)
		}
		shopClassificationIds = append(shopClassificationIds, shopClassificationId)
	}
	return shopClassificationIds, nil
}
