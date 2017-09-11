package dao

import (
	"database/sql"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"strings"
	"github.com/pkg/errors"
	"mimi/djq/util"
)

type PromotionalPartner struct {
	Conn *sql.Tx
}

func (dao *PromotionalPartner) GetConn() *sql.Tx {
	return dao.Conn
}
func (dao *PromotionalPartner) GetArgInstance() arg.BaseArgInterface {
	return &arg.PromotionalPartner{}
}
func (dao *PromotionalPartner) GetModelInstance() model.BaseModelInterface {
	return &model.PromotionalPartner{}
}

func (dao *PromotionalPartner) DeleteRelationshipWithUserByPromotionalPartnerId(ids ...string) (int64, error) {
	placeholderList := make([]string, 0, len(ids))
	for i := 0; i < len(ids); i++ {
		placeholderList = append(placeholderList, "?")
	}
	sql := "update tbl_user set promotional_partner_id = '' where promotional_partner_id in (" + strings.Join(placeholderList, ",") + ") and del_flag = false;"
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return 0, errors.Wrap(err, "conn:" + sql)
	}
	defer stmt.Close()
	result, err := stmt.Exec(util.StringArrConvert2InterfaceArr(ids)...)
	if err != nil {
		return 0, errors.Wrap(err, "stmt:" + sql)
	}
	return result.RowsAffected()
}