package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb/pkg/slices"
	"github.com/pkg/errors"
	"log"
	"mimi/djq/constant"
	"mimi/djq/dao/arg"
	"mimi/djq/model"
	"mimi/djq/service"
	"mimi/djq/session"
	"mimi/djq/util"
	"mimi/djq/wxpay"
	"net/http"
	"strconv"
	"strings"
)

func CashCouponOrderComplete4Si(c *gin.Context) {
	id := c.PostForm("id")
	sn, err := session.GetSi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	shopAccountId, err := sn.Get(session.SessionNameSiShopAccountId)
	if err != nil || shopAccountId == "" {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	serviceObj := &service.CashCouponOrder{}
	err = serviceObj.Complete(shopAccountId, id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult("")
	c.JSON(http.StatusOK, result)
}

func CashCouponOrderPost4Ui(c *gin.Context) {
	cashCouponIds := c.PostForm("ids")
	if strings.TrimSpace(cashCouponIds) == "" {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("未知代金券"))
		return
	}
	list, err := putInCartAction(c, strings.Split(cashCouponIds, constant.Split4Id)...)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(list)
	c.JSON(http.StatusOK, result)
}
func CashCouponOrderActionBuyFromCashCoupon4Ui(c *gin.Context) {
	cashCouponIds := c.PostForm("ids")
	if strings.TrimSpace(cashCouponIds) == "" {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("未知代金券"))
		return
	}
	list, err := putInCartAction(c, strings.Split(cashCouponIds, constant.Split4Id)...)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	cashCouponOrderIds := make([]string, len(list), len(list))
	for i, v := range list {
		cashCouponOrderIds[i] = v.Id
	}
	params, err := buyAction(c, cashCouponOrderIds...)
	if err != nil {
		if util.ErrNeedMiLogin.Error() == err.Error(){
			c.AbortWithStatusJSON(http.StatusOK,util.BuildNeedLoginResult())
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.BuildSuccessResult(params))
}

//func CashCouponOrderActionRefreshByPayOrderNumber4Ui(c *gin.Context) {
//	payOrderNumber := c.PostForm("payOrderNumber")
//	if payOrderNumber == "" {
//		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("未知订单编号"))
//		return
//	}
//	tradeState, totalFee, err := wxpay.OrderQueryResult(payOrderNumber)
//	if err != nil {
//		log.Println(err)
//		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
//		return
//	}
//	serviceObj := &service.CashCouponOrder{}
//	var idListStr string
//	switch tradeState {
//	case "SUCCESS":
//		idListStr, err = serviceObj.ConfirmOrder(payOrderNumber, totalFee)
//	default:
//		idListStr, err = serviceObj.CancelOrder(payOrderNumber)
//		cache.Set(cache.CacheNameWxpayPayOrderNumberCancel + payOrderNumber, idListStr, time.Hour * 24 * 7)
//	}
//}

func CashCouponOrderActionBuyFromCashCouponOrder4Ui(c *gin.Context) {
	cashCouponOrderIds := c.PostForm("ids")
	params, err := buyAction(c, strings.Split(cashCouponOrderIds, constant.Split4Id)...)
	if err != nil {
		if util.ErrNeedMiLogin.Error() == err.Error(){
			c.AbortWithStatusJSON(http.StatusOK,util.BuildNeedLoginResult())
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
	} else {
		c.JSON(http.StatusOK, util.BuildSuccessResult(params))
	}
}

func putInCartAction(c *gin.Context, ids ...string) (list []*model.CashCouponOrder, err error) {
	if len(ids) == 0 {
		err = errors.New("未知代金券")
		return
	}

	sn, err := session.GetUi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		err = errors.New(ErrUnknown.Error())
		return
	}
	userId, err := sn.Get(session.SessionNameUiUserId)
	if err != nil {
		log.Println(err)
		err = errors.New(ErrUnknown.Error())
		return
	}

	serviceObj := &service.CashCouponOrder{}
	list, err = serviceObj.BatchAddInCart(userId, ids...)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func buyAction(c *gin.Context, ids ...string) (params wxpay.Params, err error) {
	idList := make([]string, 0, len(ids))
	for _, v := range ids {
		if strings.TrimSpace(v) == "" {
			continue
		}
		idList = append(idList, strings.TrimSpace(v))
	}
	if len(idList) == 0 {
		err = errors.New("未包含代金券")
		return
	}
	sn, err := session.GetUi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		err = ErrUnknown
		return
	}
	userId, err := sn.Get(session.SessionNameUiUserId)
	if err != nil {
		log.Println(err)
		err = ErrUnknown
		return
	}
	serviceObj := &service.CashCouponOrder{}
	clientIp := c.ClientIP()
	openId, err := sn.Get(session.SessionNameUiUserOpenId)
	if err != nil {
		log.Println(err)
		err = ErrUnknown
		return
	}
	if openId == "" {
		sn2, err2 := session.GetSi(c.Writer, c.Request)
		if err2 != nil {
			log.Println(err2)
			err = ErrUnknown
			return
		}
		openId, err = sn2.Get(session.SessionNameSiShopAccountOpenId)
		if err != nil {
			log.Println(err)
			err = ErrUnknown
			return
		}
	}
	if openId == "" {
		//util.CookieCleanAll(c.Writer)
		http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameUiUserId, Value: "", Path: "/", MaxAge: -1})
		http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameUiUserMobile, Value: "", Path: "/", MaxAge: -1})
		http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameUiUserOpenId, Value: "", Path: "/", MaxAge: -1})
		http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameSiShopAccountId, Value: "", Path: "/", MaxAge: -1})
		http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameSiShopAccountName, Value: "", Path: "/", MaxAge: -1})
		http.SetCookie(c.Writer, &http.Cookie{Name: session.SessionNameSiShopAccountOpenId, Value: "", Path: "/", MaxAge: -1})
		//c.AbortWithStatusJSON(http.StatusOK, util.BuildNeedLoginResult())
		//err = util.BuildNeedLoginResult()
		//err = errors.New("未知微信openId")
		err = util.ErrNeedMiLogin
		return
	}
	params, err = serviceObj.Pay(userId, openId, clientIp, idList...)
	if err != nil {
		log.Println(err)
		err = ErrUnknown
		return
	}
	return
}

func CashCouponOrderListInCart4Ui(c *gin.Context) {
	targetPageStr := c.Query("targetPage")
	targetPage, err := strconv.Atoi(targetPageStr)
	if err != nil {
		targetPage = util.BeginPage
	}
	list, err := listCashCouponOrder4Ui(c, []int{constant.CashCouponOrderStatusInCart}, targetPage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.BuildSuccessResult(list))
}

func CashCouponOrderListUnused4Ui(c *gin.Context) {
	targetPageStr := c.Query("targetPage")
	targetPage, err := strconv.Atoi(targetPageStr)
	if err != nil {
		targetPage = util.BeginPage
	}
	list, err := listCashCouponOrder4Ui(c, []int{constant.CashCouponOrderStatusPaidNotUsed, constant.CashCouponOrderStatusNotUsedRefunding, constant.CashCouponOrderStatusNotUsedRefunded}, targetPage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.BuildSuccessResult(list))
}

func CashCouponOrderListUsed4Ui(c *gin.Context) {
	targetPageStr := c.Query("targetPage")
	targetPage, err := strconv.Atoi(targetPageStr)
	if err != nil {
		targetPage = util.BeginPage
	}
	list, err := listCashCouponOrder4Ui(c, []int{constant.CashCouponOrderStatusUsed, constant.CashCouponOrderStatusUsedRefunding, constant.CashCouponOrderStatusUsedRefunded}, targetPage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	c.JSON(http.StatusOK, util.BuildSuccessResult(list))
}

func listCashCouponOrder4Ui(c *gin.Context, statusList []int, targetPage int) (page *util.PageVO, err error) {
	sn, err := session.GetUi(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
		err = ErrUnknown
		return
	}
	userId, err := sn.Get(session.SessionNameUiUserId)
	if err != nil {
		log.Println(err)
		err = ErrUnknown
		return
	}
	serviceObj := &service.CashCouponOrder{}
	argObj := &arg.CashCouponOrder{}
	//argObj.StatusEqual = strconv.Itoa(status)
	argObj.StatusIn = statusList
	argObj.UserIdEqual = userId
	argObj.TargetPage = targetPage
	argObj.PageSize = 5
	//argObj.DisplayNames = []string{"id", "userId", "cashCouponId", "price", "refundAmount", "payOrderNumber", "number", "status"}
	argObj.OrderBy = "number"
	page, err = service.Page(serviceObj, argObj)
	if err != nil {
		log.Println(err)
		err = ErrUnknown
		return
	}
	if page.Total != 0 {
		ll := page.Datas.([]interface{})
		list := make([]*model.CashCouponOrder, len(ll), len(ll))
		ids := make([]string, 0, len(list))
		for i, cashCouponOrder := range ll {
			list[i] = cashCouponOrder.(*model.CashCouponOrder)
			if !slices.Exists(ids, list[i].CashCouponId) {
				ids = append(ids, list[i].CashCouponId)
			}
		}
		serviceCashCoupon := &service.CashCoupon{}
		argCashCoupon := &arg.CashCoupon{}
		argCashCoupon.IdsIn = ids
		var cashCouponList []interface{}
		cashCouponList, err = service.Find(serviceCashCoupon, argCashCoupon)
		if err != nil {
			log.Println(err)
			err = ErrUnknown
			return
		}

		shopIds := make([]string, 0, len(list))
		for _, v := range cashCouponList {
			shopIds = append(shopIds, v.(*model.CashCoupon).ShopId)
		}
		if len(shopIds) > 0 {
			serviceShop := &service.Shop{}
			argShop := &arg.Shop{}
			argShop.IdsIn = shopIds
			var shopList []interface{}
			shopList, err = service.Find(serviceShop, argShop)
			if err != nil {
				log.Println(err)
				c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
				return
			}
			for _, v1 := range cashCouponList {
				for _, v2 := range shopList {
					if v1.(*model.CashCoupon).ShopId == v2.(*model.Shop).Id {
						v1.(*model.CashCoupon).Shop = v2.(*model.Shop)
						break
					}
				}
			}
		}
		for _, v1 := range list {
			for _, v2 := range cashCouponList {
				if v1.CashCouponId == v2.(*model.CashCoupon).Id {
					v1.CashCoupon = v2.(*model.CashCoupon)
					break
				}
			}
		}
	}
	return
}

func CashCouponOrderDelete4Ui(c *gin.Context) {
	if strings.TrimSpace(c.Query("ids")) == "" {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("未知代金券"))
		return
	}
	ids := strings.Split(c.Query("ids"), constant.Split4Id)

	serviceObj := &service.CashCouponOrder{}
	argObj := &arg.CashCouponOrder{}
	argObj.IdsIn = ids
	argObj.StatusEqual = strconv.Itoa(constant.CashCouponOrderStatusInCart)
	result := service.ResultBatchDelete(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func CashCouponOrderList(c *gin.Context) {
	argObj := &arg.CashCouponOrder{}
	err := c.Bind(argObj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}
	argObj.OrderBy = "number"

	serviceObj := &service.CashCouponOrder{}
	argObj.DisplayNames = []string{"id", "userId", "cashCouponId", "price", "refundAmount", "payOrderNumber", "number", "status"}
	result := service.ResultList(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}

func CashCouponOrderGet(c *gin.Context) {
	serviceObj := &service.CashCouponOrder{}
	obj, err := service.Get(serviceObj, c.Param("id"))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	serviceCashCoupon := &service.CashCoupon{}
	cashCoupon, err := service.Get(serviceCashCoupon, obj.(*model.CashCouponOrder).CashCouponId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	serviceUser := &service.User{}
	user, err := service.Get(serviceUser, obj.(*model.CashCouponOrder).UserId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
		return
	}
	result := util.BuildSuccessResult(map[string]interface{}{"cashCouponOrder": obj, "cashCoupon": cashCoupon, "user": user})
	c.JSON(http.StatusOK, result)
	//serviceObj := &service.CashCouponOrder{}
	//result := service.ResultGet(serviceObj, c.Param("id"))
	//c.JSON(http.StatusOK, result)
}

func CashCouponOrderPost(c *gin.Context) {
	obj := &model.CashCouponOrder{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.CashCouponOrder{}
	result := service.ResultAdd(serviceObj, obj)
	c.JSON(http.StatusOK, result)
}

func CashCouponOrderPatch(c *gin.Context) {
	obj := &model.CashCouponOrder{}
	err := c.Bind(obj)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrParamException.Error()))
		return
	}

	serviceObj := &service.CashCouponOrder{}
	result := service.ResultUpdate(serviceObj, obj, "price", "refundAmount", "status")
	c.JSON(http.StatusOK, result)
}

func CashCouponOrderDelete(c *gin.Context) {
	ids := strings.Split(c.Query("ids"), constant.Split4Id)

	serviceObj := &service.CashCouponOrder{}
	argObj := &arg.CashCouponOrder{}
	argObj.IdsIn = ids
	result := service.ResultBatchDelete(serviceObj, argObj)
	c.JSON(http.StatusOK, result)
}
