package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mimi/djq/util"
	"mimi/djq/config"
	"net/http"
)

//func Register(w http.ResponseWriter, r *http.Request) {
//	if r.Method == "GET" {
//		if authentication, _, _ := CheckLogin(w, r); authentication {
//			http.Redirect(w, r, "/", http.StatusFound)
//			return
//		}
//		t, _ := template.ParseFiles("html/template/register.html")
//		t.Execute(w, nil)
//	} else {
//		values := make(map[string]interface{})
//		if err := RegisterAction(r); err == nil {
//			http.Redirect(w, r, "/", http.StatusFound)
//			return
//		} else {
//			values["error"] = err.Error()
//		}
//		t, _ := template.ParseFiles("html/template/register.html")
//		t.Execute(w, values)
//	}
//}

func RegisterAction(r *http.Request) error {
	return nil
}

func GeetestRegister(c *gin.Context) {
	geetestId := config.Get("geetest_id")
	geetestKey := config.Get("geetest_key")
	var userID = "test"
	//var userID = "testmimiwawa"
	gt := util.GeetestLib(geetestKey, geetestId)
	gt.PreProcess(userID)
	responseMap := gt.GetResponseMap()
	c.JSON(http.StatusOK, responseMap)
	//c.Data["json"]=responseMap
	//c.ServeJSON()
}

func GeetestValidate(c *gin.Context) {
	geetestId := config.Get("geetest_id")
	geetestKey := config.Get("geetest_key")
	var result bool
	var respstr string
	gt := util.GeetestLib(geetestKey, geetestId)
	challenge := c.GetString(util.FN_CHALLENGE)
	validate := c.GetString(util.FN_VALIDATE)
	seccode := c.GetString(util.FN_SECCODE)
	status := c.GetInt(util.GT_STATUS_SESSION_KEY)
	userID := c.GetString("user_id")
	userID = "test2"
	if status == 0 {
		fmt.Println("local")
		result = gt.FailbackValidate(challenge, validate, seccode)
	} else {
		fmt.Println("web")
		result = gt.SuccessValidate(challenge, validate, seccode, userID)
	}
	if result {
		respstr = "<html><body><h1>登录成功</h1></body></html>"
	} else {
		respstr = "<html><body><h1>登录失败</h1></body></html>"
	}
	fmt.Println(status)
	c.Writer.WriteString(respstr)
	//c.Ctx.WriteString(respstr)
}

func GeetestAjaxValidate(c *gin.Context) {
	//geetestId := config.Get("geetest_id")
	//geetestKey := config.Get("geetest_key")
	//var result bool
	var jsondata = make(map[string]string)
	//gt := util.GeetestLib(geetestKey, geetestId)
	//
	//challenge := c.PostForm(util.FN_CHALLENGE)
	//validate := c.PostForm(util.FN_VALIDATE)
	//seccode := c.PostForm(util.FN_SECCODE)
	//status := c.PostForm(util.GT_STATUS_SESSION_KEY)
	//userID := c.PostForm("user_id")
	//if status == "0" {
	//	result = gt.FailbackValidate(challenge, validate, seccode)
	//} else {
	//	result = gt.SuccessValidate(challenge, validate, seccode, userID)
	//}
	result := util.GeetestCheck(c)
	if result {
		jsondata["status"] = "success"
	} else {
		jsondata["status"] = "fail"
	}
	c.JSON(http.StatusOK, jsondata)
	//c.Data["json"]= jsondata
	//c.ServeJSON()
}
