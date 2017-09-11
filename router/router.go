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
	miAdmin.DELETE("/", handler.PermissionAdminD, handler.AdminDelete)

	mi.GET("/adminAction/self", handler.AdminGetSelf)
	mi.PATCH("/adminAction/self", handler.AdminPatchSelf)

	miRole := mi.Group("/role")
	miRole.GET("/", handler.PermissionRoleR, handler.RoleList)
	miRole.GET("/:id", handler.PermissionRoleR, handler.RoleGet)
	miRole.POST("/", handler.PermissionRoleC, handler.RolePost)
	miRole.PATCH("/:id", handler.PermissionRoleU, handler.RolePatch)
	miRole.DELETE("/", handler.PermissionRoleD, handler.RoleDelete)

	miPermission := mi.Group("/permission")
	miPermission.GET("/", handler.PermissionRoleR, handler.PermissionList)

	miAdvertisement := mi.Group("/advertisement")
	miAdvertisement.GET("/", handler.PermissionAdvertisementR, handler.AdvertisementList)
	miAdvertisement.GET("/:id", handler.PermissionAdvertisementR, handler.AdvertisementGet)
	miAdvertisement.POST("/", handler.PermissionAdvertisementC, handler.AdvertisementPost)
	miAdvertisement.PATCH("/:id", handler.PermissionAdvertisementU, handler.AdvertisementPatch)
	miAdvertisement.DELETE("/", handler.PermissionAdvertisementD, handler.AdvertisementDelete)
	mi.POST("/advertisementAction/uploadImage", handler.PermissionAdvertisementU, handler.AdvertisementUploadImage)

	miShop := mi.Group("/shop")
	miShop.GET("/", handler.PermissionShopR, handler.ShopList)
	miShop.GET("/:id", handler.PermissionShopR, handler.ShopGet)
	miShop.POST("/", handler.PermissionShopC, handler.ShopPost)
	miShop.PATCH("/:id", handler.PermissionShopU, handler.ShopPatch)
	miShop.DELETE("/", handler.PermissionShopD, handler.ShopDelete)
	mi.POST("/shopAction/uploadPreImage", handler.PermissionShopU, handler.ShopUploadPreImage)
	mi.POST("/shopAction/uploadLogo", handler.PermissionShopU, handler.ShopUploadLogo)

	miShopClassification := mi.Group("/shopClassification")
	miShopClassification.GET("/", handler.PermissionShopClassificationR, handler.ShopClassificationList)
	miShopClassification.GET("/:id", handler.PermissionShopClassificationR, handler.ShopClassificationGet)
	miShopClassification.POST("/", handler.PermissionShopClassificationC, handler.ShopClassificationPost)
	miShopClassification.PATCH("/:id", handler.PermissionShopClassificationU, handler.ShopClassificationPatch)
	miShopClassification.DELETE("/", handler.PermissionShopClassificationD, handler.ShopClassificationDelete)

	miShopAccount := mi.Group("/shopAccount")
	miShopAccount.GET("/", handler.PermissionShopR, handler.ShopAccountList)
	miShopAccount.GET("/:id", handler.PermissionShopR, handler.ShopAccountGet)
	miShopAccount.POST("/", handler.PermissionShopC, handler.ShopAccountPost)
	miShopAccount.PATCH("/:id", handler.PermissionShopU, handler.ShopAccountPatch)
	miShopAccount.DELETE("/", handler.PermissionShopD, handler.ShopAccountDelete)

	miShopIntroductionImage := mi.Group("/shopIntroductionImage")
	miShopIntroductionImage.GET("/", handler.PermissionShopR, handler.ShopIntroductionImageList)
	miShopIntroductionImage.GET("/:id", handler.PermissionShopR, handler.ShopIntroductionImageGet)
	miShopIntroductionImage.POST("/", handler.PermissionShopC, handler.ShopIntroductionImagePost)
	miShopIntroductionImage.PATCH("/:id", handler.PermissionShopU, handler.ShopIntroductionImagePatch)
	miShopIntroductionImage.DELETE("/", handler.PermissionShopD, handler.ShopIntroductionImageDelete)

	miCashCoupon := mi.Group("/cashCoupon")
	miCashCoupon.GET("/", handler.PermissionShopR, handler.CashCouponList)
	miCashCoupon.GET("/:id", handler.PermissionShopR, handler.CashCouponGet)
	miCashCoupon.POST("/", handler.PermissionShopC, handler.CashCouponPost)
	miCashCoupon.PATCH("/:id", handler.PermissionShopU, handler.CashCouponPatch)
	miCashCoupon.DELETE("/", handler.PermissionShopD, handler.CashCouponDelete)
	mi.POST("/cashCouponAction/uploadImage", handler.PermissionShopU, handler.CashCouponUploadImage)

	miPromotionalPartner := mi.Group("/promotionalPartner")
	miPromotionalPartner.GET("/", handler.PermissionPromotionalPartnerR, handler.PromotionalPartnerList)
	miPromotionalPartner.GET("/:id", handler.PermissionPromotionalPartnerR, handler.PromotionalPartnerGet)
	miPromotionalPartner.POST("/", handler.PermissionPromotionalPartnerC, handler.PromotionalPartnerPost)
	miPromotionalPartner.PATCH("/:id", handler.PermissionPromotionalPartnerU, handler.PromotionalPartnerPatch)
	miPromotionalPartner.DELETE("/", handler.PermissionPromotionalPartnerD, handler.PromotionalPartnerDelete)

	miUser := mi.Group("/user")
	miUser.GET("/", handler.PermissionUserR, handler.UserList)
	miUser.GET("/:id", handler.PermissionUserR, handler.UserGet)
	miUser.POST("/", handler.PermissionUserC, handler.UserPost)
	miUser.PATCH("/:id", handler.PermissionUserU, handler.UserPatch)
	miUser.DELETE("/", handler.PermissionUserD, handler.UserDelete)

	miPresent := mi.Group("/present")
	miPresent.GET("/", handler.PermissionPresentR, handler.PresentList)
	miPresent.GET("/:id", handler.PermissionPresentR, handler.PresentGet)
	miPresent.POST("/", handler.PermissionPresentC, handler.PresentPost)
	miPresent.PATCH("/:id", handler.PermissionPresentU, handler.PresentPatch)
	miPresent.DELETE("/", handler.PermissionPresentD, handler.PresentDelete)

	miPresentOrder := mi.Group("/presentOrder")
	miPresentOrder.GET("/", handler.PermissionPresentOrderR, handler.PresentOrderList)
	miPresentOrder.GET("/:id", handler.PermissionPresentOrderR, handler.PresentOrderGet)
	miPresentOrder.POST("/", handler.PermissionPresentOrderC, handler.PresentOrderPost)
	miPresentOrder.PATCH("/:id", handler.PermissionPresentOrderU, handler.PresentOrderPatch)
	miPresentOrder.DELETE("/", handler.PermissionPresentOrderD, handler.PresentOrderDelete)

	si := router.Group("/si", handler.ApiGlobal, handler.ShopAccountCheckLogin)
	si.POST("/login", handler.ShopAccountLogin)
	si.POST("/logout", handler.ShopAccountLogout)

	si.GET("/shopAccountAction/self", handler.ShopAccountGetSelf)

	open := router.Group("/open", handler.ApiGlobal)

	open.POST("/geetest", handler.GeetestInit)
	open.GET("/getPublicKey", handler.GetPublicKey)

	open.GET("/shop", handler.ShopList4Open)
	open.GET("/shop/:id", handler.ShopGet4Open)
	open.GET("/shopClassification", handler.ShopClassificationList4Open)

	open.GET("/advertisement", handler.AdvertisementList4Open)

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

	uploadDirectoryHead := config.Get("uploadPath")
	if uploadDirectoryHead == "" {
		uploadDirectoryHead = "upload"
	}
	router.Static("/upload", uploadDirectoryHead)

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
	//router.StaticFile("/geetest", "html/template/geetest.html")
	////router.StaticFile("/login", "html/template/login.html")
	////router.GET("/geetest", handler.GeetestRegister)
	//router.GET("/register", handler.GeetestRegister)
	//router.POST("/validate", handler.GeetestValidate)
	//router.POST("/ajax_validate", handler.GeetestAjaxValidate)
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
		routerInfoLogUrl = "logs/router_info.log"

	}
	path := filepath.Dir(routerInfoLogUrl)
	os.MkdirAll(path, 0777)

	routerErrorLogUrl := config.Get("router_error_log")
	if routerErrorLogUrl == "" {
		routerErrorLogUrl = "logs/router_error.log"
	}
	path = filepath.Dir(routerErrorLogUrl)
	os.MkdirAll(path, 0777)
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
