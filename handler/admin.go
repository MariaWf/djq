package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"mimi/djq/util"
	"mimi/djq/dao/arg"
	"log"
	"mimi/djq/service"
	"mimi/djq/model"
	"strings"
	"mimi/djq/session"
	"mimi/djq/constant"
)

func AdminLogin(c *gin.Context) {
	obj := &model.Admin{}
	err := c.Bind(obj)

	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	serviceObj := &service.Admin{}
	obj, err = serviceObj.CheckLogin(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}

	sn, err := session.GetMi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	if err = sn.Set("id", obj.Id); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	if err = sn.Set("name", obj.Name); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	if codeList := obj.GetPermissionCodeList(); codeList != nil&&len(codeList) != 0 {
		if err = sn.Set("permissionCode", strings.Join(codeList, ",")); err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
			return
		}
	}

	result := util.BuildSuccessResult(obj)
	c.JSON(http.StatusOK, result)
}

func AdminLogout(c *gin.Context) {
	sn, err := session.GetMi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	if err = sn.Del(); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(nil)
	c.JSON(http.StatusOK, result)
}

func AdminCheckLogin(c *gin.Context) {
	if c.Request.URL.Path != "/mi/login" && c.Request.URL.Path != "/mi/logout" {
		sn, err := session.GetMi(c.Writer, c.Request)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
			return
		}
		id, err := sn.Get("id")
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
			return
		}
		if id == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, util.BuildNeedLoginResult())
		}
	}
}

func AdminList(c *gin.Context) {
	argObj := &arg.Admin{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	serviceObj := &service.Admin{}
	argObj.ShowColumnNames = []string{"id", "name", "mobile", "locked"}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func AdminGet(c *gin.Context) {
	argObj := &arg.Admin{}
	argObj.SetIdEqual(c.Param("id"))
	serviceObj := &service.Admin{}
	obj, err := serviceObj.Get(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(obj)
	c.JSON(http.StatusOK, result)
}

func AdminPost(c *gin.Context) {
	admin := &model.Admin{}
	roleIds := c.PostForm("roleIds")
	err := c.Bind(admin)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	if roleIds != "" {
		roleIdList := strings.Split(roleIds, ",")
		if roleIdList != nil && len(roleIdList) != 0 {
			admin.RoleList = make([]*model.Role, len(roleIdList), len(roleIdList))
			for i, roleId := range roleIdList {
				admin.RoleList[i] = &model.Role{Id:roleId}
			}
		}
	}

	serviceObj := &service.Admin{}
	obj, err := serviceObj.Add(admin)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(obj.Id)
	c.JSON(http.StatusOK, result)
}

func AdminPatch(c *gin.Context) {
	admin := &model.Admin{}
	roleIds := c.PostForm("roleIds")
	err := c.Bind(admin)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	sn, err := session.GetMi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	id, err := sn.Get("id")
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	if id == admin.GetId() {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("禁止修改当前登录管理员信息"))
		return
	}
	if id == constant.AdminId {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("禁止修改超级管理员信息"))
		return
	}
	if roleIds != "" {
		roleIdList := strings.Split(roleIds, ",")
		if roleIdList != nil && len(roleIdList) != 0 {
			admin.RoleList = make([]*model.Role, len(roleIdList), len(roleIdList))
			for i, roleId := range roleIdList {
				admin.RoleList[i] = &model.Role{Id:roleId}
			}
		}
	}

	serviceObj := &service.Admin{}
	obj, err := serviceObj.Update(admin)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(obj.Id)
	c.JSON(http.StatusOK, result)
	//admin := &model.Admin{}
	//err := c.Bind(admin)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
	//}
	//conn, _ := mysql.Get()
	//defer mysql.Close(conn)
	//adminDao := &dao.AdminDao{conn}
	//admin, err = adminDao.Update(admin)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
	//}
	//c.JSON(http.StatusOK, util.BuildSuccessResult(admin))
}

func AdminDelete(c *gin.Context) {
	//args := &arg.Admin{}
	//err := c.Bind(args)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
	//	return
	//}
	//conn, _ := mysql.Get()
	//defer mysql.Close(conn)
	//adminDao := &dao.AdminDao{conn}
	//count, err := adminDao.Delete(args)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
	//}
	//c.JSON(http.StatusOK, util.BuildSuccessResult(count))
}
