package dao

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/util"
	"strings"
)

type Admin struct {
	Conn *sql.Tx
}

func (dao *Admin) GetConn() *sql.Tx {
	return dao.Conn
}
func (dao *Admin) GetArgInstance() arg.BaseArgInterface {
	return &arg.Admin{}
}
func (dao *Admin) GetModelInstance() model.BaseModelInterface {
	return &model.Admin{}
}

func (dao *Admin) AddRelationshipWithRole(adminId string, roleId string) error {
	sql := "insert into tr_admin_role (id,admin_id,role_id) values (?,?,?);"
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return errors.Wrap(err, "conn:"+sql)
	}
	defer stmt.Close()
	id := BuildId()
	_, err = stmt.Exec(id, adminId, roleId)
	if err != nil {
		return errors.Wrap(err, "stmt:"+sql)
	}
	return nil
}

func (dao *Admin) DeleteRelationshipWithRoleByAdminId(adminIds ...string) (int64, error) {
	placeholderList := make([]string, 0, len(adminIds))
	for i := 0; i < len(adminIds); i++ {
		placeholderList = append(placeholderList, "?")
	}
	sql := "update tr_admin_role set del_flag = true where admin_id in (" + strings.Join(placeholderList, ",") + ") and del_flag = false;"
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return 0, errors.Wrap(err, "conn:"+sql)
	}
	defer stmt.Close()
	result, err := stmt.Exec(util.StringArrConvert2InterfaceArr(adminIds)...)
	if err != nil {
		return 0, errors.Wrap(err, "stmt:"+sql)
	}
	return result.RowsAffected()
}
