package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/service"
	"mimi/djq/util"
	"net/http"
	"strings"
	"mimi/djq/constant"
)

func RoleList(c *gin.Context) {
	argObj := &arg.Role{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	argObj.IdsNotIn = []string{constant.AdminRoleId}
	argObj.OrderBy = "name"
	c.JSON(http.StatusOK, service.ResultList(&service.Role{}, argObj))
}

func RoleGet(c *gin.Context) {
	serviceObj := &service.Role{}
	var result *util.ResultVO
	obj, err := serviceObj.Get(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result = util.BuildSuccessResult(obj)
	c.JSON(http.StatusOK, result)
}

func RolePost(c *gin.Context) {
	obj := &model.Role{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	serviceObj := &service.Role{}
	obj.BindStr2PermissionList()
	obj, err = serviceObj.Add(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(obj)
	c.JSON(http.StatusOK, result)
}

func RolePatch(c *gin.Context) {
	obj := &model.Role{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	serviceObj := &service.Role{}
	if obj.Id == constant.AdminRoleId {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("禁止修改超级管理员角色信息"))
		return
	}
	obj.BindStr2PermissionList()
	obj, err = serviceObj.Update(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(obj)
	c.JSON(http.StatusOK, result)
}

func RoleDelete(c *gin.Context) {
	ids := strings.Split(c.Query("ids"), constant.Split4Id)
	for _, v := range ids {
		if v == constant.AdminRoleId {
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("禁止删除超级管理员信息"))
			return
		}
	}

	serviceObj := &service.Role{}
	count, err := serviceObj.Delete(ids...)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(count)
	c.JSON(http.StatusOK, result)
}
