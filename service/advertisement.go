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
	if !util.MatchLen(obj.Name, 2, 32) {
		return errors.New("名称长度为2到32")
	}
	if obj.Image == "" {
		return errors.New("图片不能为空")
	}
	if !util.MatchLen(obj.Image, 0, 200) {
		return errors.New("图片路径最大长度为200")
	}
	if !util.MatchLen(obj.Link, 0, 200) {
		return errors.New("链接路径最大长度为200")
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
