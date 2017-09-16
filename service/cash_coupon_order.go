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
	"fmt"
	"mimi/djq/constant"
)

type CashCouponOrder struct {
}

func (service *CashCouponOrder) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.CashCouponOrder{conn}
}

func (service *CashCouponOrder) BatchAddInCart(userId string, ids ... string) (list []*model.CashCouponOrder, err error) {
	list = make([]*model.CashCouponOrder,0,len(ids))
	if len(ids)==0{
		err = errors.New("未包含代金券")
		return
	}
	if userId == ""{
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
	daoCashCoupon := &dao.CashCoupon{}
	argCashCoupon := &arg.CashCoupon{}
	argCashCoupon.IdsIn = ids
	cashCouponList,err := dao.Find(daoCashCoupon,argCashCoupon)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	if len(cashCouponList)==0{
		err = errors.New("找不到代金券")
		return
	}
	for _,v := range cashCouponList{
		cashCoupon := v.(*model.CashCoupon)
		obj := &model.CashCouponOrder{}
		obj.CashCouponId = cashCoupon.Id
		obj.Number = util.BuildCashCouponOrderNumber()
		obj.Price = cashCoupon.Price
		obj.Status = constant.CashCouponOrderStatusInCart
		obj.UserId = userId
		_,err = dao.Add(daoObj,obj)
		if err != nil {
			rollback = true
			err = checkErr(err)
			return
		}
		list = append(list,obj)
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
		if obj.PayOrderNumber != "" {
			err = errors.New("代金券订单处于支付状态，避免重复付款请稍后重试")
			return
		}
		totalFee += obj.Price
	}

	//状态改为
	payOrderNumber := util.BuildUUID()
	params, err := wxpay.UnifiedOrder(payOrderNumber, totalFee, clientIp, openId)
	if err != nil {
		err = checkErr(err)
		return
	}
	fmt.Println(params)
	if params["return_code"] == "SUCCESS" && params["result_code"] == "SUCCESS" {
		client := wxpay.NewDefaultClient()
		np.SetString("appId", client.AppId)
		np.SetString("nonceStr", util.BuildUUID())
		np.SetString("package", "prepay_id=" + params["prepay_id"])
		np.SetString("signType", "SHA1")//MD5
		np.SetString("paySign", wxpay.Signature(np))
	} else {
		log.Println(params)
		err = errors.New("支付请求失败")
	}
	return
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
