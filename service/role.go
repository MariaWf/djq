package service

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao"
	"mimi/djq/dao/arg"
	"mimi/djq/db/mysql"
	"mimi/djq/model"
	"mimi/djq/util"
)

type Role struct {
}

func (service *Role) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.Role{conn}
}

func (service *Role) Get(id string) (*model.Role, error) {
	conn, err := mysql.Get()
	if err != nil {
		return nil, checkErr(err)
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn)
	obj, err := dao.Get(daoObj, id)
	if err != nil {
		rollback = true
		return nil, checkErr(err)
	}
	obj.(*model.Role).BindStr2PermissionList()
	return obj.(*model.Role), nil
}

func (service *Role) Add(obj *model.Role) (*model.Role, error) {
	conn, err := mysql.Get()
	if err != nil {
		return nil, checkErr(err)
	}
	err = service.checkAdd(obj)
	if err != nil {
		return nil, err
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn)
	obj.BindPermissionList2Str()
	_, err = dao.Add(daoObj, obj)
	if err != nil {
		rollback = true
		return obj, checkErr(err)
	}
	return obj, nil
}

func (service *Role) Update(obj *model.Role) (*model.Role, error) {
	conn, err := mysql.Get()
	if err != nil {
		return nil, checkErr(err)
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn)
	obj.BindPermissionList2Str()
	_, err = dao.Update(daoObj, obj, "name", "description", "permissionListStr")
	if err != nil {
		rollback = true
		return obj, checkErr(err)
	}
	return obj, nil
}

func (service *Role) Delete(roleIds ...string) (int64, error) {
	if roleIds == nil || len(roleIds) == 0 {
		return 0, ErrIdEmpty
	}
	conn, err := mysql.Get()
	if err != nil {
		return 0, checkErr(err)
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.Role)
	argObj := &arg.Role{}
	argObj.SetIdsIn(roleIds)
	count, err := dao.LogicalDelete(daoObj, argObj)
	if err != nil {
		rollback = true
		return 0, checkErr(err)
	}
	_, err = daoObj.DeleteRelationshipWithAdmin(roleIds...)
	if err != nil {
		rollback = true
		return 0, checkErr(err)
	}
	return count, nil
}

func (service *Role) check(obj *model.Role) error {
	if obj == nil {
		return ErrObjectEmpty
	}
	if !util.MatchLen(obj.Name, 2, 32) {
		return errors.New("名称长度为2到32")
	}
	return nil
}

func (service *Role) checkUpdate(obj *model.Role) error {
	if obj != nil && obj.Id == "" {
		return ErrIdEmpty
	}
	return service.check(obj)
}

func (service *Role) checkAdd(obj *model.Role) error {
	return service.check(obj)
}
