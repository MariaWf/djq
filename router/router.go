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
	mi.POST("/advertisementAction/uploadImage", handler.PermissionAdvertisementCU, handler.AdvertisementUploadImage)

	miShop := mi.Group("/shop")
	miShop.GET("/", handler.PermissionShopR, handler.ShopList)
	miShop.GET("/:id", handler.PermissionShopR, handler.ShopGet)
	miShop.POST("/", handler.PermissionShopC, handler.ShopPost)
	miShop.PATCH("/:id", handler.PermissionShopU, handler.ShopPatch)
	miShop.DELETE("/", handler.PermissionShopD, handler.ShopDelete)
	mi.POST("/shopAction/uploadPreImage", handler.PermissionShopCU, handler.ShopUploadPreImage)
	mi.POST("/shopAction/uploadLogo", handler.PermissionShopCU, handler.ShopUploadLogo)

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
	mi.POST("/cashCouponAction/uploadImage", handler.PermissionShopCU, handler.CashCouponUploadImage)

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
	mi.POST("/presentAction/uploadImage", handler.PermissionShopCU, handler.PresentUploadImage)

	miPresentOrder := mi.Group("/presentOrder")
	miPresentOrder.GET("/", handler.PermissionPresentOrderR, handler.PresentOrderList)
	miPresentOrder.GET("/:id", handler.PermissionPresentOrderR, handler.PresentOrderGet)
	miPresentOrder.POST("/", handler.PermissionPresentOrderC, handler.PresentOrderPost)
	miPresentOrder.PATCH("/:id", handler.PermissionPresentOrderU, handler.PresentOrderPatch)
	miPresentOrder.DELETE("/", handler.PermissionPresentOrderD, handler.PresentOrderDelete)

	miCashCouponOrder := mi.Group("/cashCouponOrder")
	miCashCouponOrder.GET("/", handler.PermissionCashCouponOrderR, handler.CashCouponOrderList)
	miCashCouponOrder.GET("/:id", handler.PermissionCashCouponOrderR, handler.CashCouponOrderGet)
	miCashCouponOrder.POST("/", handler.PermissionCashCouponOrderC, handler.CashCouponOrderPost)
	miCashCouponOrder.PATCH("/:id", handler.PermissionCashCouponOrderU, handler.CashCouponOrderPatch)
	miCashCouponOrder.DELETE("/", handler.PermissionCashCouponOrderD, handler.CashCouponOrderDelete)

	miRefund := mi.Group("/refund")
	miRefund.GET("/", handler.PermissionRefundR, handler.RefundList)
	miRefund.GET("/:id", handler.PermissionRefundR, handler.RefundGet)
	miRefund.POST("/", handler.PermissionRefundC, handler.RefundPost)
	miRefund.PATCH("/:id", handler.PermissionRefundU, handler.RefundPatch)
	miRefund.DELETE("/", handler.PermissionRefundD, handler.RefundDelete)
	miRefund.POST("/refundAction/uploadEvidence",handler.PermissionRefundCU, handler.RefundUploadEvidence)

	miRefundReason := mi.Group("/refundReason")
	miRefundReason.GET("/", handler.PermissionRefundReasonR, handler.RefundReasonList)
	miRefundReason.GET("/:id", handler.PermissionRefundReasonR, handler.RefundReasonGet)
	miRefundReason.POST("/", handler.PermissionRefundReasonC, handler.RefundReasonPost)
	miRefundReason.PATCH("/:id", handler.PermissionRefundReasonU, handler.RefundReasonPatch)
	miRefundReason.DELETE("/", handler.PermissionRefundReasonD, handler.RefundReasonDelete)

	si := router.Group("/si", handler.ApiGlobal, handler.ShopAccountCheckLogin)
	si.POST("/login", handler.ShopAccountLogin)
	si.POST("/logout", handler.ShopAccountLogout)

	si.GET("/shopAccountAction/self", handler.ShopAccountGetSelf)

	si.GET("/presentOrderOrCashCouponOrder", handler.ShopAccountActionGetPresentOrderOrCashCouponOrder4Si)
	si.POST("/completePresentOrder", handler.PresentOrderComplete4Si)
	si.POST("/completeCashCouponOrder", handler.CashCouponOrderComplete4Si)
	si.POST("/getMoney", handler.ShopAccountGetMoney4Si)


	ui := router.Group("/ui", handler.ApiGlobal, handler.UserCheckLogin)
	ui.POST("/login", handler.UserLogin)
	ui.POST("/logout", handler.UserLogout)

	ui.GET("/cashCouponOrderInCart",handler.CashCouponOrderListInCart4Ui)
	ui.GET("/cashCouponOrderUnused",handler.CashCouponOrderListUnused4Ui)
	ui.GET("/cashCouponOrderUsed",handler.CashCouponOrderListUsed4Ui)

	ui.DELETE("/cashCouponOrder",handler.CashCouponOrderDelete4Ui)
	ui.POST("/cashCouponOrder",handler.CashCouponOrderPost4Ui)
	ui.POST("/cashCouponOrderAction/buyFromCashCoupon",handler.CashCouponOrderActionBuyFromCashCoupon4Ui)
	ui.POST("/cashCouponOrderAction/buyFromCashCouponOrder", handler.CashCouponOrderActionBuyFromCashCouponOrder4Ui)


	ui.GET("/refund",handler.RefundList4Ui)

	ui.GET("/refundReason",handler.RefundReasonList4Ui)

	ui.POST("/refund",handler.RefundPost4Ui)
	ui.POST("/refundAction/cancel",handler.RefundCancel4Ui)
	ui.POST("/refundAction/uploadEvidence", handler.RefundUploadEvidence)

	ui.GET("/present",handler.PresentList4Ui)

	ui.GET("/presentOrder",handler.PresentOrderList4Ui)
	ui.POST("/presentOrder",handler.PresentOrderPost4Ui)

	open := router.Group("/open", handler.ApiGlobal)


	open.GET("/getServerRootUrl", handler.GetServerRootUrl)

	open.POST("/geetest", handler.GeetestInit)
	open.POST("/captcha", handler.GetCaptcha)

	open.GET("/getPublicKey", handler.GetPublicKey)

	open.GET("/shop", handler.ShopList4Open)
	open.GET("/shop/:id", handler.ShopGet4Open)
	open.GET("/shopClassification", handler.ShopClassificationList4Open)

	open.GET("/advertisement", handler.AdvertisementList4Open)


	wxpay := open.Group("/wxpay")
	wxpay.GET("config",handler.WxpayConfig)
	wxpay.GET("query",handler.WxpayQuery)
	wxpay.GET("getOpenId",handler.WxpayGetOpenId)

	wxpay.POST("notify4UnifiedOrder",handler.WxpayNotifyUnifiedOrder)
	wxpay.POST("notify4refund",handler.WxpayNotifyRefund)

	wxpay.GET("downloadBill",handler.WxpayDownloadBill)

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
	router.Static("/html", "html")

	uploadDirectoryHead := config.Get("uploadPath")
	if uploadDirectoryHead == "" {
		uploadDirectoryHead = "upload"
	}
	router.Static("/upload", uploadDirectoryHead)

	router.GET("/", handler.Index2)
	router.GET("/testUser", handler.TestUser)
	router.GET("/testShop", handler.TestShop)
	router.GET("/api", handler.Api)
	wxpay.GET("/index", handler.Wxpay)
	router.StaticFile("/MP_verify_IrIxUvnx9Bob0ktY.txt", "/home/zhaohao/app/MP_verify_IrIxUvnx9Bob0ktY.txt")
	//router.GET("/getPublicKey", handler.ApiGlobal,handler.GetPublicKey)
	//router.GET("/user/:id", handler.UserGet)
	//router.GET("/user", handler.UserList)
	//router.POST("/upload", handler.Upload)
	//mi.POST("/advertisementAction/uploadImage", handler.PermissionAdvertisementU, handler.AdvertisementUploadImage)
	//router.POST("/upload",  handler.PermissionShopU,handler.CashCouponUploadImage)
	router.POST("/upload1", handler.ApiGlobal, handler.AdminCheckLogin, handler.PermissionShopU, handler.CashCouponUploadImage)
	router.POST("/upload2", handler.AdminCheckLogin, handler.PermissionShopU, handler.CashCouponUploadImage)
	router.POST("/upload3", handler.PermissionShopU, handler.CashCouponUploadImage)
	router.POST("/upload4", handler.CashCouponUploadImage)
	router.POST("/upload", handler.Test, handler.Upload)
	router.POST("/test", handler.Test, handler.TestGet)
	//router.POST("/upload", handler.CashCouponUploadImage)
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
