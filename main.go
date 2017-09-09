package main

import (
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
	"strconv"
	"math/rand"
)

func main() {
	initLog()
	log.Println("------LstdFlags：" + time.Now().String())
	log.SetFlags(log.LstdFlags | log.Llongfile)
	initData()
	if config.Get("buildTestData") == "true" {
		initTestData()
	}
	router.Begin()
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
	logFile, err := os.OpenFile(globalLogUrl, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.Println()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func initData() {
	constant.AdminId = initAdmin()
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

func initTestData() {
	initTestRole()
	initTestAdmin()
	initTestAdvertisement()
}

func initTestRole() {
	serviceRole := &service.Role{}
	argRole := &arg.Role{}
	count, err := service.Count(serviceRole, argRole)
	checkErr(err)
	if count < 5 {
		for i := 0; i < 50; i++ {
			obj := &model.Role{}
			pl := make([]*model.Permission, 0, 10)
			for _, v := range model.GetPermissionList() {
				if rand.Intn(2) < 1 {
					pl = append(pl, v)
				}
			}
			obj.PermissionList = pl
			obj.Name = "name" + strconv.Itoa(i)
			obj.Description = "description" + strconv.Itoa(i)
			obj, err := serviceRole.Add(obj)
			checkErr(err)
		}
	}
}

func initTestAdmin() {
	serviceAdmin := &service.Admin{}
	argAdmin := &arg.Admin{}
	count, err := service.Count(serviceAdmin, argAdmin)
	checkErr(err)
	if count < 5 {
		serviceRole := &service.Role{}
		argRole := &arg.Role{}
		roleList, err := service.Find(serviceRole, argRole)
		checkErr(err)
		for i := 0; i < 50; i++ {
			obj := &model.Admin{}
			rl := make([]*model.Role, 0, 10)
			for _, v := range roleList {
				if rand.Intn(10) < 2 {
					rl = append(rl, v.(*model.Role))
				}
			}
			obj.RoleList = rl
			obj.Name = "name" + strconv.Itoa(i)
			obj.Password = "123123"
			obj.Password, err = util.EncryptPassword(obj.Password)
			checkErr(err)
			obj, err := serviceAdmin.Add(obj)
			checkErr(err)
			checkErr(err)
		}
	}
}

func initTestAdvertisement() {
	serviceAdvertisement := &service.Advertisement{}
	argAdvertisement := &arg.Advertisement{}
	count, err := service.Count(serviceAdvertisement, argAdvertisement)
	checkErr(err)
	if count < 5 {
		for i := 0; i < 50; i++ {
			obj := &model.Advertisement{}
			obj.Name = "name" + strconv.Itoa(i)
			obj.Image = "https://www.baidu.com/img/bd_logo1.png"
			obj.Link = "https://www.baidu.com"
			obj.Priority = rand.Intn(1000)
			obj.Hide = rand.Intn(2) < 1
			obj.Description = "description" + strconv.Itoa(i)
			_, err := service.Add(serviceAdvertisement, obj)
			checkErr(err)
		}
	}
}