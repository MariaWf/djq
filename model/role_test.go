package model

import "testing"

func TestRole_BindPermissionList2Str(t *testing.T) {
	str := "admin_r,role_d"
	obj := &Role{}
	obj.PermissionListStr = str
	obj.BindStr2PermissionList()
	obj.PermissionListStr = ""
	obj.BindPermissionList2Str()
	if str != obj.PermissionListStr {
		t.Error(str, obj.PermissionListStr)
	} else {
		t.Log(str)
	}
}

func TestRole_BindStr2PermissionList(t *testing.T) {
	obj := &Role{}
	obj.PermissionListStr = "admin_r,role_d"
	obj.BindStr2PermissionList()
	for _, pm := range obj.PermissionList {
		t.Log(pm)
	}
}
