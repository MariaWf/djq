package main

import (
	"fmt"
	"log"
	"mimi/djq/config"
	"mimi/djq/constant"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/router"
	"mimi/djq/service"
	"mimi/djq/util"
	"os"
	"path/filepath"
	"time"
)

func main() {
	initLog()
	log.Println("------LstdFlags：" + time.Now().String())
	log.SetFlags(log.LstdFlags | log.Llongfile)
	initData()
	if config.Get("initTestData") == "true" {
		initTestData()
	}
	router.Begin()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func initData() {
	constant.AdminId = initAdmin()
}

func initTestData() {

}

func initRole() string {
	serviceRole := &service.Role{}
	argRole := &arg.Role{}
	argRole.NameEqual = config.Get("adminRole")
	roleList, err := service.Find(serviceRole, argRole)
	checkErr(err)
	if roleList == nil || len(roleList) == 0 {
		role := &model.Role{}
		role.PermissionList = model.GetPermissionList()
		fmt.Println(config.Get("adminRole"))
		role.Name = config.Get("adminRole")
		role.Description = "超级管理员不能删除，不能修改"
		role, err := serviceRole.Add(role)
		checkErr(err)
		return role.Id
	}
	return roleList[0].(*model.Role).Id
}

func initAdmin() string {
	roleId := initRole()
	constant.AdminRoleId = roleId
	serviceAdmin := &service.Admin{}
	argAdmin := &arg.Admin{}
	argAdmin.NameEqual = config.Get("adminName")
	objList, err := service.Find(serviceAdmin, argAdmin)
	checkErr(err)
	if objList == nil || len(objList) == 0 {
		obj := &model.Admin{}
		obj.Name = config.Get("adminName")
		obj.Password = config.Get("adminPassword")
		obj.Password, err = util.EncryptPassword(obj.Password)
		checkErr(err)
		obj.RoleList = []*model.Role{{Id: roleId}}
		obj, err := serviceAdmin.Add(obj)
		checkErr(err)
		return obj.Id
	}
	return objList[0].(*model.Admin).Id
}

func initLog() {
	if "false" == config.Get("output_log") {
		return
	}
	globalLogUrl := config.Get("global_log")
	if globalLogUrl == "" {
		globalLogUrl = "global.log"
	} else {
		path := filepath.Dir(globalLogUrl)
		os.MkdirAll(path, 0777)
	}
	logFile, err := os.OpenFile(globalLogUrl, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.Println()
}
