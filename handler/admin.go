package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"mimi/djq/constant"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/service"
	"mimi/djq/session"
	"mimi/djq/util"
	"net/http"
	"strings"
)

func AdminLogin(c *gin.Context) {
	if !util.GeetestCheck(c) {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(util.ErrParamException.Error()))
		return
	}

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
	if err = sn.Set(session.SessionNameMiAdminId, obj.Id); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	if err = sn.Set(session.SessionNameMiAdminName, obj.Name); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	if codeList := obj.GetPermissionCodeList(); codeList != nil && len(codeList) != 0 {
		if err = sn.Set(session.SessionNameMiPermission, strings.Join(codeList, constant.Split4Permission)); err != nil {
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
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
			return
		}
		id, err := sn.Get(session.SessionNameMiAdminId)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
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
				admin.RoleList[i] = &model.Role{Id: roleId}
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
	obj := &model.Admin{}
	roleIds := c.PostForm("roleIds")
	err := c.Bind(obj)
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
	if id == obj.GetId() {
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
			obj.RoleList = make([]*model.Role, len(roleIdList), len(roleIdList))
			for i, roleId := range roleIdList {
				obj.RoleList[i] = &model.Role{Id: roleId}
			}
		}
	}

	serviceObj := &service.Admin{}
	obj, err = serviceObj.Update(obj)
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
	ids := strings.Split(c.PostForm("ids"), constant.Split4Id)
	for _, v := range ids {
		if v == id {
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("禁止删除当前登录管理员信息"))
			return
		}
		if v == constant.AdminId {
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("禁止删除超级管理员信息"))
			return
		}
	}

	serviceObj := &service.Admin{}
	count, err := serviceObj.Delete(ids...)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(count)
	c.JSON(http.StatusOK, result)
}
