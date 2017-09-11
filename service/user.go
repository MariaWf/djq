package service

import (
	"database/sql"
	"mimi/djq/dao"
	"mimi/djq/model"
	"mimi/djq/util"
)

type User struct {
}

func (service *User) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.User{conn}
}

func (service *User) check(obj *model.User) error {
	if obj == nil {
		return ErrObjectEmpty
	}
	if !util.MatchMobile(obj.Mobile) {
		return util.ErrMobileFormat
	}
	if err := util.MatchNonnegativeNumberWithErr(obj.PresentChance, "抽奖次数"); err != nil {
		return err
	}
	return nil
}

func (service *User) CheckUpdate(obj model.BaseModelInterface) error {
	if obj != nil && obj.GetId() == "" {
		return ErrIdEmpty
	}
	return service.check(obj.(*model.User))
}

func (service *User) CheckAdd(obj model.BaseModelInterface) error {
	return service.check(obj.(*model.User))
}
