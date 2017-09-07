package dao

import (
	"testing"
	"mimi/djq/model"
	"strconv"
	"mimi/djq/db/mysql"
	"mimi/djq/dao/arg"
	"fmt"
	"math/rand"
)

func TestRole_Find(t *testing.T) {

	conn, err := mysql.Get()
	if err != nil {
		t.Error(err)
	}
	defer mysql.Rollback(conn)
	dao := &Role{conn}
	objList, err := bathAdd(dao, 15)
	if err != nil {
		t.Error(err)
	}
	if len(objList) != 15 {
		t.Error(len(objList), 15, "len not equal")
	}
	arg := &arg.Role{}
	arg.NameLike = "3"
	list, err := Find(dao, arg)
	if err != nil {
		t.Error(err)
	}
	for _, obj := range list {
		t.Log(obj.(*model.Role).Name)
	}

}

func TestRole_LogicalDelete(t *testing.T) {
	conn, err := mysql.Get()
	if err != nil {
		t.Error(err)
	}
	defer mysql.Rollback(conn)
	dao := &Role{conn}

	oldCount, err := batchDelete(dao)
	if err != nil {
		t.Error(err)
	}
	t.Log("旧数据量", oldCount)

	objList, err := bathAdd(dao, 15)
	if err != nil {
		t.Error(err)
	}
	if len(objList) != 15 {
		t.Error(len(objList), 15, "len not equal")
	}

	argObj := &arg.Role{}
	argObj.NameLike = "3"
	list, err := Find(dao, argObj)
	if err != nil {
		t.Error(err)
	}
	if len(list) != 2 {
		t.Error(len(list), 2, "len not equal")
	}

	count, err := LogicalDelete(dao, argObj)
	if err != nil {
		t.Error(err)
	}
	if count != 2 {
		t.Error(count, 2, "logical delete len not equal")
	}

	arg2 := &arg.Role{}
	total, err := Count(dao, arg2)
	if err != nil {
		t.Error(err)
	}
	if total != 13 {
		t.Error(total, 13, "logical delete len not equal")
	}

	arg2.IncludeDeleted = true
	total, err = Count(dao, arg2)
	if err != nil {
		t.Error(err)
	}
	if total != 15 {
		t.Error(total, 15, "logical delete len not equal")
	}

}

func TestRole_Update(t *testing.T) {
	conn, err := mysql.Get()
	if err != nil {
		t.Error(err)
	}
	defer mysql.Rollback(conn)
	dao := &Role{conn}

	oldCount, err := batchDelete(dao)
	if err != nil {
		t.Error(err)
	}
	t.Log("旧数据量", oldCount)

	objList, err := bathAdd(dao, 15)
	if err != nil {
		t.Error(err)
	}
	if len(objList) != 15 {
		t.Error(len(objList), 15, "len not equal")
	}

	argObj := &arg.Role{}
	argObj.NameLike = "3"
	argObj.UpdateColumnNames = []string{"name", "description"}
	argObj.UpdateObject = &model.Role{Name:"updateName", Description:"updateDesc"}
	count, err := BatchUpdate(dao, argObj)
	if err != nil {
		t.Error(err)
	}
	if count != 2 {
		t.Error(count, 2, "update len not equal")
	}
	list, err := findAll(dao)
	if err != nil {
		t.Error(err)
	}

	list[0].(*model.Role).Name = "newUpdateName"
	list[0].(*model.Role).Description = "newUpdateDesc"
	fmt.Println(list[0].(*model.Role).GetId())
	newRole, err := Update(dao, list[0].(*model.Role), "name")
	if err != nil {
		t.Error(err)
	}

	id := list[0].(*model.Role).Id
	roleObj, err := Get(dao, id)
	if err != nil {
		t.Error(err)
	}
	if newRole.(*model.Role).Name != roleObj.(*model.Role).Name {
		t.Error(newRole.(*model.Role).Name, roleObj.(*model.Role).Name)
	}
	if newRole.(*model.Role).Description == roleObj.(*model.Role).Description {
		t.Error(newRole.(*model.Role).Description, roleObj.(*model.Role).Description)
	}
	t.Log(roleObj)
}

func TestRoleUpdate(t *testing.T) {
	conn, err := mysql.Get()
	if err != nil {
		t.Error(err)
	}
	defer mysql.Rollback(conn)
	dao := &Role{conn}

	oldCount, err := batchDelete(dao)
	if err != nil {
		t.Error(err)
	}
	t.Log("旧数据量", oldCount)

	objList, err := bathAdd(dao, 15)
	if err != nil {
		t.Error(err)
	}
	if len(objList) != 15 {
		t.Error(len(objList), 15, "len not equal")
	}
	findAll(dao)
	changeCount := 10
	argObj := &arg.Role{}
	ids := make([]string, 0, changeCount)
	for i, obj := range objList {
		if i == changeCount {
			break
		}
		ids = append(ids, obj.GetId())
	}
	argObj.SetIdsIn(ids)
	argObj.UpdateColumnNames = []string{"name", "description"}
	argObj.UpdateObject = &model.Role{Name:"updateName", Description:"updateDesc"}
	count, err := BatchUpdate(dao, argObj)
	if err != nil {
		t.Error(err)
	}
	if int(count) != changeCount {
		t.Error(count, changeCount, "update len not equal")
	}
	list, err := findAll(dao)
	if err != nil {
		t.Error(err)
	}

	list[0].(*model.Role).Name = "newUpdateName"
	list[0].(*model.Role).Description = "newUpdateDesc"
	fmt.Println(list[0].(*model.Role).GetId())
}

func batchDelete(dao *Role) (int64, error) {
	argObj := &arg.Role{}
	argObj.IncludeDeleted = true
	return Delete(dao, argObj)
}

func bathAdd(dao *Role, total int) ([]*model.Role, error) {
	objList := make([]*model.Role, 0, total)
	permissionList := model.GetPermissionList()
	for i := 0; i < total; i++ {
		obj := &model.Role{}
		obj.Name = "name" + strconv.Itoa(i)
		pl := make([]*model.Permission, 0, 5)
		for _, pm := range permissionList {
			if rand.Int31n(5) > 2 {
				pl = append(pl, pm)
			}
		}
		obj.PermissionList = pl
		obj.BindPermissionList2Str()
		_, err := Add(dao, obj)
		if err != nil {
			return objList, err
		}
		objList = append(objList, obj)
	}
	return objList, nil
}

func findAll(dao *Role) ([]interface{}, error) {
	argObj := &arg.Role{}
	argObj.ShowColumnNames = []string{"id", "name", "description", "permissionListStr"}
	argObj.OrderBy = "name"
	list, err := Find(dao, argObj)
	if err != nil {
		return nil, err
	}
	for _, role := range list {
		fmt.Println(role)
	}
	return list, err
}