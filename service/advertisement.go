package service

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao"
	"mimi/djq/model"
	"mimi/djq/util"
)

type Advertisement struct {
}

func (service *Advertisement) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.Advertisement{conn}
}

func (service *Advertisement) check(obj *model.Advertisement) error {
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
	if err := util.MatchLenWithErr(obj.Link, 0, 200, "链接"); err != nil {
		return err
	}
	if err := util.MatchPriority(obj.Priority); err != nil {
		return err
	}
	return nil
}

func (service *Advertisement) CheckUpdate(obj model.BaseModelInterface) error {
	if obj != nil && obj.GetId() == "" {
		return ErrIdEmpty
	}
	return service.check(obj.(*model.Advertisement))
}

func (service *Advertisement) CheckAdd(obj model.BaseModelInterface) error {
	return service.check(obj.(*model.Advertisement))
}
