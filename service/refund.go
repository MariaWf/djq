package service

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao"
	"mimi/djq/model"
	"mimi/djq/util"
)

type Refund struct {
}

func (service *Refund) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.Refund{conn}
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
