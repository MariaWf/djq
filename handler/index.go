package handler

import (
	"net/http"
	"html/template"
	"log"
	"github.com/gin-gonic/gin"
	"fmt"
	"mimi/djq/util"
)

func NotFound(c *gin.Context) {
	if c.Request.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		c.AbortWithStatusJSON(http.StatusNotFound, util.BuildFailResult("未知资源"))
	}
}

func ApiGlobal(c *gin.Context) {
	//if !strings.Contains(c.Request.Host, "djq.51zxiu.cn") && !strings.Contains(c.Request.Host, "djq.tunnel.qydev.com") {
	//	c.AbortWithStatusJSON(http.StatusForbidden, util.BuildFailResult("禁止访问"))
	//}
}

func Index(w http.ResponseWriter, r *http.Request) {
	values := make(map[string]interface{})
	if authentication, session, err := CheckLogin(w, r); err != nil {
		log.Println(err)
	} else {
		values["authentication"] = authentication
		if loginName, err := session.Get("loginName"); err != nil {
			log.Println(err)
		} else {
			values["loginName"] = loginName
		}
	}
	t, _ := template.ParseFiles("html/template/index.html", "html/template/head.html")
	t.ExecuteTemplate(w, "head", values)
	//t.Execute(w, values)
}

func Index2(c *gin.Context) {
	values := make(map[string]interface{})
	//if authentication, session, err := CheckLogin(w, r); err != nil {
	//	log.Println(err)
	//} else {
	//	values["authentication"] = authentication
	//	if loginName, err := session.Get("loginName"); err != nil {
	//		log.Println(err)
	//	} else {
	//		values["loginName"] = loginName
	//	}
	//}
	//t, _ := template.ParseFiles("html/template/index.html", "html/template/head.html")
	//t.ExecuteTemplate(c.Writer, "head", values)

	t, _ := template.ParseFiles("html/template/index.html")
	t.Execute(c.Writer, values)
}

func Upload(c *gin.Context) {
	file, _ := c.FormFile("theFile")
	fmt.Println(file.Filename + "_" + c.PostForm("wawaName"))

	// Upload the file to specific dst.
	// c.SaveUploadedFile(file, dst)
	c.SaveUploadedFile(file, "c:/upload/" + file.Filename)

	//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	result := util.BuildSuccessResult(file.Filename)
	c.JSON(http.StatusOK, result)
}

func Api(c *gin.Context) {
	//firstname := c.DefaultQuery("firstname", "Guest")
	//lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	//
	//c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	str := `规则：
		GET /zoos：列出所有动物园
		POST /zoos：新建一个动物园
		GET /zoos/ID：获取某个指定动物园的信息
		PUT /zoos/ID：更新某个指定动物园的信息（提供该动物园的全部信息）
		PATCH /zoos/ID：更新某个指定动物园的信息（提供该动物园的部分信息）
		DELETE /zoos/ID：删除某个动物园
		GET /zoos/ID/animals：列出某个指定动物园的所有动物
		DELETE /zoos/ID/animals/ID：删除某个指定动物园的指定动物

		路径以“/html/”开头的，返回页面，其他返回json数据
		Json数据统一格式{ “status”:1/0,”msg”:”xxxxxxx”,”result”:{}};
		查询操作返回数组datas、当前页curPage、每页数量pageSize、总页数totalPage、总数total
		详情操作返回对象data
		添加操作返回id
		修改操作返回id
		删除操作返回id
		上传操作返回url
		批处理操作返回数组ids，行为action

		`
	c.Writer.WriteString(str)
}