package service

import (
	"database/sql"
	"mimi/djq/dao"
	"mimi/djq/model"
	"mimi/djq/db/mysql"
	"mimi/djq/util"
	"mimi/djq/dao/arg"
)

type Admin struct {

}

func (service *Admin) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.Admin{conn}
}

func (service *Admin) Get(id string) (*model.Admin, error) {
	conn, err := mysql.Get()
	if err != nil {
		return nil, checkErr(err)
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.Admin)
	obj, err := dao.Get(daoObj, id)
	if err != nil {
		rollback = true
		return nil, checkErr(err)
	}

	roleDao := &dao.Role{conn}
	roleIds, err := roleDao.ListRoleIdsByAdminId(id)
	if err != nil {
		rollback = true
		return nil, checkErr(err)
	}

	if roleIds != nil && len(roleIds) != 0 {
		argRole := roleDao.GetArgInstance().(*arg.Role)
		argRole.SetIdsIn(roleIds)
		roles, err := dao.Find(roleDao, argRole)
		if err != nil {
			rollback = true
			return nil, checkErr(err)
		}
		obj.(*model.Admin).SetRoleListFromInterfaceArr(roles)
		obj.(*model.Admin).BindPermissionList()
	}
	return obj.(*model.Admin), nil
}

func (service *Admin) Add(obj *model.Admin) (*model.Admin, error) {
	conn, err := mysql.Get()
	if err != nil {
		return nil, checkErr(err)
	}

	sourcePassword, err := util.DecryptPassword(obj.Password)
	if err != nil {
		return nil, checkErr(err)
	}
	obj.Password = sourcePassword

	err = service.checkAdd(obj)
	if err != nil {
		return nil, err
	}

	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.Admin)

	_, err = dao.Add(daoObj, obj)
	if err != nil {
		rollback = true
		return nil, checkErr(err)
	}

	if obj.RoleList != nil && len(obj.RoleList) != 0 {
		toAddRoleIds := make([]string, len(obj.RoleList), len(obj.RoleList))
		for i, role := range obj.RoleList {
			toAddRoleIds[i] = role.GetId()
		}
		if !util.IsStringArrEmpty(toAddRoleIds) {
			for _, roleId := range toAddRoleIds {
				err := daoObj.AddRelationshipWithRole(obj.GetId(), roleId)
				if err != nil {
					rollback = true
					return nil, err
				}
			}
		}
	}
	return obj, nil
}

func (service *Admin) refreshRelationshipWithRole(conn *sql.Tx, obj *model.Admin) error {
	daoObj := service.GetDaoInstance(conn).(*dao.Admin)

	if obj.RoleList == nil && len(obj.RoleList) == 0 {
		_, err := daoObj.DeleteRelationshipWithRoleByAdminId(obj.GetId())
		if err != nil {
			return err
		}
		return nil
	}

	toUpdateRoleIds := make([]string, 0, len(obj.RoleList))
	for i, role := range obj.RoleList {
		toUpdateRoleIds[i] = role.GetId()
	}

	roleDao := &dao.Role{conn}

	existRoleIds, err := roleDao.ListRoleIdsByAdminId(obj.GetId())
	if err != nil {
		return err
	}

	var toAddRoleIds []string
	var toDeleteRoleIds []string
	if existRoleIds == nil || len(existRoleIds) == 0 {
		toAddRoleIds = util.StringArrCopy(toUpdateRoleIds)
		toDeleteRoleIds = nil
	} else {
		toAddRoleIds = make([]string, 0, 10)
		toDeleteRoleIds := util.StringArrCopy(existRoleIds)
		for _, toUpdateRoleId := range toUpdateRoleIds {
			exist := false
			for _, existRoleId := range existRoleIds {
				if toUpdateRoleId == existRoleId {
					exist = true
					toDeleteRoleIds = util.StringArrDelete(toDeleteRoleIds, toUpdateRoleId)
				}
			}
			if !exist {
				toAddRoleIds = append(toAddRoleIds, toUpdateRoleId)
			}
		}
	}
	if !util.IsStringArrEmpty(toAddRoleIds) {
		for _, roleId := range toAddRoleIds {
			err := daoObj.AddRelationshipWithRole(obj.GetId(), roleId)
			if err != nil {
				return err
			}
		}
	}
	if !util.IsStringArrEmpty(toDeleteRoleIds) {
		_, err := roleDao.DeleteRelationshipWithAdmin(toDeleteRoleIds...)
		if err != nil {
			return err
		}
	}

	return nil
}

func (service *Admin) Update(obj *model.Admin) (*model.Admin, error) {
	conn, err := mysql.Get()
	if err != nil {
		return nil, checkErr(err)
	}

	sourcePassword, err := util.DecryptPassword(obj.Password)
	if err != nil {
		return nil, checkErr(err)
	}
	obj.Password = sourcePassword

	err = service.checkUpdate(obj)
	if err != nil {
		return nil, err
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn)

	if obj.Password != "" {
		obj.Password = util.BuildPassword4DB(obj.Password)
	}

	_, err = dao.Update(daoObj, obj, "name", "mobile", "password", "Locked")
	if err != nil {
		rollback = true
		return nil, checkErr(err)
	}
	err = service.refreshRelationshipWithRole(conn, obj)
	obj.Password = ""
	if err != nil {
		rollback = true
		return nil, checkErr(err)
	}
	return obj, checkErr(err)
}

func (service *Admin) Count() {

}

func (service *Admin) Delete(ids ... string) (int64, error) {
	if ids == nil || len(ids) == 0 {
		return 0, nil
	}
	conn, err := mysql.Get()
	if err != nil {
		return 0, checkErr(err)
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.Admin)
	argObj := daoObj.GetArgInstance().(*arg.Admin)
	argObj.SetIdsIn(ids)
	count, err := dao.LogicalDelete(daoObj, argObj)
	if err != nil {
		rollback = true
		return 0, checkErr(err)
	}
	_, err = daoObj.DeleteRelationshipWithRoleByAdminId(ids...)
	if err != nil {
		rollback = true
		return 0, checkErr(err)
	}
	return count, checkErr(err)

}

func (service *Admin) check(obj *model.Admin) error {
	if obj == nil {
		return ErrObjectEmpty
	}
	if !util.MatchName(obj.Name) {
		return util.ErrNameFormat
	}
	if obj.Mobile != "" && !util.MatchMobile(obj.Mobile) {
		return util.ErrMobileFormat
	}
	if obj.Password != "" && !util.MatchPassword(obj.Password) {
		return util.ErrPasswordFormat
	}
	return nil
}

func (service *Admin) checkUpdate(obj *model.Admin) error {
	if obj != nil && obj.Id == "" {
		return ErrIdEmpty
	}
	return service.check(obj)
}

func (service *Admin) checkAdd(obj *model.Admin) error {
	if obj != nil && obj.Password == "" {
		return util.ErrPasswordFormat
	}
	return service.check(obj)
}