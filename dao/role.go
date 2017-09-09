package dao

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/util"
)

type Role struct {
	Conn *sql.Tx
}

func (dao *Role) GetConn() *sql.Tx {
	return dao.Conn
}
func (dao *Role) GetArgInstance() arg.BaseArgInterface {
	return &arg.Role{}
}
func (dao *Role) GetModelInstance() model.BaseModelInterface {
	return &model.Role{}
}

func (dao *Role) DeleteRelationshipWithAdmin(roleIds ...string) (int64, error) {
	if roleIds == nil || len(roleIds) == 0 {
		return 0, ErrIdEmpty
	}
	sql := "update tr_admin_role set del_flag = true where role_id in ("
	for i, _ := range roleIds {
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
	result, err := stmt.Exec(util.StringArrConvert2InterfaceArr(roleIds)...)
	if err != nil {
		return 0, errors.Wrap(err, "stmt:"+sql)
	}
	return result.RowsAffected()
}

func (dao *Role) ListRoleIdsByAdminId(adminId string) ([]string, error) {
	if adminId == "" {
		return nil, ErrIdEmpty
	}
	sql := "select distinct role_id from tr_admin_role where admin_id = ? and del_flag = false;"
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return nil, errors.Wrap(err, "conn:"+sql)
	}
	defer stmt.Close()
	rows, err := stmt.Query(adminId)
	if err != nil {
		return nil, errors.Wrap(err, "stmt:"+sql)
	}
	defer rows.Close()
	roleIds := make([]string, 0, 10)
	for rows.Next() {
		var roleId string
		err = rows.Scan(&roleId)
		if err != nil {
			return nil, errors.Wrap(err, "rows:"+sql)
		}
		roleIds = append(roleIds, roleId)
	}
	return roleIds, nil
}