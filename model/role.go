package model

import (
	"errors"
	"strings"
)

//id varchar(64) not null primary key,
//name varchar(32) not null default '',
//description varchar(200) not null default '',
//del_flag tinyint(1) default false

type Role struct {
	Id                string        `form:"id" json:"id" db:"id" desc:"id"`
	Name              string        `form:"name" json:"name" db:"name" desc:"名称"`
	Description       string        `form:"description" json:"description" db:"description" desc:"描述"`
	PermissionListStr string        `form:"permissionListStr" json:"permissionListStr" db:"permission_list_str" desc:"权限列表字符串"`
	PermissionList    []*Permission `json:"permissionList" desc:"权限列表"`
}

func (obj *Role) GetId() string {
	return obj.Id
}

func (obj *Role) SetId(id string) {
	obj.Id = id
}

func (obj *Role) GetPointer4DB(name string) interface{} {
	switch name {
	case "id":
		return &obj.Id
	case "name":
		return &obj.Name
	case "description":
		return &obj.Description
	case "permission_list_str":
		return &obj.PermissionListStr
	}
	panic(errors.New("对象role属性[" + name + "]不存在"))
}

func (obj *Role) GetValue4DB(name string) interface{} {
	switch name {
	case "id":
		return obj.Id
	case "name":
		return obj.Name
	case "description":
		return obj.Description
	case "permission_list_str":
		return obj.PermissionListStr
	}
	panic(errors.New("对象role属性[" + name + "]不存在"))
}

func (obj *Role) BindPermissionList2Str() {
	str := ""
	if obj.PermissionList != nil && len(obj.PermissionList) != 0 {
		permissionList := GetPermissionList()
		for _, permission := range obj.PermissionList {
			for _, permission2 := range permissionList {
				if permission.Code == permission2.Code {
					str += "," + permission.Code
					break
				}
			}
		}
		if str != "" {
			str = str[1:]
		}
	}
	obj.PermissionListStr = str
}

func (obj *Role) BindStr2PermissionList() {
	codes := strings.Split(obj.PermissionListStr, ",")
	oPermissionList := make([]*Permission, 0, len(codes))
	permissionList := GetPermissionList()
	for _, code := range codes {
		for _, permission := range permissionList {
			if code == permission.Code {
				oPermissionList = append(oPermissionList, permission)
			}
		}
	}
	obj.PermissionList = oPermissionList
}
