package service

import (
	"math/rand"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"strconv"
	"testing"
)

func TestRole_Update(t *testing.T) {
	list, err := listAllRole()
	service := &Role{}
	if err != nil {
		t.Error(err)
	} else if list != nil && len(list) != 0 {
		role := list[0].(*model.Role)
		role.Name = "mimi"
		role.Description = "mimiDesc"
		_, err := Update(service, role, "name")
		if err != nil {
			t.Error(err)
		} else {
			newObj, err := Get(service, role.GetId())
			if err != nil {
				t.Error(err)
			} else {
				if newObj.(*model.Role).Name != role.Name {
					t.Error(newObj.(*model.Role).Name, role.Name)
				}
				if newObj.(*model.Role).Description == role.Description {
					t.Error(newObj.(*model.Role).Description, role.Description)
				}
				t.Log(newObj)
			}
		}
	}
}

func TestRole_Add(t *testing.T) {
	total := 5
	roles, err := randomAddRole(total)
	if err != nil {
		t.Error(err)
	}
	if len(roles) != 5 {
		t.Error(len(roles), total)
	}
	t.Log(roles)
}

func TestRole_Delete(t *testing.T) {
	count, err := countAllRole()
	if err != nil {
		t.Error(err)
	}
	c2, err := deleteAllRole()
	if err != nil {
		t.Error(err)
	}
	if int64(count) != c2 {
		t.Error(count, c2)
	}
	count, err = countAllRole()
	if count != 0 {
		t.Error(count, err)
	}
}

func randomAddRole(total int) ([]*model.Role, error) {
	service := &Role{}
	objList := make([]*model.Role, 0, total)
	permissionList := model.GetPermissionList()
	for i := 0; i < total; i++ {
		obj := &model.Role{}
		obj.Name = "name" + strconv.Itoa(i)
		obj.Description = "description" + strconv.Itoa(i)
		pl := make([]*model.Permission, 0, 5)
		for _, pm := range permissionList {
			if rand.Int31n(5) > 2 {
				pl = append(pl, pm)
			}
		}
		obj.PermissionList = pl
		obj.BindPermissionList2Str()
		_, err := service.Add(obj)
		if err != nil {
			return objList, err
		}
		objList = append(objList, obj)
	}
	return objList, nil
}

func deleteAllRole() (int64, error) {
	service := &Role{}

	roles, err := listAllRole()
	if err != nil {
		return 0, err
	}
	if roles == nil || len(roles) == 0 {
		return 0, nil
	}
	len := len(roles)
	roleIds := make([]string, 0, len)
	for _, role := range roles {
		roleIds = append(roleIds, role.(*model.Role).GetId())
	}
	return service.Delete(roleIds...)
}

func listAllRole() ([]interface{}, error) {
	service := &Role{}
	argObj := &arg.Role{}
	list, err := Find(service, argObj)
	return list, err
}

func countAllRole() (int, error) {
	service := &Role{}
	argObj := &arg.Role{}
	return Count(service, argObj)
}
