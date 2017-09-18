package service

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao"
	"mimi/djq/model"
	"mimi/djq/db/mysql"
	"mimi/djq/dao/arg"
	"mimi/djq/wxpay"
	"mimi/djq/util"
	"log"
	"mimi/djq/constant"
	"strings"
	"time"
)

type CashCouponOrder struct {
}

func (service *CashCouponOrder) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.CashCouponOrder{conn}
}

func (service *CashCouponOrder) Complete(shopAccountId, id string) (err error) {
	if shopAccountId == "" {
		err = errors.New("未知商户")
	}
	if id == "" {
		err = errors.New("未知代金券订单")
	}

	conn, err := mysql.Get()
	if err != nil {
		err = checkErr(err)
		return
	}
	rollback := false
	defer mysql.Close(conn, &rollback)

	daoObj := service.GetDaoInstance(conn).(*dao.CashCouponOrder)
	cashCouponOrderO, err := dao.Get(daoObj, id)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	cashCouponOrder := cashCouponOrderO.(*model.CashCouponOrder)
	if cashCouponOrder.Status != constant.CashCouponOrderStatusPaidNotUsed {
		rollback = true
		err = errors.New("代金券订单状态异常")
		return
	}
	daoCashCoupon := &dao.CashCoupon{conn}
	cashCouponO, err := dao.Get(daoCashCoupon, cashCouponOrder.CashCouponId)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	cashCoupon := cashCouponO.(*model.CashCoupon)

	daoShopAccount := &dao.ShopAccount{conn}
	shopAccountO, err := dao.Get(daoShopAccount, shopAccountId)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	shopAccount := shopAccountO.(*model.ShopAccount)
	if shopAccount.ShopId != cashCoupon.ShopId {
		rollback = true
		err = errors.New("代金券不属于本店")
		return
	}
	cashCouponOrder.Status = constant.CashCouponOrderStatusUsed
	_, err = dao.Update(daoObj, cashCouponOrder, "status")
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	shopAccount.MoneyChance = shopAccount.MoneyChance + 1
	_, err = dao.Update(daoShopAccount, shopAccount, "moneyChance")
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	return
}

func (service *CashCouponOrder) BatchAddInCart(userId string, ids ... string) (list []*model.CashCouponOrder, err error) {
	list = make([]*model.CashCouponOrder, 0, len(ids))
	if len(ids) == 0 {
		err = errors.New("未包含代金券")
		return
	}
	if userId == "" {
		err = errors.New("未知用户")
		return
	}
	conn, err := mysql.Get()
	if err != nil {
		err = checkErr(err)
		return
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.CashCouponOrder)
	daoCashCoupon := &dao.CashCoupon{conn}
	argCashCoupon := &arg.CashCoupon{}
	argCashCoupon.IdsIn = ids
	cashCouponList, err := dao.Find(daoCashCoupon, argCashCoupon)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	if len(cashCouponList) == 0 {
		err = errors.New("找不到代金券")
		return
	}
	for _, v := range cashCouponList {
		cashCoupon := v.(*model.CashCoupon)
		obj := &model.CashCouponOrder{}
		obj.CashCouponId = cashCoupon.Id
		obj.Number = util.BuildCashCouponOrderNumber()
		obj.Price = cashCoupon.Price
		obj.Status = constant.CashCouponOrderStatusInCart
		obj.UserId = userId
		obj.PayBegin = util.StringDefaultTime4DB()
		obj.PayEnd = util.StringDefaultTime4DB()
		_, err = dao.Add(daoObj, obj)
		if err != nil {
			rollback = true
			err = checkErr(err)
			return
		}
		list = append(list, obj)
	}
	return
}

func (service *CashCouponOrder) Pay(userId string, openId string, clientIp string, ids ... string) (np wxpay.Params, err error) {
	np = make(wxpay.Params)
	if len(ids) == 0 {
		err = errors.New("未包含代金券订单")
		return
	}
	if userId == "" {
		err = errors.New("未知用户ID")
		return
	}
	if openId == "" {
		err = errors.New("未知用户OPENID")
		return
	}
	if clientIp == "" {
		err = errors.New("未知用户IP")
		return
	}
	conn, err := mysql.Get()
	if err != nil {
		err = checkErr(err)
		return
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.CashCouponOrder)
	argObj := daoObj.GetArgInstance().(*arg.CashCouponOrder)
	argObj.IdsIn = ids
	//argObj.UserIdEqual = userId
	list, err := dao.Find(daoObj, argObj)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	if len(list) == 0 {
		err = errors.New("找不到代金券订单")
		return
	}
	totalFee := 0
	for _, v := range list {
		obj := v.(*model.CashCouponOrder)
		if obj.UserId != userId {
			err = errors.New("代金券订单不属于当前用户")
			return
		}
		if obj.PayOrderNumber != "" || obj.Status != constant.CashCouponOrderStatusInCart {
			err = errors.New("代金券订单处于支付状态，避免重复付款请稍后重试")
			return
		}
		totalFee += obj.Price
	}
	totalFee *= 100

	payOrderNumber := util.BuildUUID()

	obj := &model.CashCouponOrder{}
	obj.PayOrderNumber = payOrderNumber
	obj.PayBegin = util.StringTime4DB(time.Now())
	obj.Status = constant.CashCouponOrderStatusPaying
	argObj.UpdateObject = obj
	argObj.UpdateNames = []string{"payOrderNumber", "payBegin", "status"}
	_, err = dao.BatchUpdate(daoObj, argObj)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}

	params, err := wxpay.UnifiedOrder(payOrderNumber, totalFee, clientIp, openId)
	if err != nil {
		err = checkErr(err)
		return
	}
	if params["return_code"] == "SUCCESS" && params["result_code"] == "SUCCESS" {
		client := wxpay.NewDefaultClient()
		np.SetString("appId", client.AppId)
		np.SetString("nonceStr", util.BuildUUID())
		np.SetString("package", "prepay_id=" + params["prepay_id"])
		np.SetString("signType", "SHA1")//MD5
		np.SetString("paySign", wxpay.Signature(np))
		return
	} else {
		log.Println(params)
		err = errors.New("支付请求失败")
		rollback = true
		return
	}
	//状态改为
	//totalFee := rand.Intn(100) + 1
	//fmt.Println(c.Request.RemoteAddr)
	//fmt.Println(c.Request.Header.Get("Remote_addr"))
	//clientIp := "211.162.109.98"//c.ClientIP()
	//fmt.Println(payOrderNumber, totalFee, clientIp)
	//sn, err := session.GetUi(c.Writer, c.Request)
	//if err != nil {
	//	log.Println(err)
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
	//	return
	//}
	//openId, err := sn.Get(session.SessionNameUiUserOpenId)
	//if err != nil {
	//	log.Println(err)
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
	//	return
	//}
	//if openId == "" {
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult("未知微信openId"))
	//	return
	//}
	//params, err := wxpay.UnifiedOrder(payOrderNumber, totalFee, clientIp, openId)
	//if err != nil {
	//	log.Println(err)
	//	c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
	//	return
	//}
	//fmt.Println(params)
	//var result *util.ResultVO
	//if params["return_code"] == "SUCCESS" && params["result_code"] == "SUCCESS" {
	//	//nonceStr: '', // 支付签名随机串，不长于 32 位
	//	//package: '', // 统一支付接口返回的prepay_id参数值，提交格式如：prepay_id=***）
	//	//	signType: '', // 签名方式，默认为'SHA1'，使用新版支付需传入'MD5'
	//	//	paySign: '', // 支付签名
	//	client := wxpay.NewDefaultClient()
	//	np := make(wxpay.Params)
	//	np.SetString("appId", client.AppId)
	//	np.SetString("nonceStr", util.BuildUUID())
	//	np.SetString("package", "prepay_id=" + params["prepay_id"])
	//	np.SetString("signType", "SHA1")//MD5
	//	np.SetString("paySign", wxpay.Signature(np))
	//	result = util.BuildSuccessResult(np)
	//} else {
	//	result = util.BuildFailResult(params["return_msg"] + params["err_code"] + params["err_code_des"])
	//}
}

func (service *CashCouponOrder) ConfirmOrder(payOrderNumber string, totalFee int) (idListStr string, err error) {
	if payOrderNumber == "" {
		err = errors.New("订单号为空")
		return
	}
	if totalFee <= 0 {
		err = errors.New("支付金额不合法")
		return
	}
	conn, err := mysql.Get()
	if err != nil {
		err = checkErr(err)
		return
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.CashCouponOrder)
	argObj := daoObj.GetArgInstance().(*arg.CashCouponOrder)
	argObj.PayOrderNumberEqual = payOrderNumber
	list, err := dao.Find(daoObj, argObj)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	total := 0
	ids := make([]string, len(list), len(list))
	for i, v := range list {
		total += v.(*model.CashCouponOrder).Price
		ids[i] = v.(*model.CashCouponOrder).Id
		status := v.(*model.CashCouponOrder).Status
		if status != constant.CashCouponOrderStatusInCart && status != constant.CashCouponOrderStatusPaidNotUsed {
			rollback = true
			err = constant.ErrWxpayConfirmIllegalOrderStatus
			return
		}
	}
	idListStr = strings.Join(ids, constant.Split4Id)
	total *= 100
	if total != totalFee {
		rollback = true
		err = constant.ErrWxpayConfirmTotalFeeNotMatch
		return
	}
	obj := &model.CashCouponOrder{}
	obj.Status = constant.CashCouponOrderStatusPaidNotUsed
	argObj.UpdateObject = obj
	argObj.UpdateNames = []string{"status"}
	_, err = dao.BatchUpdate(daoObj, argObj)
	if err != nil {
		rollback = true
		err = checkErr(err)
	}
	return
}

func (service *CashCouponOrder) CancelOrder(payOrderNumber string) (idListStr string, err error) {
	if payOrderNumber == "" {
		err = errors.New("订单号为空")
		return
	}
	conn, err := mysql.Get()
	if err != nil {
		err = checkErr(err)
		return
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.CashCouponOrder)
	argObj := daoObj.GetArgInstance().(*arg.CashCouponOrder)
	argObj.PayOrderNumberEqual = payOrderNumber
	list, err := dao.Find(daoObj, argObj)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	ids := make([]string, len(list), len(list))
	for i, v := range list {
		status := v.(*model.CashCouponOrder).Status
		ids[i] = v.(*model.CashCouponOrder).Id
		if status != constant.CashCouponOrderStatusInCart {
			rollback = true
			err = constant.ErrWxpayCancelIllegalOrderStatus
			return
		}
	}
	idListStr = strings.Join(ids, constant.Split4Id)
	obj := &model.CashCouponOrder{}
	obj.PayOrderNumber = ""
	argObj.UpdateObject = obj
	argObj.UpdateNames = []string{"payOrderNumber"}
	_, err = dao.BatchUpdate(daoObj, argObj)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	return
}

func (service *CashCouponOrder) check(obj *model.CashCouponOrder) error {
	if obj == nil {
		return ErrObjectEmpty
	}
	return nil
}

func (service *CashCouponOrder) CheckUpdate(obj model.BaseModelInterface) error {
	if obj != nil && obj.GetId() == "" {
		return ErrIdEmpty
	}
	return service.check(obj.(*model.CashCouponOrder))
}

func (service *CashCouponOrder) CheckAdd(obj model.BaseModelInterface) error {
	if obj != nil && obj.(*model.CashCouponOrder).Number == "" {
		return errors.New("代金券编号为空")
	}
	if obj != nil && obj.(*model.CashCouponOrder).UserId == "" {
		return errors.New("用户ID为空")
	}
	if obj != nil && obj.(*model.CashCouponOrder).CashCouponId == "" {
		return errors.New("代金券ID为空")
	}
	return service.check(obj.(*model.CashCouponOrder))
}
