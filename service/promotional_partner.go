package service

import (
	"database/sql"
	"mimi/djq/dao"
	"mimi/djq/dao/arg"
	"mimi/djq/db/mysql"
	"mimi/djq/model"
	"mimi/djq/util"
)

type PromotionalPartner struct {
}

func (service *PromotionalPartner) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.PromotionalPartner{conn}
}

func (service *PromotionalPartner) Delete(ids ...string) (int64, error) {
	if ids == nil || len(ids) == 0 {
		return 0, nil
	}
	conn, err := mysql.Get()
	if err != nil {
		return 0, checkErr(err)
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.PromotionalPartner)
	argObj := daoObj.GetArgInstance().(*arg.PromotionalPartner)
	argObj.SetIdsIn(ids)
	count, err := dao.LogicalDelete(daoObj, argObj)
	if err != nil {
		rollback = true
		return 0, checkErr(err)
	}
	_, err = daoObj.DeleteRelationshipWithUserByPromotionalPartnerId(ids...)
	if err != nil {
		rollback = true
		return 0, checkErr(err)
	}
	return count, checkErr(err)
}

func (service *PromotionalPartner) check(obj *model.PromotionalPartner) error {
	if obj == nil {
		return ErrObjectEmpty
	}
	if err := util.MatchLenWithErr(obj.Name, 2, 32, "名称"); err != nil {
		return err
	}
	if err := util.MatchLenWithErr(obj.Description, 0, 200, "描述"); err != nil {
		return err
	}
	if err := util.MatchNonnegativeNumberWithErr(obj.TotalPay, "已提现金额"); err != nil {
		return err
	}
	if err := util.MatchNonnegativeNumberWithErr(obj.TotalPrice, "收益金额"); err != nil {
		return err
	}
	if err := util.MatchNonnegativeNumberWithErr(obj.TotalUser, "用户数"); err != nil {
		return err
	}
	return nil
}

func (service *PromotionalPartner) CheckUpdate(obj model.BaseModelInterface) error {
	if obj != nil && obj.GetId() == "" {
		return ErrIdEmpty
	}
	return service.check(obj.(*model.PromotionalPartner))
}

func (service *PromotionalPartner) CheckAdd(obj model.BaseModelInterface) error {
	return service.check(obj.(*model.PromotionalPartner))
}
