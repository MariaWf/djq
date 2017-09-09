package model

import "errors"

type Admin struct {
	Id       string `form:"id" json:"id" db:"id" desc:"id"`
	Name     string `form:"name" json:"name" db:"name" desc:"名称"`
	Mobile   string `form:"mobile" json:"mobile" db:"mobile" desc:"手机"`
	Password string `form:"password" json:"password" db:"password" desc:"密码"`
	Locked   bool   `form:"locked" json:"locked" db:"locked" desc:"锁定"`

	RoleList       []*Role       `json:"roleList" desc:"角色列表"`
	PermissionList []*Permission `json:"permissionList" desc:"权限列表"`
}

func (obj *Admin) GetId() string {
	return obj.Id
}

func (obj *Admin) SetId(id string) {
	obj.Id = id
}

func (obj *Admin) SetRoleListFromInterfaceArr(list []interface{}) {
	obj.RoleList = make([]*Role, 0, 10)
	if list != nil && len(list) != 0 {
		for _, role := range list {
			obj.RoleList = append(obj.RoleList, role.(*Role))
		}
	}
}

func (obj *Admin) GetPointer4DB(name string) interface{} {
	switch name {
	case "id":
		return &obj.Id
	case "name":
		return &obj.Name
	case "mobile":
		return &obj.Mobile
	case "locked":
		return &obj.Locked
	case "password":
		return &obj.Password
	}
	panic(errors.New("对象admin属性[" + name + "]不存在"))
}

//func (obj *Admin) GetPointers4DB(names []string) []interface{} {
//	pointers := make([]interface{}, 0, 5)
//	for _, name := range names {
//		pointers = append(pointers, obj.GetPointer4DB(name))
//	}
//	return pointers
//}

func (obj *Admin) GetValue4DB(name string) interface{} {
	switch name {
	case "id":
		return obj.Id
	case "name":
		return obj.Name
	case "locked":
		return obj.Locked
	case "mobile":
		return obj.Mobile
	case "password":
		return obj.Password
	}
	panic(errors.New("对象admin属性[" + name + "]不存在"))
}

//func (obj *Admin) GetValues4DB(names []string) []interface{} {
//	values := make([]interface{}, 0, 5)
//	for _, name := range names {
//		values = append(values, obj.GetValue4DB(name))
//	}
//	return values
//}

func (obj *Admin) BindPermissionList() {
	obj.PermissionList = make([]*Permission, 0, 10)
	if obj.RoleList != nil && len(obj.RoleList) != 0 {
		for _, role := range obj.RoleList {
			role.BindStr2PermissionList()
			if role.PermissionList != nil && len(role.PermissionList) != 0 {
				for _, permission := range role.PermissionList {
					exist := false
					for _, existP := range obj.PermissionList {
						if existP.Code == permission.Code {
							exist = true
							break
						}
					}
					if !exist {
						obj.PermissionList = append(obj.PermissionList, permission)
					}
				}
			}
		}
	}
}

func (obj *Admin) GetPermissionCodeList() []string {
	if obj.PermissionList != nil && len(obj.PermissionList) != 0 {
		list := make([]string, len(obj.PermissionList), len(obj.PermissionList))
		for i, pn := range obj.PermissionList {
			list[i] = pn.Code
		}
		return list
	}
	return nil
}
