package router

import (
	"github.com/gin-gonic/gin"
	"io"
	"mimi/djq/config"
	"mimi/djq/handler"
	"os"
	"path/filepath"
)

func Begin() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()
	initLog()
	router := gin.Default()
	router.NoRoute(handler.NotFound)
	ui := router.Group("/ui", handler.ApiGlobal)
	ui.POST("/login", handler.AdminList)

	mi := router.Group("/mi", handler.ApiGlobal, handler.AdminCheckLogin)

	mi.POST("/login", handler.AdminLogin)
	mi.POST("/logout", handler.AdminLogout)

	miAdmin := mi.Group("/admin")
	miAdmin.GET("/", handler.PermissionAdminR, handler.AdminList)
	miAdmin.GET("/:id", handler.PermissionAdminR, handler.AdminGet)
	miAdmin.POST("/", handler.PermissionAdminC, handler.AdminPost)
	miAdmin.PATCH("/:id", handler.PermissionAdminU, handler.AdminPatch)
	miAdmin.DELETE("/:id", handler.PermissionAdminD, handler.AdminDelete)

	miRole := mi.Group("/role")
	miRole.GET("/", handler.PermissionRoleR, handler.RoleList)
	miRole.GET("/:id", handler.PermissionRoleR, handler.RoleGet)
	miRole.POST("/", handler.PermissionRoleC, handler.RolePost)
	miRole.PATCH("/:id", handler.PermissionRoleU, handler.RolePatch)
	miRole.DELETE("/:id", handler.PermissionRoleD, handler.RoleDelete)

	miPermission := mi.Group("/permission")
	miPermission.GET("/", handler.PermissionRoleR, handler.PermissionList)

	si := router.Group("/si", handler.ApiGlobal)
	si.POST("/login", handler.AdminList)

	open := router.Group("/open", handler.ApiGlobal)

	open.POST("/geetest", handler.GeetestInit)
	open.GET("/getPublicKey", handler.GetPublicKey)

	open.GET("/shop", handler.AdminList)
	open.GET("/shop/:id", handler.AdminGet)
	open.GET("/advertisement", handler.AdminGet)
	open.GET("/shopClassification", handler.AdminGet)

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	//router.Static("/", "html/1.0.0")
	//router.GET("/welcome", func(c *gin.Context) {
	//	firstname := c.DefaultQuery("firstname", "Guest")
	//	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	//
	//	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	//})
	router.Static("/static", "html/assets")
	router.GET("/", handler.Index2)
	router.GET("/api", handler.Api)
	//router.GET("/getPublicKey", handler.ApiGlobal,handler.GetPublicKey)
	//router.GET("/user/:id", handler.UserGet)
	//router.GET("/user", handler.UserList)
	//router.POST("/upload", handler.Upload)
	//
	//router.GET("/html/ui", handler.Index2)
	//router.GET("/html/mi", handler.Index2)
	//router.GET("/advertisementList4Index", handler.AdvertisementList4Index)
	//router.GET("/shopClassificationList4Index", handler.ShopClassificationList4Index)
	//router.GET("/shop4Index", handler.Shop4IndexList)
	//router.GET("/shop4Index/:id", handler.Shop4Index)
	//
	router.StaticFile("/geetest", "html/template/geetest.html")
	//router.StaticFile("/login", "html/template/login.html")
	//router.GET("/geetest", handler.GeetestRegister)
	router.GET("/register", handler.GeetestRegister)
	router.POST("/validate", handler.GeetestValidate)
	router.POST("/ajax_validate", handler.GeetestAjaxValidate)
	//
	//ui.GET("/", handler.Index2)
	////mi.GET("/", handler.Index2)
	//si.GET("/", handler.Index2)
	//open.GET("/", handler.Index2)

	router.Run(":" + config.Get("server_port"))
}

func initLog() {
	if "false" == config.Get("output_log") {
		return
	}
	gin.SetMode(gin.ReleaseMode)
	// Logging to a file.
	//gin_file, _ := os.Create("gin.log")
	routerInfoLogUrl := config.Get("router_info_log")
	if routerInfoLogUrl == "" {
		routerInfoLogUrl = "router_info.log"

	} else {
		path := filepath.Dir(routerInfoLogUrl)
		os.MkdirAll(path, 0777)
	}
	routerErrorLogUrl := config.Get("router_error_log")
	if routerErrorLogUrl == "" {
		routerErrorLogUrl = "router_error.log"
	} else {
		path := filepath.Dir(routerErrorLogUrl)
		os.MkdirAll(path, 0777)
	}
	infoFile, err := os.OpenFile(routerInfoLogUrl, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	errorFile, err2 := os.OpenFile(routerErrorLogUrl, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err2 != nil {
		panic(err2)
	}
	//ginLogUrl := "gin.log"
	//var gin_file *os.File
	//if _,err := os.Stat(ginLogUrl);err!=nil{
	//	gin_file, _ = os.Create(ginLogUrl)
	//}else{
	//	gin_file, _ = os.OpenFile("gin.log", os.O_RDWR|os.O_APPEND, 0666)
	//}
	//gin_error_file, _ := os.Create(routerErrorLogUrl)
	gin.DefaultWriter = io.MultiWriter(infoFile)
	gin.DefaultErrorWriter = errorFile
	//gin.RecoveryWithWriter(gin_error_file)
}
