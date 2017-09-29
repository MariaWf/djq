package initialization

import (
	"mimi/djq/config"
	"mimi/djq/constant"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/service"
	"mimi/djq/util"
	"strings"
)

func InitData() string {
	id := InitAdmin()
	InitSystemAdmin()
	InitBusinessAdmin()
	InitSimpleBusinessAdmin()
	InitReadAndDeleteAdmin()
	InitReadOnlyAdmin()
	InitReadAndAddAdmin()
	InitReadAndUpdateAdmin()
	return id
}

func InitRole(name, desc string, permissions []*model.Permission) string {
	serviceRole := &service.Role{}
	argRole := &arg.Role{}
	argRole.NameEqual = name
	roleList, err := service.Find(serviceRole, argRole)
	var role *model.Role
	checkErr(err)
	if roleList == nil || len(roleList) == 0 {
		role = &model.Role{}
		role.PermissionList = permissions
		role.Name = name
		role.Description = desc
		role, err := serviceRole.Add(role)
		checkErr(err)
		return role.Id
	} else {
		role = roleList[0].(*model.Role)
		role.BindStr2PermissionList()
		if len(role.PermissionList) != len(permissions) {
			role.PermissionList = permissions
			role, err = serviceRole.Update(role)
			checkErr(err)
		}
	}
	return role.Id
}

func InitAdminRole() string {
	name := config.Get("adminRole")
	if name == "" {
		name = "超级管理员"
	}
	desc := "超级管理员拥有最高权限，不能被删除，不能被修改"
	return InitRole(name, desc, model.GetPermissionList())
}

func InitSystemRole() string {
	name := "系统管理员"
	desc := "拥有删除行为以外的一切权限"
	permissions := make([]*model.Permission, 0, 10)
	for _, v := range model.GetPermissionList() {
		if strings.HasSuffix(v.Code, "_d") {
			continue
		}
		permissions = append(permissions, v)
	}
	return InitRole(name, desc, permissions)
}

func InitBusinessRole() string {
	name := "业务管理员"
	desc := "不包含角色管理、管理员管理权限，拥有对其他内容删除行为以外的一切权限"
	permissions := make([]*model.Permission, 0, 10)
	for _, v := range model.GetPermissionList() {
		if strings.HasSuffix(v.Code, "_d") || strings.HasPrefix(v.Code, "admin_") || strings.HasPrefix(v.Code, "role_") {
			continue
		}
		permissions = append(permissions, v)
	}
	return InitRole(name, desc, permissions)
}

func InitSimpleBusinessRole() string {
	name := "简单业务管理员"
	desc := "拥有对广告、退款原因、运维等内容删除行为以外的一切权限"
	permissions := make([]*model.Permission, 0, 10)
	for _, v := range model.GetPermissionList() {
		if strings.HasSuffix(v.Code, "_manage") || strings.HasPrefix(v.Code, "advertisement_") || strings.HasPrefix(v.Code, "refundReason_") {
			permissions = append(permissions, v)
		}
	}
	return InitRole(name, desc, permissions)
}

func InitReadOnlyRole() string {
	name := "阅读管理员"
	desc := "一切内容的阅读权限"
	permissions := make([]*model.Permission, 0, 10)
	for _, v := range model.GetPermissionList() {
		if strings.HasSuffix(v.Code, "_r") {
			permissions = append(permissions, v)
		}
	}
	return InitRole(name, desc, permissions)
}

func InitReadAndAddRole() string {
	name := "阅读添加管理员"
	desc := "一切内容的阅读、添加权限"
	permissions := make([]*model.Permission, 0, 10)
	for _, v := range model.GetPermissionList() {
		if strings.HasSuffix(v.Code, "_c") || strings.HasSuffix(v.Code, "_r") {
			permissions = append(permissions, v)
		}
	}
	return InitRole(name, desc, permissions)
}

func InitReadAndUpdateRole() string {
	name := "阅读更新管理员"
	desc := "一切内容的阅读、更新权限"
	permissions := make([]*model.Permission, 0, 10)
	for _, v := range model.GetPermissionList() {
		if strings.HasSuffix(v.Code, "_u") || strings.HasSuffix(v.Code, "_r") {
			permissions = append(permissions, v)
		}
	}
	return InitRole(name, desc, permissions)
}

func InitReadAndDeleteRole() string {
	name := "阅读删除管理员"
	desc := "一切内容的阅读、删除权限"
	permissions := make([]*model.Permission, 0, 10)
	for _, v := range model.GetPermissionList() {
		if strings.HasSuffix(v.Code, "_d") || strings.HasSuffix(v.Code, "_r") {
			permissions = append(permissions, v)
		}
	}
	return InitRole(name, desc, permissions)
}

func InitAdmin() string {
	roleId := InitAdminRole()
	name := config.Get("adminName")
	password := config.Get("adminPassword")
	constant.AdminRoleId = roleId
	return InitCommonAdmin(name, password, roleId)
}

func InitSystemAdmin() string {
	roleId := InitSystemRole()
	name := "systemAdmin"
	password := "123456"
	return InitCommonAdmin(name, password, roleId)
}

func InitBusinessAdmin() string {
	roleId := InitBusinessRole()
	name := "businessAdmin"
	password := "123456"
	return InitCommonAdmin(name, password, roleId)
}

func InitSimpleBusinessAdmin() string {
	roleId := InitSimpleBusinessRole()
	name := "simpleBusinessAdmin"
	password := "123456"
	return InitCommonAdmin(name, password, roleId)
}

func InitReadOnlyAdmin() string {
	roleId := InitReadOnlyRole()
	name := "readOnlyAdmin"
	password := "123456"
	return InitCommonAdmin(name, password, roleId)
}

func InitReadAndAddAdmin() string {
	roleId := InitReadAndAddRole()
	name := "readAndAddAdmin"
	password := "123456"
	return InitCommonAdmin(name, password, roleId)
}

func InitReadAndUpdateAdmin() string {
	roleId := InitReadAndUpdateRole()
	name := "readAndUpdateAdmin"
	password := "123456"
	return InitCommonAdmin(name, password, roleId)
}

func InitReadAndDeleteAdmin() string {
	roleId := InitReadAndDeleteRole()
	name := "readAndDeleteAdmin"
	password := "123456"
	return InitCommonAdmin(name, password, roleId)
}

func InitCommonAdmin(name, password, roleId string) string {
	serviceAdmin := &service.Admin{}
	argAdmin := &arg.Admin{}
	argAdmin.NameEqual = name
	objList, err := service.Find(serviceAdmin, argAdmin)
	checkErr(err)
	if objList == nil || len(objList) == 0 {
		obj := &model.Admin{}
		obj.Name = name
		obj.Password = password
		obj.Password, err = util.EncryptPassword(obj.Password)
		checkErr(err)
		obj.RoleList = []*model.Role{{Id: roleId}}
		obj, err := serviceAdmin.Add(obj)
		checkErr(err)
		return obj.Id
	}
	return objList[0].(*model.Admin).Id
}
