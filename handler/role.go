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
	c.JSON(http.StatusOK, service.ResultList(GetRoleServiceInstance(), argObj))
}

func RoleGet(c *gin.Context) {
	serviceObj := GetRoleServiceInstance()
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
	serviceObj := GetRoleServiceInstance()
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
	serviceObj := GetRoleServiceInstance()
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
	serviceObj := GetRoleServiceInstance()
	count, err := serviceObj.Delete(strings.Split(c.PostForm("ids"), constant.Split4Id)...)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(count)
	c.JSON(http.StatusOK, result)
}
