package service

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao"
	"mimi/djq/model"
	"mimi/djq/util"
	"mimi/djq/dao/arg"
	"mimi/djq/db/mysql"
)

type CashCoupon struct {
}

func (service *CashCoupon) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.CashCoupon{conn}
}

func (service *CashCoupon) RefreshExpired() (err error) {
	conn, err := mysql.Get()
	if err != nil {
		err = checkErr(err)
		return
	}
	rollback := false
	defer mysql.Close(conn, &rollback)

	daoObj := service.GetDaoInstance(conn).(*dao.CashCoupon)
	cashCoupon := &model.CashCoupon{}
	cashCoupon.Expired = true
	argCashCoupon := &arg.CashCoupon{}
	argCashCoupon.OverExpiryDate = true
	argCashCoupon.UpdateNames = []string{"expired"}
	argCashCoupon.UpdateObject = cashCoupon
	_, err = dao.BatchUpdate(daoObj, argCashCoupon)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	return
	//cashCouponListO,err := dao.Find(daoObj,argCashCoupon)
	//if err != nil {
	//	rollback = true
	//	err = checkErr(err)
	//	return
	//}
	//cashCouponIds := make([]string,len(cashCouponListO),len(cashCouponListO))
	//for i,v:=range cashCouponListO{
	//	cashCouponIds[i] = v.(*model.CashCoupon).Id
	//}
	//
	//daoCashCouponOrder := &dao.CashCouponOrder{}
	//argCashCouponOrder := &arg.CashCouponOrder{}
	//argCashCouponOrder.CashCouponIdsIn = cashCouponIds
	//argCashCouponOrder.NotComplete = true
	//cashCouponOrder := &model.CashCouponOrder{}
	//cashCouponOrderListO,err := dao.BatchUpdate(daoCashCouponOrder,argCashCouponOrder)
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
