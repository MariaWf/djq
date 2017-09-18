package service

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao"
	"mimi/djq/model"
	"mimi/djq/util"
	"mimi/djq/db/mysql"
	"mimi/djq/constant"
)

type Refund struct {
}

func (service *Refund) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.Refund{conn}
}

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
	var cashCouponOrderStatus int
	var cashCouponOrder interface{}
	daoCashCouponOrder := &dao.CashCouponOrder{conn}
	switch obj.(*model.Refund).Status {
	case constant.RefundStatusNotUsedRefunding:
		status = constant.RefundStatusNotUsedRefundCancel
		cashCouponOrderStatus = constant.CashCouponOrderStatusPaidNotUsed
	case constant.RefundStatusUsedRefunding:
		status = constant.RefundStatusUsedRefundCancel
		cashCouponOrder,err = dao.Get(daoCashCouponOrder,obj.(*model.CashCouponOrder).CashCouponId)
		if err!=nil{
			rollback = true
			err = errors.New("退款申请状态不合法")
		}
		if cashCouponOrder.(*model.CashCouponOrder).RefundAmount == 0{
			cashCouponOrderStatus = constant.CashCouponOrderStatusUsed
		}else{
			cashCouponOrderStatus = constant.CashCouponOrderStatusUsedRefunded
		}
	default:
		rollback = true
		err = errors.New("退款申请状态不合法")
		return
	}
	obj.(*model.Refund).Status = status
	_,err = dao.Update(daoObj,obj.(*model.Refund),"status")
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	cashCouponOrder.(*model.CashCouponOrder).Status = cashCouponOrderStatus
	_,err = dao.Update(daoCashCouponOrder,cashCouponOrder.(*model.CashCouponOrder),"status")
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}

	return
}
func (service *Refund) check(obj *model.Refund) error {
	if obj == nil {
		return ErrObjectEmpty
	}
	if obj.Common == "" {
		return errors.New("图片不能为空")
	}
	if err := util.MatchLenWithErr(obj.Reason, 0, 200, "退款理由"); err != nil {
		return err
	}
	if err := util.MatchLenWithErr(obj.Evidence, 0, 200, "退款凭证"); err != nil {
		return err
	}
	if err := util.MatchLenWithErr(obj.Common, 0, 200, "平台意见"); err != nil {
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
