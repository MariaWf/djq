package service

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao"
	"mimi/djq/model"
)

type CashCouponOrder struct {
}

func (service *CashCouponOrder) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.CashCouponOrder{conn}
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
