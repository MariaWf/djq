package service

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao"
	"mimi/djq/model"
	"mimi/djq/util"
	"mimi/djq/db/mysql"
	"mimi/djq/constant"
	"strconv"
	"time"
	"mimi/djq/wxpay"
	"mimi/djq/dao/arg"
	"log"
	"fmt"
	"mimi/djq/config"
)

type Refund struct {
}

func (service *Refund) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.Refund{conn}
}

//创建一次退款申请
//状态为退款中
//退款订单编码为空
//设置关联代金券订单状态
func (service *Refund) Add(obj *model.Refund) (refund *model.Refund, err error) {
	refund = obj
	err = service.CheckAdd(obj)
	if err != nil {
		return
	}

	conn, err := mysql.Get()
	if err != nil {
		err = checkErr(err)
		return
	}
	rollback := false
	defer mysql.Close(conn, &rollback)

	//代金券订单处理
	daoCashCouponOrder := &dao.CashCouponOrder{conn}
	cashCouponOrderO, err := dao.Get(daoCashCouponOrder, refund.CashCouponOrderId)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	cashCouponOrder := cashCouponOrderO.(*model.CashCouponOrder)
	var payEnd time.Time
	payEnd, err = util.ParseTimeFromDB(cashCouponOrder.PayEnd)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	if payEnd.Add(time.Hour * 24 * 350).Before(time.Now()) {
		rollback = true
		err = errors.New("账单太久远，无法执行退款")
		return
	}
	switch cashCouponOrder.Status {
	case constant.CashCouponOrderStatusPaidNotUsed:
		cashCouponOrder.Status = constant.CashCouponOrderStatusNotUsedRefunding
		obj.Status = constant.RefundStatusNotUsedRefunding
		obj.RefundAmount = cashCouponOrder.Price
	case constant.CashCouponOrderStatusUsed:
		fallthrough
	case constant.CashCouponOrderStatusUsedRefunded:
		cashCouponOrder.Status = constant.CashCouponOrderStatusUsedRefunding
		obj.Status = constant.RefundStatusUsedRefunding
		if err = util.MatchLenWithErr(obj.Evidence, 0, 200, "退款凭证"); err != nil {
			return
		}
	default:
		rollback = true
		err = errors.New("代金券订单不符合退款状态_status:" + strconv.Itoa(cashCouponOrder.Status) + "_cashCouponOrderId:" + cashCouponOrder.Id)
		return
	}
	if obj.RefundAmount + cashCouponOrder.RefundAmount > cashCouponOrder.Price {
		rollback = true
		err = errors.New("退款总额超出订单金额_退款总额：" + strconv.Itoa(cashCouponOrder.RefundAmount) + "_订单金额：" + strconv.Itoa(cashCouponOrder.Price))
		return
	}
	_, err = dao.Update(daoCashCouponOrder, cashCouponOrder, "status")
	daoObj := service.GetDaoInstance(conn).(*dao.Refund)
	obj.RefundBegin = util.StringDefaultTime4DB()
	obj.RefundEnd = util.StringDefaultTime4DB()
	obj.RefundOrderNumber = ""
	_, err = dao.Add(daoObj, obj)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	if config.Get("task_run") != "true" {
		if obj.Status == constant.RefundStatusNotUsedRefunding {
			err = send2Wxpay(obj, obj.Comment, obj.RefundAmount, conn, &rollback)
		}
	}
	return
}

func send2Wxpay(refund *model.Refund, comment string, refundAmount int, conn *sql.Tx, rollback *bool) (err error) {
	//退款申请处理
	daoObj := &dao.Refund{conn}
	if refund.Status != constant.RefundStatusNotUsedRefunding && refund.Status != constant.RefundStatusUsedRefunding {
		*rollback = true
		err = errors.New("退款申请状态异常")
		log.Println(errors.Wrap(err, fmt.Sprintf("Status:%v", refund.Status)))
		return
	}
	if refund.RefundOrderNumber != "" {
		*rollback = true
		err = errors.New("已经进行退款流程_refundOrderNumber:" + refund.RefundOrderNumber)
		log.Println(errors.Wrap(err, fmt.Sprintf("refundOrderNumber:%v", refund.RefundOrderNumber)))
		return
	}
	refund.RefundBegin = util.StringTime4DB(time.Now())
	refund.RefundEnd = util.StringDefaultTime4DB()
	refund.RefundOrderNumber = util.BuildUUID()
	refund.RefundAmount = refundAmount
	refund.Comment = comment
	_, err = dao.Update(daoObj, refund, "refundBegin", "refundEnd", "refundOrderNumber", "refundAmount", "comment")
	if err != nil {
		*rollback = true
		err = checkErr(err)
		return
	}

	//代金券订单处理
	daoCashCouponOrder := &dao.CashCouponOrder{conn}
	cashCouponOrderO, err := dao.Get(daoCashCouponOrder, refund.CashCouponOrderId)
	if err != nil {
		*rollback = true
		err = checkErr(err)
		return
	}
	cashCouponOrder := cashCouponOrderO.(*model.CashCouponOrder)
	argCashCouponOrder := &arg.CashCouponOrder{}
	argCashCouponOrder.PayOrderNumberEqual = cashCouponOrder.PayOrderNumber
	list, err := dao.Find(daoCashCouponOrder, argCashCouponOrder)
	if err != nil {
		*rollback = true
		err = checkErr(err)
		return
	}
	totalFee := 0
	oldRefundAmount := 0
	for _, obj := range list {
		totalFee += obj.(*model.CashCouponOrder).Price
		oldRefundAmount += obj.(*model.CashCouponOrder).RefundAmount
	}
	if refund.RefundAmount + oldRefundAmount > totalFee {
		*rollback = true
		err = errors.New("退款累计金额超出订单总金额")
		log.Println(errors.Wrap(err, fmt.Sprintf("oldRefundAmount:%v_refundAmount:%v_totalFee:%v", oldRefundAmount, refund.RefundAmount, totalFee)))
		return
	}
	//调用微信退款
	err = wxpay.RefundResult(cashCouponOrder.PayOrderNumber, totalFee * 100, refund.RefundOrderNumber, refund.RefundAmount * 100)
	if err != nil {
		*rollback = true
		err = checkErr(err)
		return
	}
	return
}

//同意退款并发送微信退款请求，生成退款编码
func (service *Refund) Agree(id string, comment string, refundAmount int) (err error) {
	if id == "" {
		err = ErrIdEmpty
		return
	}
	if err = util.MatchLenWithErr(comment, 0, 200, "平台意见"); err != nil {
		return
	}
	conn, err := mysql.Get()
	if err != nil {
		err = checkErr(err)
		return
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	//退款申请处理
	daoObj := service.GetDaoInstance(conn).(*dao.Refund)
	obj, err := dao.Get(daoObj, id)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	refund := obj.(*model.Refund)
	if refund.Status != constant.RefundStatusNotUsedRefunding && refund.Status != constant.RefundStatusUsedRefunding {
		rollback = true
		err = errors.New("退款申请状态异常")
		log.Println(errors.Wrap(err, fmt.Sprintf("Status:%v", refund.Status)))
		return
	}
	if refund.RefundOrderNumber != "" {
		rollback = true
		err = errors.New("已经进行退款流程_refundOrderNumber:" + refund.RefundOrderNumber)
		log.Println(errors.Wrap(err, fmt.Sprintf("refundOrderNumber:%v", refund.RefundOrderNumber)))
		return
	}
	refund.RefundBegin = util.StringTime4DB(time.Now())
	refund.RefundEnd = util.StringDefaultTime4DB()
	refund.RefundOrderNumber = util.BuildUUID()
	refund.RefundAmount = refundAmount
	refund.Comment = comment
	_, err = dao.Update(daoObj, refund, "refundBegin", "refundEnd", "refundOrderNumber", "refundAmount", "comment")
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}

	//代金券订单处理
	daoCashCouponOrder := &dao.CashCouponOrder{conn}
	cashCouponOrderO, err := dao.Get(daoCashCouponOrder, refund.CashCouponOrderId)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	cashCouponOrder := cashCouponOrderO.(*model.CashCouponOrder)
	argCashCouponOrder := &arg.CashCouponOrder{}
	argCashCouponOrder.PayOrderNumberEqual = cashCouponOrder.PayOrderNumber
	list, err := dao.Find(daoCashCouponOrder, argCashCouponOrder)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	totalFee := 0
	oldRefundAmount := 0
	for _, obj := range list {
		totalFee += obj.(*model.CashCouponOrder).Price
		oldRefundAmount += obj.(*model.CashCouponOrder).RefundAmount
	}
	if refund.RefundAmount + oldRefundAmount > totalFee {
		rollback = true
		err = errors.New("退款累计金额超出订单总金额")
		log.Println(errors.Wrap(err, fmt.Sprintf("oldRefundAmount:%v_refundAmount:%v_totalFee:%v", oldRefundAmount, refund.RefundAmount, totalFee)))
		return
	}
	//调用微信退款
	err = wxpay.RefundResult(cashCouponOrder.PayOrderNumber, totalFee * 100, refund.RefundOrderNumber, refund.RefundAmount * 100)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	return
	//if refund.Status != constant.RefundStatusNotUsedRefunding && refund.Status != constant.RefundStatusUsedRefunding {
	//	rollback = true
	//	err = errors.New("退款申请状态异常" + strconv.Itoa(refund.Status))
	//	return
	//}
	//refund.RefundBegin = util.StringTime4DB(time.Now())
	//refund.RefundEnd = util.StringDefaultTime4DB()
	//refund.RefundOrderNumber = util.BuildUUID()
	//_, err = dao.Update(daoObj, refund, "refundBegin","refundEnd", "refundOrderNumber")
	//if err != nil {
	//	rollback = true
	//	err = checkErr(err)
	//	return
	//}
	//
	////代金券订单处理
	//daoCashCouponOrder := &dao.CashCouponOrder{conn}
	//cashCouponOrderO, err := dao.Get(daoCashCouponOrder, refund.CashCouponOrderId)
	//if err != nil {
	//	rollback = true
	//	err = checkErr(err)
	//	return
	//}
	//cashCouponOrder := cashCouponOrderO.(*model.CashCouponOrder)
	//argCashCouponOrder := &arg.CashCouponOrder{}
	//argCashCouponOrder.PayOrderNumberEqual = cashCouponOrder.PayOrderNumber
	//list, err := dao.Find(daoCashCouponOrder, argCashCouponOrder)
	//if err != nil {
	//	rollback = true
	//	err = checkErr(err)
	//	return
	//}
	//totalFee := 0
	//for _, obj := range list {
	//	totalFee += obj.(*model.CashCouponOrder).Price
	//}
	////调用微信退款
	//err = wxpay.RefundResult(cashCouponOrder.PayOrderNumber, totalFee * 100, refund.RefundOrderNumber, refund.RefundAmount * 100)
	//if err != nil {
	//	rollback = true
	//	err = checkErr(err)
	//	return
	//}
	//return
}


//拒绝退款,恢复退款关联的代金券订单状态
func (service *Refund) Reject(id string, comment string) (err error) {
	if id == "" {
		err = ErrIdEmpty
		return
	}
	if err = util.MatchLenWithErr(comment, 0, 200, "平台意见"); err != nil {
		return
	}
	conn, err := mysql.Get()
	if err != nil {
		err = checkErr(err)
		return
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	//退款申请处理
	daoObj := service.GetDaoInstance(conn).(*dao.Refund)
	obj, err := dao.Get(daoObj, id)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	refund := obj.(*model.Refund)
	if refund.RefundOrderNumber != "" {
		rollback = true
		err = errors.New("已进入汇款阶段，不能撤回")
		log.Println(errors.Wrap(err, fmt.Sprintf("RefundOrderNumber:%v", refund.RefundOrderNumber)))
		return
	}
	if refund.Status != constant.RefundStatusNotUsedRefunding && refund.Status != constant.RefundStatusUsedRefunding {
		rollback = true
		err = errors.New("退款申请状态异常" + strconv.Itoa(refund.Status))
		log.Println(errors.Wrap(err, fmt.Sprintf("Status:%v", refund.Status)))
		return
	}

	switch refund.Status {
	case constant.RefundStatusNotUsedRefunding:
		refund.Status = constant.RefundStatusNotUsedRefundFail
	case constant.RefundStatusUsedRefunding:
		refund.Status = constant.RefundStatusUsedRefundFail
	default:
		rollback = true
		err = errors.New("退款申请状态异常")
		log.Println(errors.Wrap(err, fmt.Sprintf("Status:%v", refund.Status)))
		return
	}
	refund.Comment = comment
	_, err = dao.Update(daoObj, refund, "status", "comment")
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	err = rollbackCashCouponOrderWithFailCloseOrCancel(refund.CashCouponOrderId, conn, &rollback)
	return
}

//以退款失败的方式关闭退款,恢复退款关联的代金券订单状态
func (service *Refund) FailCloseByRefundOrderNumber(refundOrderNumber string) (err error) {
	if refundOrderNumber == "" {
		err = errors.New("未知退款单号")
		return
	}
	conn, err := mysql.Get()
	if err != nil {
		err = checkErr(err)
		return
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.Refund)
	argObj := &arg.Refund{}
	argObj.RefundOrderNumberEqual = refundOrderNumber
	argObj.PageSize = 1
	list, err := dao.Find(daoObj, argObj)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	if len(list) == 0 {
		rollback = true
		err = ErrObjectNotFound
		return
	}
	refund := list[0].(*model.Refund)
	switch refund.Status {
	case constant.RefundStatusNotUsedRefunding:
		refund.Status = constant.RefundStatusNotUsedRefundFail
	case constant.RefundStatusUsedRefunding:
		refund.Status = constant.RefundStatusUsedRefundFail
	case constant.RefundStatusNotUsedRefundFail:
		fallthrough
	case constant.RefundStatusUsedRefundFail:
		rollback = true
		return
	default:
		rollback = true
		err = errors.New("退款申请状态异常" + strconv.Itoa(refund.Status))
		return
	}
	_, err = dao.Update(daoObj, refund, "status")
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	err = rollbackCashCouponOrderWithFailCloseOrCancel(refund.CashCouponOrderId, conn, &rollback)
	return
}

func rollbackCashCouponOrderWithFailCloseOrCancel(cashCouponOrderId string, conn *sql.Tx, rollback *bool) (err error) {
	daoCashCouponOrder := &dao.CashCouponOrder{conn}
	cashCouponOrderO, err := dao.Get(daoCashCouponOrder, cashCouponOrderId)
	if err != nil {
		*rollback = true
		err = checkErr(err)
		return
	}
	cashCouponOrder := cashCouponOrderO.(*model.CashCouponOrder)
	switch cashCouponOrder.Status {
	case constant.CashCouponOrderStatusNotUsedRefunding:
		cashCouponOrder.Status = constant.CashCouponOrderStatusPaidNotUsed
	case constant.CashCouponOrderStatusUsedRefunding:
		if cashCouponOrder.RefundAmount > 0 {
			cashCouponOrder.Status = constant.CashCouponOrderStatusUsedRefunded
		} else {
			cashCouponOrder.Status = constant.CashCouponOrderStatusUsed
		}
	default:
		*rollback = true
		err = errors.New("代金券订单状态异常")
		log.Println(errors.Wrap(err, fmt.Sprintf("Status:%v", cashCouponOrder.Status)))
		return
	}
	_, err = dao.Update(daoCashCouponOrder, cashCouponOrder, "status")
	if err != nil {
		*rollback = true
		err = checkErr(err)
		return
	}
	return
}

//通过退款编号确认退款完成，设置关联代金券订单状态及金额
func (service *Refund) ConfirmByRefundOrderNumber(refundOrderNumber string) (err error) {
	if refundOrderNumber == "" {
		err = errors.New("未知退款单号")
		return
	}
	conn, err := mysql.Get()
	if err != nil {
		err = checkErr(err)
		return
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	//退款申请处理
	daoObj := service.GetDaoInstance(conn).(*dao.Refund)
	argObj := &arg.Refund{}
	argObj.RefundOrderNumberEqual = refundOrderNumber
	argObj.PageSize = 1
	list, err := dao.Find(daoObj, argObj)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	if len(list) == 0 {
		rollback = true
		err = ErrObjectNotFound
		return
	}
	refund := list[0].(*model.Refund)
	err = confirmAction(refund, conn, &rollback)
	return
}

//通过退款ID确认退款完成，设置关联代金券订单状态及金额
func (service *Refund) Confirm(id string) (err error) {
	if id == "" {
		err = ErrIdEmpty
		return
	}
	conn, err := mysql.Get()
	if err != nil {
		err = checkErr(err)
		return
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	//退款申请处理
	daoObj := service.GetDaoInstance(conn).(*dao.Refund)
	obj, err := dao.Get(daoObj, id)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	refund := obj.(*model.Refund)
	err = confirmAction(refund, conn, &rollback)
	return
}

func confirmAction(refund *model.Refund, conn *sql.Tx, rollback *bool) (err error) {
	daoObj := &dao.Refund{conn}
	switch refund.Status {
	case constant.RefundStatusNotUsedRefunding:
		refund.Status = constant.RefundStatusNotUsedRefundSuccess
	case constant.RefundStatusUsedRefunding:
		refund.Status = constant.RefundStatusUsedRefundSuccess
	case constant.RefundStatusNotUsedRefundSuccess:
		fallthrough
	case constant.RefundStatusUsedRefundSuccess:
		*rollback = true
		return
	default:
		*rollback = true
		err = errors.New("退款申请状态异常")
		log.Println(errors.Wrap(err, fmt.Sprintf("Status:%v", refund.Status)))
		return
	}
	if refund.RefundOrderNumber == "" {
		*rollback = true
		err = errors.New("未知退款编号")
		log.Println(errors.Wrap(err, fmt.Sprintf("refundId:%v", refund.Id)))
		return
	}
	refund.RefundEnd = util.StringTime4DB(time.Now())
	_, err = dao.Update(daoObj, refund, "refundEnd", "status")
	if err != nil {
		*rollback = true
		err = checkErr(err)
		return
	}

	//代金券订单处理
	daoCashCouponOrder := &dao.CashCouponOrder{conn}
	cashCouponOrderO, err := dao.Get(daoCashCouponOrder, refund.CashCouponOrderId)
	if err != nil {
		*rollback = true
		err = checkErr(err)
		return
	}
	cashCouponOrder := cashCouponOrderO.(*model.CashCouponOrder)
	cashCouponOrder.RefundAmount = refund.RefundAmount + cashCouponOrder.RefundAmount
	if cashCouponOrder.RefundAmount > cashCouponOrder.Price {
		*rollback = true
		err = errors.New("退款总额超出订单金额")
		log.Println(errors.Wrap(err, fmt.Sprintf("退款总额：%v_订单金额：%v", cashCouponOrder.RefundAmount, cashCouponOrder.Price)))
		return
	}
	switch cashCouponOrder.Status {
	case constant.CashCouponOrderStatusNotUsedRefunding:
		cashCouponOrder.Status = constant.CashCouponOrderStatusNotUsedRefunded
	case constant.CashCouponOrderStatusUsedRefunding:
		cashCouponOrder.Status = constant.CashCouponOrderStatusUsedRefunded
	default:
		*rollback = true
		err = errors.New("代金券订单不符合退款状态" + strconv.Itoa(cashCouponOrder.Status))
		log.Println(errors.Wrap(err, fmt.Sprintf("cashCouponOrder.Status：%v", cashCouponOrder.Status)))
		return
	}
	_, err = dao.Update(daoCashCouponOrder, cashCouponOrder, "refundAmount", "status")
	if err != nil {
		*rollback = true
		err = checkErr(err)
		return
	}
	return
}

//func (service *Refund) ConfirmNotUsed(id string) (err error) {
//	if id == "" {
//		err = ErrIdEmpty
//		return
//	}
//	conn, err := mysql.Get()
//	if err != nil {
//		err = checkErr(err)
//		return
//	}
//	rollback := false
//	defer mysql.Close(conn, &rollback)
//	//退款申请处理
//	daoObj := service.GetDaoInstance(conn).(*dao.Refund)
//	obj, err := dao.Get(daoObj, id)
//	if err != nil {
//		rollback = true
//		err = checkErr(err)
//		return
//	}
//	refund := obj.(*model.Refund)
//	if refund.Status != constant.RefundStatusNotUsedRefunding {
//		rollback = true
//		err = errors.New("退款申请状态异常" + strconv.Itoa(refund.Status))
//		return
//	}
//	refund.RefundEnd = util.StringTime4DB(time.Now())
//	refund.Status = constant.RefundStatusNotUsedRefundSuccess
//	_, err = dao.Update(daoObj, refund, "refundEnd", "status")
//	if err != nil {
//		rollback = true
//		err = checkErr(err)
//		return
//	}
//
//	//代金券订单处理
//	daoCashCouponOrder := &dao.CashCouponOrder{conn}
//	cashCouponOrderO, err := dao.Get(daoCashCouponOrder, refund.CashCouponOrderId)
//	if err != nil {
//		rollback = true
//		err = checkErr(err)
//		return
//	}
//	cashCouponOrder := cashCouponOrderO.(*model.CashCouponOrder)
//	cashCouponOrder.RefundAmount = refund.RefundAmount + cashCouponOrder.RefundAmount
//	if cashCouponOrder.RefundAmount > cashCouponOrder.Price {
//		rollback = true
//		err = errors.New("退款总额超出订单金额_退款总额：" + strconv.Itoa(cashCouponOrder.RefundAmount) + "_订单金额：" + strconv.Itoa(cashCouponOrder.Price))
//		return
//	}
//	if cashCouponOrder.Status != constant.CashCouponOrderStatusNotUsedRefunding {
//		rollback = true
//		err = errors.New("代金券订单状态异常" + strconv.Itoa(cashCouponOrder.Status))
//		return
//	}
//	cashCouponOrder.Status = constant.CashCouponOrderStatusNotUsedRefunded
//	_, err = dao.Update(daoCashCouponOrder, cashCouponOrder, "refundAmount", "status")
//	if err != nil {
//		rollback = true
//		err = checkErr(err)
//		return
//	}
//	argCashCouponOrder := &arg.CashCouponOrder{}
//	argCashCouponOrder.PayOrderNumberEqual = cashCouponOrder.PayOrderNumber
//	list, err := dao.Find(daoCashCouponOrder, argCashCouponOrder)
//	if err != nil {
//		rollback = true
//		err = checkErr(err)
//		return
//	}
//	totalFee := 0
//	for _, obj := range list {
//		totalFee += obj.(*model.CashCouponOrder).Price
//	}
//	totalFee *= 100
//	//调用微信退款
//	err = wxpay.RefundResult(cashCouponOrder.PayOrderNumber, totalFee, refund.RefundOrderNumber, refund.RefundAmount)
//	if err != nil {
//		rollback = true
//		err = checkErr(err)
//		return
//	}
//	return
//}

//通过退款ID撤销退款申请，若存在退款编码（正在微信退款流程中）则拒绝撤销
func (service *Refund) Cancel(id string) (status int, err error) {
	if id == "" {
		err = ErrIdEmpty
		return
	}

	conn, err := mysql.Get()
	if err != nil {
		err = checkErr(err)
		return
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.Refund)
	obj, err := dao.Get(daoObj, id)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	refund := obj.(*model.Refund)
	if refund.RefundOrderNumber != "" {
		rollback = true
		err = errors.New("已进入汇款阶段，不能撤回")
		return
	}
	switch refund.Status {
	case constant.RefundStatusNotUsedRefunding:
		status = constant.RefundStatusNotUsedRefundCancel
	case constant.RefundStatusUsedRefunding:
		status = constant.RefundStatusUsedRefundCancel
	default:
		rollback = true
		err = errors.New("退款申请状态不合法")
		return
	}
	refund.Status = status
	_, err = dao.Update(daoObj, refund, "status")
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	err = rollbackCashCouponOrderWithFailCloseOrCancel(refund.CashCouponOrderId, conn, &rollback)
	return
}

func (service *Refund) check(obj *model.Refund) error {
	if obj == nil {
		return ErrObjectEmpty
	}
	//if obj.Comment == "" {
	//	return errors.New("图片不能为空")
	//}
	if obj.RefundAmount < 1 {
		return errors.New("退款金额不能为0")
	}
	if err := util.MatchLenWithErr(obj.Reason, 1, 200, "退款理由"); err != nil {
		return err
	}
	if err := util.MatchLenWithErr(obj.Evidence, 0, 200, "退款凭证"); err != nil {
		return err
	}
	if err := util.MatchLenWithErr(obj.Comment, 0, 200, "平台意见"); err != nil {
		return err
	}
	return nil
}

func (service *Refund) CheckUpdate(obj model.BaseModelInterface) error {
	if obj != nil && obj.GetId() == "" {
		return ErrIdEmpty
	}
	return service.check(obj.(*model.Refund))
}

func (service *Refund) CheckAdd(obj model.BaseModelInterface) error {
	if obj != nil && obj.(*model.Refund).CashCouponOrderId == "" {
		return errors.New("代金券订单ID为空")
	}
	return service.check(obj.(*model.Refund))
}
