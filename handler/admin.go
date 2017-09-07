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
)

func AdminLogin(c *gin.Context) {
}

func AdminLogout(c *gin.Context) {

}

func AdminCheckLogin(c *gin.Context) {
	if c.DefaultQuery("name", "") != "mimi" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, util.BuildNeedLoginResult())
	}
}

func AdminList(c *gin.Context) {
	argObj := &arg.Admin{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, util.BuildFailResult(ErrParamException.Error()))
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, util.BuildFailResult(err.Error()))
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	log.Println(admin)
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(obj.Id)
	log.Println(obj)
	c.JSON(http.StatusOK, result)
}

func AdminPatch(c *gin.Context) {
	//admin := &model.Admin{}
	//err := c.Bind(admin)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, util.BuildFailResult(err.Error()))
	//}
	//conn, _ := mysql.Get()
	//defer mysql.Close(conn)
	//adminDao := &dao.AdminDao{conn}
	//admin, err = adminDao.Update(admin)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, util.BuildFailResult(err.Error()))
	//}
	//c.JSON(http.StatusOK, util.BuildSuccessResult(admin))
}

func AdminDelete(c *gin.Context) {
	//args := &arg.Admin{}
	//err := c.Bind(args)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, util.BuildFailResult(err.Error()))
	//	return
	//}
	//conn, _ := mysql.Get()
	//defer mysql.Close(conn)
	//adminDao := &dao.AdminDao{conn}
	//count, err := adminDao.Delete(args)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, util.BuildFailResult(err.Error()))
	//}
	//c.JSON(http.StatusOK, util.BuildSuccessResult(count))
}
