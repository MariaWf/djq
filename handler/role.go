package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/service"
	"mimi/djq/util"
	"log"
	"mimi/djq/dao"
)

func RoleList(c *gin.Context) {
	argObj := &arg.Role{}
	err := c.Bind(argObj)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, util.BuildFailResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.ResultList(GetRoleServiceInstance(), argObj))
}

func RoleGet(c *gin.Context) {
	service := GetRoleServiceInstance()
	var result *util.ResultVO
	obj, err := service.Get(c.Param("id"))
	if err != nil {
		log.Println(err)
		if err == dao.ErrObjectEmpty || err == dao.ErrIdEmpty {
			result = util.BuildFailResult(err.Error())
		} else {
			result = util.BuildFailResult("操作失败，请稍后重试")
		}
	} else {
		result = util.BuildSuccessResult(obj)
	}
	c.JSON(http.StatusOK, result)
}

func RolePost(c *gin.Context) {
	obj := &model.Role{}
	err := c.Bind(obj)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, util.BuildFailResult(err.Error()))
		return
	}
	service := GetRoleServiceInstance()
	var result *util.ResultVO
	obj, err = service.Add(obj)
	if err != nil {
		log.Println(err)
		if err == dao.ErrObjectEmpty {
			result = util.BuildFailResult(err.Error())
		} else {
			result = util.BuildFailResult("操作失败，请稍后重试")
		}
	} else {
		result = util.BuildSuccessResult(obj)
	}
	c.JSON(http.StatusOK, result)
}

func RolePatch(c *gin.Context) {
	role := &model.Role{}
	err := c.Bind(role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, util.BuildFailResult(err.Error()))
	}
	c.JSON(http.StatusOK, service.ResultUpdate(GetRoleServiceInstance(), role, "name", "description"))
}

func RoleDelete(c *gin.Context) {
	args := &arg.Role{}
	err := c.Bind(args)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, util.BuildFailResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.ResultBatchDelete(GetRoleServiceInstance(), args))
}

func RoleAllPermission(c *gin.Context) {
	permissionList := model.GetPermissionList()
	c.JSON(http.StatusOK, util.BuildSuccessResult(permissionList))
}