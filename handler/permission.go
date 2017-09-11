package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"mimi/djq/constant"
	"mimi/djq/model"
	"mimi/djq/session"
	"mimi/djq/util"
	"net/http"
	"strings"
)

func PermissionList(c *gin.Context) {
	permissionList := model.GetPermissionList()
	c.JSON(http.StatusOK, util.BuildSuccessResult(permissionList))
}

func checkPermission(c *gin.Context, permission string) {
	sn, err := session.GetMi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	permissionStr, err := sn.Get(session.SessionNameMiPermission)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	permissionList := strings.Split(permissionStr, constant.Split4Permission)
	if permissionList != nil || len(permissionList) != 0 {
		for _, pn := range permissionList {
			if permission == pn {
				return
			}
		}
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, util.BuildNeedPermissionResult())
}

func PermissionAdminC(c *gin.Context) {
	checkPermission(c, "admin_c")
}

func PermissionAdminR(c *gin.Context) {
	checkPermission(c, "admin_r")
}

func PermissionAdminU(c *gin.Context) {
	checkPermission(c, "admin_u")
}

func PermissionAdminD(c *gin.Context) {
	checkPermission(c, "admin_d")
}

func PermissionRoleC(c *gin.Context) {
	checkPermission(c, "role_c")
}

func PermissionRoleR(c *gin.Context) {
	checkPermission(c, "role_r")
}

func PermissionRoleU(c *gin.Context) {
	checkPermission(c, "role_u")
}

func PermissionRoleD(c *gin.Context) {
	checkPermission(c, "role_d")
}

func PermissionAdvertisementC(c *gin.Context) {
	checkPermission(c, "advertisement_c")
}

func PermissionAdvertisementR(c *gin.Context) {
	checkPermission(c, "advertisement_r")
}

func PermissionAdvertisementU(c *gin.Context) {
	checkPermission(c, "advertisement_u")
}

func PermissionAdvertisementD(c *gin.Context) {
	checkPermission(c, "advertisement_d")
}

func PermissionShopC(c *gin.Context) {
	checkPermission(c, "shop_c")
}

func PermissionShopR(c *gin.Context) {
	checkPermission(c, "shop_r")
}

func PermissionShopU(c *gin.Context) {
	checkPermission(c, "shop_u")
}

func PermissionShopD(c *gin.Context) {
	checkPermission(c, "shop_d")
}

func PermissionShopClassificationC(c *gin.Context) {
	checkPermission(c, "shopClassification_c")
}

func PermissionShopClassificationR(c *gin.Context) {
	checkPermission(c, "shopClassification_r")
}

func PermissionShopClassificationU(c *gin.Context) {
	checkPermission(c, "shopClassification_u")
}

func PermissionShopClassificationD(c *gin.Context) {
	checkPermission(c, "shopClassification_d")
}

func PermissionUserC(c *gin.Context) {
	checkPermission(c, "user_c")
}

func PermissionUserR(c *gin.Context) {
	checkPermission(c, "user_r")
}

func PermissionUserU(c *gin.Context) {
	checkPermission(c, "user_u")
}

func PermissionUserD(c *gin.Context) {
	checkPermission(c, "user_d")
}

func PermissionPromotionalPartnerC(c *gin.Context) {
	checkPermission(c, "promotionalPartner_c")
}

func PermissionPromotionalPartnerR(c *gin.Context) {
	checkPermission(c, "promotionalPartner_r")
}

func PermissionPromotionalPartnerU(c *gin.Context) {
	checkPermission(c, "promotionalPartner_u")
}

func PermissionPromotionalPartnerD(c *gin.Context) {
	checkPermission(c, "promotionalPartner_d")
}

//func PermissionCashCouponC(c *gin.Context) {
//	checkPermission(c, "cashCoupon_c")
//}
//
//func PermissionCashCouponR(c *gin.Context) {
//	checkPermission(c, "cashCoupon_r")
//}
//
//func PermissionCashCouponU(c *gin.Context) {
//	checkPermission(c, "cashCoupon_u")
//}
//
//func PermissionCashCouponD(c *gin.Context) {
//	checkPermission(c, "cashCoupon_d")
//}


