package service

import (
	"database/sql"
	"mimi/djq/dao"
	"mimi/djq/model"
	"mimi/djq/util"
)

type RefundReason struct {
}

func (service *RefundReason) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.RefundReason{conn}
}

func (service *RefundReason) check(obj *model.RefundReason) error {
	if obj == nil {
		return ErrObjectEmpty
	}
	if err := util.MatchLenWithErr(obj.Description, 0, 200, "描述"); err != nil {
		return err
	}
	if err := util.MatchPriority(obj.Priority); err != nil {
		return err
	}
	return nil
}

func (service *RefundReason) CheckUpdate(obj model.BaseModelInterface) error {
	if obj != nil && obj.GetId() == "" {
		return ErrIdEmpty
	}
	return service.check(obj.(*model.RefundReason))
}

func (service *RefundReason) CheckAdd(obj model.BaseModelInterface) error {
	return service.check(obj.(*model.RefundReason))
}
