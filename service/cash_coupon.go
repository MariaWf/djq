package service

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao"
	"mimi/djq/model"
	"mimi/djq/util"
)

type CashCoupon struct {
}

func (service *CashCoupon) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.CashCoupon{conn}
}

func (service *CashCoupon) check(obj *model.CashCoupon) error {
	if obj == nil {
		return ErrObjectEmpty
	}
	if obj.ShopId == "" {
		return errors.New("店铺ID为空")
	}
	if err := util.MatchPriority(obj.Priority); err != nil {
		return err
	}
	if err := util.MatchLenWithErr(obj.Name, 2, 32, "名称"); err != nil {
		return err
	}
	if err := util.MatchLenWithErr(obj.PreImage, 0, 200, "图片路径"); err != nil {
		return err
	}
	return nil
}

func (service *CashCoupon) CheckUpdate(obj model.BaseModelInterface) error {
	if obj != nil && obj.GetId() == "" {
		return ErrIdEmpty
	}
	return service.check(obj.(*model.CashCoupon))
}

func (service *CashCoupon) CheckAdd(obj model.BaseModelInterface) error {
	return service.check(obj.(*model.CashCoupon))
}
