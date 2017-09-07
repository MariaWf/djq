package dao

import (
	"database/sql"
	"mimi/djq/model"
	"mimi/djq/dao/arg"
	"github.com/pkg/errors"
	"strings"
	"mimi/djq/util"
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

//func (dao *Admin) RefreshRelationshipWithRole(obj *model.Admin) (int64, error) {
//	if obj == nil {
//		return ErrObjectEmpty
//	}
//	if obj.GetId() == "" {
//		return ErrIdEmpty
//	}
//	if obj.RoleList == nil || len(obj.RoleList) == 0 {
//		return dao.DeleteRelationshipWithRoleByAdminId(obj.GetId())
//	}
//	roleIds := make([]string,0,len(obj.RoleList))
//	count,err := dao.DeleteRelationshipWithRoleBy2AddRoleIds(roleIds...)
//
//	existRoleIds,err := dao.ListRoleIdByAdminId(adminId)
//
//	sql := "update tr_admin_role set del_flag = false where role_id in ("
//	for i, _ := range roleIds {
//		if i != 0 {
//			sql += ","
//		}
//		sql += "?"
//	}
//	sql += ");"
//	stmt, err := dao.GetConn().Prepare(sql)
//	if err != nil {
//		return 0, errors.Wrap(err, "conn:" + sql)
//	}
//	defer stmt.Close()
//	result, err := stmt.Exec(roleIds)
//	if err != nil {
//		return 0, errors.Wrap(err, "stmt:" + sql)
//	}
//	return result.RowsAffected()
//}


func (dao *Admin) AddRelationshipWithRole(adminId string, roleId string) error {
	sql := "insert into tr_admin_role (id,admin_id,role_id) values (?,?,?);"
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return errors.Wrap(err, "conn:" + sql)
	}
	defer stmt.Close()
	id := BuildId()
	_, err = stmt.Exec(id, adminId, roleId)
	if err != nil {
		return errors.Wrap(err, "stmt:" + sql)
	}
	return nil
}

func (dao *Admin) DeleteRelationshipWithRoleByAdminId(adminIds ... string) (int64, error) {
	placeholderList := make([]string, 0, len(adminIds))
	for i := 0; i < len(adminIds); i++ {
		placeholderList = append(placeholderList, "?")
	}
	sql := "update tr_admin_role set del_flag = true where admin_id in (" + strings.Join(placeholderList, ",") + ") and del_flag = false;"
	stmt, err := dao.GetConn().Prepare(sql)
	if err != nil {
		return 0, errors.Wrap(err, "conn:" + sql)
	}
	defer stmt.Close()
	result, err := stmt.Exec(util.StringArrConvert2InterfaceArr(adminIds)...)
	if err != nil {
		return 0, errors.Wrap(err, "stmt:" + sql)
	}
	return result.RowsAffected()
}


//type AdminDao struct {
//	Conn *sql.Tx
//}
//
//func (u *AdminDao)SetConn(conn *sql.Tx) {
//	u.Conn = conn
//}
//
//func (u *AdminDao)Find(arg *arg.Admin) ([]*model.Admin, error) {
//	admins := make([]*model.Admin, 0, arg.PageSize)
//	sql, params, columnNames := arg.BuildFindSql()
//	stmt, err := u.Conn.Prepare(sql)
//	if err != nil {
//		return nil, errors.Wrap(err, "conn")
//	}
//	defer stmt.Close()
//	rows, err := stmt.Query(params...)
//	if err != nil {
//		return nil, errors.Wrap(err, "stmt")
//	}
//	defer rows.Close()
//	for rows.Next() {
//		admin := new(model.Admin)
//		err = rows.Scan(admin.GetPointers4DB(columnNames)...)
//		if err != nil {
//			return nil, errors.Wrap(err, "rows")
//		}
//		admins = append(admins, admin)
//	}
//	return admins, nil
//	return nil,nil
//}
//
//func (u *AdminDao)Count(arg *arg.Admin) (int, error) {
//	sql, params := arg.BuildCountSql()
//	stmt, err := u.Conn.Prepare(sql)
//	if err != nil {
//		return 0, errors.Wrap(err, "conn")
//	}
//	defer stmt.Close()
//	rows, err := stmt.Query(params...)
//	if err != nil {
//		return 0, errors.Wrap(err, "stmt")
//	}
//	defer rows.Close()
//	for rows.Next() {
//		total := 0
//		err = rows.Scan(&total)
//		if err != nil {
//			return 0, errors.Wrap(err, "rows")
//		}
//		return total, nil
//	}
//	return 0, errors.Wrap(ErrUnknown, "dao:admin:count")
//}
//
//func (u *AdminDao)Get(id string) (*model.Admin, error) {
//	if id == "" {
//		return nil, ErrIdEmpty
//	}
//	arg := &arg.Admin{}
//	arg.IdEqual = id
//	arg.PageSize = 1
//	arg.TargetPage = util.BeginPage
//	sql, params, columnNames := arg.BuildFindSql()
//	stmt, err := u.Conn.Prepare(sql)
//	if err != nil {
//		return nil, err
//	}
//	defer stmt.Close()
//	rows, err := stmt.Query(params...)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//	for rows.Next() {
//		admin := new(model.Admin)
//		err = rows.Scan(admin.GetPointers4DB(columnNames)...)
//		if err != nil {
//			return nil, err
//		}
//		return admin, nil
//	}
//	return nil, ErrObjectNotFound
//}
//
//func (u *AdminDao)Add(admin *model.Admin) (*model.Admin, error) {
//	if admin == nil {
//		return nil, ErrObjectEmpty
//	}
//	arg := &arg.Admin{}
//	sql, columnNames := arg.BuildInsertSql()
//	stmt, err := u.Conn.Prepare(sql)
//	if err != nil {
//		return nil, errors.Wrap(err, "conn")
//	}
//	defer stmt.Close()
//	id := BuildId()
//	params := admin.GetValues4DB(columnNames)[:]
//	params[0] = id
//	_, err = stmt.Exec(params...)
//	if err != nil {
//		return nil, errors.Wrap(err, "stmt")
//	}
//	admin.Id = id
//	return admin, nil
//}
//
//func (u *AdminDao)Update(admin *model.Admin, args ... string) (*model.Admin, error) {
//	if admin == nil {
//		return nil, ErrObjectEmpty
//	}
//	if admin.Id == "" {
//		return nil, ErrIdEmpty
//	}
//	arg := &arg.Admin{}
//	arg.IdEqual = admin.Id
//	arg.UpdateObject = admin
//	arg.UpdateColumnNames = args
//	sql, params := arg.BuildUpdateSql()
//	stmt, err := u.Conn.Prepare(sql)
//	if err != nil {
//		return admin, errors.Wrap(err, "conn")
//	}
//	defer stmt.Close()
//	_, err = stmt.Exec(params...)
//	if err != nil {
//		return admin, errors.Wrap(err, "stmt")
//	}
//	return admin, nil
//}
//
//func (u *AdminDao)BatchUpdate(arg *arg.Admin) (int64, error) {
//	sql, params := arg.BuildUpdateSql()
//	stmt, err := u.Conn.Prepare(sql)
//	if err != nil {
//		return 0, errors.Wrap(err, "conn")
//	}
//	defer stmt.Close()
//	result, err := stmt.Exec(params...)
//	if err != nil {
//		return 0, errors.Wrap(err, "stmt")
//	}
//	return result.RowsAffected()
//}
//
//func (u *AdminDao)Delete(arg *arg.Admin) (int64, error) {
//	sql, params := arg.BuildDeleteSql()
//	stmt, err := u.Conn.Prepare(sql)
//	if err != nil {
//		return 0, errors.Wrap(err, "conn")
//	}
//	defer stmt.Close()
//	result, err := stmt.Exec(params...)
//	if err != nil {
//		return 0, errors.Wrap(err, "stmt")
//	}
//	return result.RowsAffected()
//
//}
//
//func (u *AdminDao)LogicalDelete(arg *arg.Admin) (int64, error) {
//	sql, params := arg.BuildUpdateSql()
//	stmt, err := u.Conn.Prepare(sql)
//	if err != nil {
//		return 0, errors.Wrap(err, "conn")
//	}
//	defer stmt.Close()
//	result, err := stmt.Exec(params...)
//	if err != nil {
//		return 0, errors.Wrap(err, "stmt")
//	}
//	return result.RowsAffected()
//
//}
