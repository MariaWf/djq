package service

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao"
	"mimi/djq/model"
	"mimi/djq/util"
)

type ShopIntroductionImage struct {
}

func (service *ShopIntroductionImage) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.ShopIntroductionImage{conn}
}

func (service *ShopIntroductionImage) check(obj *model.ShopIntroductionImage) error {
	if obj == nil {
		return ErrObjectEmpty
	}
	if obj.ShopId == "" {
		return errors.New("店铺ID为空")
	}
	if err := util.MatchPriority(obj.Priority); err != nil {
		return err
	}
	if err := util.MatchLenWithErr(obj.ContentUrl, 0, 200, "图片路径"); err != nil {
		return err
	}
	return nil
}

func (service *ShopIntroductionImage) CheckUpdate(obj model.BaseModelInterface) error {
	if obj != nil && obj.GetId() == "" {
		return ErrIdEmpty
	}
	return service.check(obj.(*model.ShopIntroductionImage))
}

func (service *ShopIntroductionImage) CheckAdd(obj model.BaseModelInterface) error {
	return service.check(obj.(*model.ShopIntroductionImage))
}
