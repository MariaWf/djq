package service

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao"
	"mimi/djq/model"
)

type PresentOrder struct {
}

func (service *PresentOrder) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.PresentOrder{conn}
}

func (service *PresentOrder) check(obj *model.PresentOrder) error {
	if obj == nil {
		return ErrObjectEmpty
	}
	return nil
}

func (service *PresentOrder) CheckUpdate(obj model.BaseModelInterface) error {
	if obj != nil && obj.GetId() == "" {
		return ErrIdEmpty
	}
	return service.check(obj.(*model.PresentOrder))
}

func (service *PresentOrder) CheckAdd(obj model.BaseModelInterface) error {
	if obj != nil && obj.(*model.PresentOrder).PresentId == "" {
		return errors.New("礼品ID为空")
	}
	if obj != nil && obj.(*model.PresentOrder).UserId == "" {
		return errors.New("用户ID为空")
	}
	return service.check(obj.(*model.PresentOrder))
}
