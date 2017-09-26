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
	"strconv"
)

type CashCouponOrder struct {
}

func (service *CashCouponOrder) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.CashCouponOrder{conn}
}

//商家账户确认使用订单，完成
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
	if cashCoupon.Expired{
		rollback = true
		err = errors.New("代金券已过期")
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
	daoUser := &dao.User{conn}
	userO,err := dao.Get(daoUser,cashCouponOrder.UserId)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	user := userO.(*model.User)
	user.PresentChance = user.PresentChance+1
	_,err = dao.Update(daoUser,user,"presentChance")
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
		obj.CashCoupon = cashCoupon
		list = append(list, obj)
	}
	return
}

//创建支付订单，
// 如果已经创建且在生命周期内重复使用支付订单，
// 如果不在生命周期内则关闭旧订单重新创建
func (service *CashCouponOrder) Pay(userId string, openId string, clientIp string, ids ... string) (np wxpay.Params, err error) {
	np = make(wxpay.Params)
	if len(ids) == 0 {
		err = errors.New("未包含代金券订单")
		return
	}
	if len(ids) != 1 {
		err = errors.New("目前只支持提交单个代金券订单")
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
	argObj.UserIdEqual = userId
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
		if obj.Status != constant.CashCouponOrderStatusInCart {
			err = errors.New("代金券订单处于支付状态，避免重复付款请稍后重试")
			return
		}
		if obj.PayOrderNumber != "" {
			var t time.Time
			t, err = util.ParseTimeFromDB(obj.PayBegin)
			if err != nil {
				rollback = true
				err = checkErr(err)
				return
			}
			if t.Add(time.Second * 7000).After(time.Now()) {
				np = buildPayReturnParams(obj.PayOrderNumber, obj.PrepayId)
				return
			}
		}
		totalFee += obj.Price
	}
	totalFee *= 100

	if list[0].(*model.CashCouponOrder).PayOrderNumber != "" {
		var paid bool
		paid, err = wxpay.CloseOrderResult(list[0].(*model.CashCouponOrder).PayOrderNumber)
		if err != nil {
			rollback = true
			err = checkErr(err)
			return
		}
		if paid {
			rollback = true
			err = errors.New("订单已支付，请勿重复支付")
			return
		}
	}

	payOrderNumber := util.BuildUUID()

	params, err := wxpay.UnifiedOrder(payOrderNumber, totalFee, clientIp, openId)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	if params["return_code"] == "SUCCESS" && params["result_code"] == "SUCCESS" {
		np = buildPayReturnParams(payOrderNumber, params["prepay_id"])
		obj := &model.CashCouponOrder{}
		obj.PayOrderNumber = payOrderNumber
		obj.PayBegin = util.StringTime4DB(time.Now())
		obj.PrepayId = params["prepay_id"]
		//obj.Status = constant.CashCouponOrderStatusPaying
		argObj.UpdateObject = obj
		argObj.UpdateNames = []string{"payOrderNumber", "payBegin", "prepayId"}
		//argObj.UpdateNames = []string{"payOrderNumber", "payBegin", "status"}
		_, err = dao.BatchUpdate(daoObj, argObj)
		if err != nil {
			rollback = true
			err = checkErr(err)
			return
		}
		return
	} else {
		log.Println(params)
		err = errors.New("支付请求失败")
		rollback = true
		return
	}
}

func buildPayReturnParams(payOrderNumber, prepayId string) (np wxpay.Params) {
	np = make(wxpay.Params)
	client := wxpay.NewDefaultClient()
	np.SetString("appId", client.AppId)
	np.SetString("nonceStr", util.BuildUUID())
	np.SetString("timeStamp", strconv.FormatInt(time.Now().Unix(), 10))
	np.SetString("package", "prepay_id=" + prepayId)
	np.SetString("signType", "MD5")//MD5
	np.SetString("paySign", client.Sign(np))
	np.SetString("payOrderNumber", payOrderNumber)
	return
}

//确认订单已支付，改变订单状况，若是重复确认则跳过
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
	paid := false
	for i, v := range list {
		total += v.(*model.CashCouponOrder).Price
		ids[i] = v.(*model.CashCouponOrder).Id
		status := v.(*model.CashCouponOrder).Status
		if status == constant.CashCouponOrderStatusPaidNotUsed {
			paid = true
		}
		if status != constant.CashCouponOrderStatusInCart && status != constant.CashCouponOrderStatusPaidNotUsed {
			rollback = true
			err = constant.ErrWxpayConfirmIllegalOrderStatus
			return
		}
	}
	idListStr = strings.Join(ids, constant.Split4Id)
	if paid {
		rollback = true
		return
	}
	total *= 100
	if total != totalFee {
		rollback = true
		err = constant.ErrWxpayConfirmTotalFeeNotMatch
		return
	}
	obj := &model.CashCouponOrder{}
	obj.Status = constant.CashCouponOrderStatusPaidNotUsed
	obj.PayEnd = util.StringTime4DB(time.Now())
	argObj.UpdateObject = obj
	argObj.UpdateNames = []string{"status", "payEnd"}
	_, err = dao.BatchUpdate(daoObj, argObj)
	if err != nil {
		rollback = true
		err = checkErr(err)
	}
	return
}

//取消订单，先检查支付订单状态，确保安全后重置代金券订单状态
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
	paid := false
	for i, v := range list {
		status := v.(*model.CashCouponOrder).Status
		ids[i] = v.(*model.CashCouponOrder).Id
		if status == constant.CashCouponOrderStatusPaidNotUsed {
			paid = true
			continue
		}
		if status != constant.CashCouponOrderStatusInCart {
			rollback = true
			err = constant.ErrWxpayCancelIllegalOrderStatus
			return
		}
	}
	idListStr = strings.Join(ids, constant.Split4Id)
	if paid {
		rollback = true
		return
	}
	tradeState, _, err := wxpay.OrderQueryResult(payOrderNumber)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	switch tradeState {
	case "SUCCESS", "USERPAYING", "REFUND":
		rollback = true
		err = errors.New("非法支付状态_tradeState:" + tradeState + "_payOrderNumber:" + payOrderNumber)
		return
	default:
		obj := &model.CashCouponOrder{}
		obj.PayOrderNumber = ""
		obj.PayBegin = util.StringDefaultTime4DB()
		obj.PayEnd = util.StringDefaultTime4DB()
		obj.PrepayId = ""
		obj.Status = constant.CashCouponOrderStatusInCart
		argObj.UpdateObject = obj
		argObj.UpdateNames = []string{"payOrderNumber", "payBegin", "payEnd", "prepayId","status"}
		_, err = dao.BatchUpdate(daoObj, argObj)
		if err != nil {
			rollback = true
			err = checkErr(err)
			return
		}
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
