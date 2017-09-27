package dao

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/util"
	"strings"
)

type Shop struct {
	Conn *sql.Tx
}

func (dao *Shop) GetConn() *sql.Tx {
	return dao.Conn
}
func (dao *Shop) GetArgInstance() arg.BaseArgInterface {
	return &arg.Shop{}
}
func (dao *Shop) GetModelInstance() model.BaseModelInterface {
	return &model.Shop{}
}

func (dao *Shop) AddRelationshipWithShopClassification(shopId string, shopClassificationId string) error {
	sql := "insert into tr_shop_classification_shop (id,shop_id,shop_classification_id) values (?,?,?);"
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return errors.Wrap(err, "conn:"+sql)
	}
	defer stmt.Close()
	id := BuildId()
	_, err = stmt.Exec(id, shopId, shopClassificationId)
	if err != nil {
		return errors.Wrap(err, "stmt:"+sql)
	}
	return nil
}

func (dao *Shop) DeleteRelationshipWithShopClassificationByShopId(shopIds ...string) (int64, error) {
	placeholderList := make([]string, 0, len(shopIds))
	for i := 0; i < len(shopIds); i++ {
		placeholderList = append(placeholderList, "?")
	}
	sql := "update tr_shop_classification_shop set del_flag = true where shop_id in (" + strings.Join(placeholderList, ",") + ") and del_flag = false;"
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return 0, errors.Wrap(err, "conn:"+sql)
	}
	defer stmt.Close()
	result, err := stmt.Exec(util.StringArrConvert2InterfaceArr(shopIds)...)
	if err != nil {
		return 0, errors.Wrap(err, "stmt:"+sql)
	}
	return result.RowsAffected()
}


func (dao *Shop) ListShopIdsByShopClassificationId(shopClassificationId string) ([]string, error) {
	if shopClassificationId == "" {
		return nil, ErrIdEmpty
	}
	sql := "select distinct shop_id from tr_shop_classification_shop where shop_classification_id = ? and del_flag = false;"
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return nil, errors.Wrap(err, "conn:"+sql)
	}
	defer stmt.Close()
	rows, err := stmt.Query(shopClassificationId)
	if err != nil {
		return nil, errors.Wrap(err, "stmt:"+sql)
	}
	defer rows.Close()
	shopIds := make([]string, 0, 10)
	for rows.Next() {
		var shopClassificationId string
		err = rows.Scan(&shopClassificationId)
		if err != nil {
			return nil, errors.Wrap(err, "rows:"+sql)
		}
		shopIds = append(shopIds, shopClassificationId)
	}
	return shopIds, nil
}
