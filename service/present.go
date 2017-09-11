package service

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao"
	"mimi/djq/model"
	"mimi/djq/util"
)

type Present struct {
}

func (service *Present) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.Present{conn}
}

func (service *Present) check(obj *model.Present) error {
	if obj == nil {
		return ErrObjectEmpty
	}
	if obj.Image == "" {
		return errors.New("图片不能为空")
	}
	if err := util.MatchLenWithErr(obj.Name, 2, 32, "名称"); err != nil {
		return err
	}
	if err := util.MatchLenWithErr(obj.Image, 0, 200, "图片路径"); err != nil {
		return err
	}
	if err := util.MatchLenWithErr(obj.Address, 0, 200, "地址"); err != nil {
		return err
	}
	if err := util.MatchWeight(obj.Weight); err != nil {
		return err
	}
	return nil
}

func (service *Present) CheckUpdate(obj model.BaseModelInterface) error {
	if obj != nil && obj.GetId() == "" {
		return ErrIdEmpty
	}
	return service.check(obj.(*model.Present))
}

func (service *Present) CheckAdd(obj model.BaseModelInterface) error {
	return service.check(obj.(*model.Present))
}
