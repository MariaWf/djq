package model

import "errors"

type Admin struct {
	Id             string `form:"id" json:"id" db:"id" desc:"id"`
	Name           string `form:"name" json:"name" db:"name" desc:"名称"`
	Mobile         string `form:"mobile" json:"mobile" db:"mobile" desc:"手机"`
	Password       string `form:"password" json:"password" db:"password" desc:"密码"`
	Locked         bool   `form:"locked" json:"locked" db:"locked" desc:"锁定"`

	RoleList       []*Role       `json:"roleList" desc:"角色列表"`
	PermissionList []*Permission `json:"permissionList" desc:"权限列表"`
}

func (obj *Admin) GetId() string {
	return obj.Id
}

func (obj *Admin) SetId(id string) {
	obj.Id = id
}

func (obj *Admin) GetTableName() string {
	return "tbl_admin"
}

func (obj *Admin) GetDBNames() []string {
	return []string{"id", "name", "mobile", "password", "locked"}
}

func (obj *Admin) GetMapNames() []string {
	return []string{"id", "name", "mobile", "password", "locked"}
}

func (obj *Admin) GetValue4Map(name string) interface{} {
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

func (obj *Admin) GetDBFromMapName(name string) string {
	str := GetDBFromMapName(obj, name)
	if str != "" {
		return str
	}
	panic(errors.New("对象admin属性[" + name + "]不存在"))
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


func (obj *Admin) SetRoleListFromInterfaceArr(list []interface{}) {
	if list != nil && len(list) != 0 {
		obj.RoleList = make([]*Role, len(list), len(list))
		for i, role := range list {
			obj.RoleList[i] = role.(*Role)
		}
	} else {
		obj.RoleList = make([]*Role, 0, 1)
	}
}