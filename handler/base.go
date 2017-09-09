package handler

import (
	"github.com/pkg/errors"
	"mimi/djq/service"
	"sync"
)

var ErrParamException = errors.New("参数异常")

var once sync.Once
var roleService *service.Role
var adminService *service.Admin

func GetRoleServiceInstance() *service.Role {
	//once.Do(func() {
	//	roleService = &service.Role{}
	//})
	//if roleService !=nil{
	//	roleService = &service.Role{}
	//}
	//return roleService
	return &service.Role{}
}
