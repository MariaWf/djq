package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"mimi/djq/config"
	"mimi/djq/util"
	"net/http"
	"strings"
	"mimi/djq/session"
	"log"
	"math/rand"
	"strconv"
	"mimi/djq/aliyun"
	"mimi/djq/cache"
)

func NotFound(c *gin.Context) {
	if c.Request.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		c.AbortWithStatusJSON(http.StatusNotFound, util.BuildFailResult("未知资源"))
	}
}

func ApiGlobal(c *gin.Context) {
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "http://192.168.1.104:8080")
	//c.Writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:8080")
	c.MultipartForm()
	origins := strings.Split(config.Get("allowOrigin"), ",")
	if origins != nil && len(origins) != 0 {
		for _, origin := range origins {
			if c.Request.Header.Get("origin") == origin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
				c.Writer.Header().Set("Access-Control-Allow-Method", "POST,GET")
				c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")
				break
			}
		}
	}
	//c.Writer.Header().Set("Access-Control-Allow-Origin","*")
	//Access-Control-Allow-Origin: http://api.bob.com
	//Access-Control-Allow-Credentials: true
	//Access-Control-Expose-Headers: FooBar
	//c.Writer.Header().Add("Access-Control-Allow-Origin","http://localhost:8080,http://192.168.1.104:8080")
	//c.Writer.Header().Add("Access-Control-Allow-Credentials","true")
	//c.Writer.Header().Add("Access-Control-Expose-Headers","FooBar")
	//c.Request.Referer()
	//if !strings.Contains(c.Request.Host, "djq.51zxiu.cn") && !strings.Contains(c.Request.Host, "djq.tunnel.qydev.com") {
	//	c.AbortWithStatusJSON(http.StatusForbidden, util.BuildFailResult("禁止访问"))
	//}
}

//func Index(w http.ResponseWriter, r *http.Request) {
//	values := make(map[string]interface{})
//	if authentication, session, err := CheckLogin(w, r); err != nil {
//		log.Println(err)
//	} else {
//		values["authentication"] = authentication
//		if loginName, err := session.Get("loginName"); err != nil {
//			log.Println(err)
//		} else {
//			values["loginName"] = loginName
//		}
//	}
//	t, _ := template.ParseFiles("html/template/index.html", "html/template/head.html")
//	t.ExecuteTemplate(w, "head", values)
//	//t.Execute(w, values)
//}

func Index2(c *gin.Context) {
	values := make(map[string]interface{})
	t, _ := template.ParseFiles("html/template/index.html")
	t.Execute(c.Writer, values)
}

func TestUser(c *gin.Context) {
	values := make(map[string]interface{})
	t, _ := template.ParseFiles("html/template/user.html")
	t.Execute(c.Writer, values)
}

func TestShop(c *gin.Context) {
	values := make(map[string]interface{})
	t, _ := template.ParseFiles("html/template/shop.html")
	t.Execute(c.Writer, values)
}

func GetServerRootUrl(c *gin.Context){
	result := util.BuildSuccessResult(config.Get("server_root_url"))
	c.JSON(http.StatusOK, result)
}

func GetPublicKey(c *gin.Context) {
	result := util.BuildSuccessResult(string(util.GetPublicKey()))
	c.JSON(http.StatusOK, result)
}

func GeetestInit(c *gin.Context) {
	sn, err := session.GetOpen(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("验证码初始化异常"))
		return
	}
	geetestId := config.Get("geetest_id")
	geetestKey := config.Get("geetest_key")
	var userID = sn.Id
	gt := util.GeetestLib(geetestKey, geetestId)
	gt.PreProcess(userID)
	responseMap := gt.GetResponseMap()
	c.JSON(http.StatusOK, responseMap)
}

func GetCaptcha(c *gin.Context){
	if !util.GeetestCheck(c) {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(util.ErrParamException.Error()))
		return
	}
	mobile := c.PostForm("mobile")
	if !util.MatchMobile(mobile){
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(util.ErrMobileFormat.Error()))
		return
	}
	sn ,err := session.GetUi(c.Writer,c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	captcha := strconv.Itoa(rand.Intn(8888)+1000)
	err = sn.Set(session.SessionNameUiUserCaptcha,captcha)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	err = aliyun.CaptchaSend(mobile, captcha)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	result := util.BuildSuccessResult(mobile)
	//result := util.BuildSuccessResult(captcha)
	c.JSON(http.StatusOK, result)
}

func GetTotalCashCouponPrice(c *gin.Context){
	price,err := cache.Get(cache.CacheNameGlobalTotalCashCouponPrice)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	result := util.BuildSuccessResult(price)
	c.JSON(http.StatusOK, result)
}

func Test(c *gin.Context) {
	//_, _ = c.FormFile("theFile")
	//c.Request.MultipartReader()
	c.MultipartForm()
	fmt.Println(c.ContentType())
	//c.MultipartForm()

	c.AbortWithStatusJSON(http.StatusOK, util.BuildNeedLoginResult())
}

func TestGet(c *gin.Context) {
	result := util.BuildSuccessResult("mimi")
	c.JSON(http.StatusOK, result)
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
