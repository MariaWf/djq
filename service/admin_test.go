package service

import (
	"github.com/pkg/errors"
	"math/rand"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/util"
	"strconv"
	"testing"
)

func TestAdmin_Update(t *testing.T) {
	list, err := listAllAdmin()
	service := &Admin{}
	if err != nil {
		t.Error(err)
	} else if list != nil && len(list) != 0 {
		obj := list[0].(*model.Admin)
		obj.Name = "mimi"
		obj.Locked = true
		obj.Mobile = "11111111111"
		_, err := Update(service, obj, "name", "locked")
		if err != nil {
			t.Error(err)
		} else {
			newObj, err := Get(service, obj.GetId())
			if err != nil {
				t.Error(err)
			} else {
				if newObj.(*model.Admin).Name != obj.Name {
					t.Error(newObj.(*model.Admin).Name, obj.Name)
				}
				if newObj.(*model.Admin).Locked != obj.Locked {
					t.Error(newObj.(*model.Admin).Locked, obj.Locked)
				}
				if newObj.(*model.Admin).Mobile == obj.Mobile {
					t.Error(newObj.(*model.Admin).Mobile, obj.Mobile)
				}
				t.Log(newObj)
			}
		}
	}
}

func TestAdmin_Add(t *testing.T) {
	_, err := randomAddAdmin(5)
	if err != nil {
		t.Error(err)
	}
}

func TestAdmin_Delete(t *testing.T) {
	count, err := countAllAdmin()
	if err != nil {
		t.Error(err)
	}
	c2, err := deleteAllAdmin()
	//service := &Admin{}
	//c2, err := service.Delete("76411cc0c43c47fdbbd9b100afc12cbc")

	if err != nil {
		t.Error(err)
	}
	if int64(count) != c2 {
		t.Error(count, c2)
	}
	count, err = countAllAdmin()
	if count != 0 {
		t.Error(count, err)
	}
}

func randomAddAdmin(total int) ([]*model.Admin, error) {
	objList := make([]*model.Admin, 0, total)
	service := &Admin{}
	roles, err := listAllRole()
	if err != nil {
		return nil, errors.Wrap(err, "listRole")
	}
	for i := 0; i < total; i++ {
		obj := &model.Admin{}
		obj.Name = "name" + strconv.Itoa(i)
		obj.Password = "password" + strconv.Itoa(i)
		obj.Password, err = util.EncryptPassword(obj.Password)
		if err != nil {
			return objList, errors.Wrap(err, "encryptPassword")
		}
		//obj.Mobile = "mobile" + strconv.Itoa(i)
		obj.Mobile = "12345678910"
		obj.Locked = rand.Intn(2) > 0
		roleList := make([]*model.Role, 0, 5)
		for _, role := range roles {
			if rand.Int31n(5) > 2 {
				roleList = append(roleList, role.(*model.Role))
			}
		}
		obj.RoleList = roleList
		_, err := service.Add(obj)
		if err != nil {
			return objList, errors.Wrap(err, "addAdmin")
			//return objList, err
		}
		objList = append(objList, obj)
	}
	return objList, nil
}

func deleteAllAdmin() (int64, error) {
	service := &Admin{}

	objList, err := listAllAdmin()
	if err != nil {
		return 0, err
	}
	if objList == nil || len(objList) == 0 {
		return 0, nil
	}
	len := len(objList)
	ids := make([]string, 0, len)
	for _, obj := range objList {
		ids = append(ids, obj.(*model.Admin).GetId())
	}
	return service.Delete(ids...)
}

func listAllAdmin() ([]interface{}, error) {
	service := &Admin{}
	argObj := &arg.Admin{}
	list, err := Find(service, argObj)
	return list, err
}

func countAllAdmin() (int, error) {
	service := &Admin{}
	argObj := &arg.Admin{}
	return Count(service, argObj)
}
