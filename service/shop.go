package service

import (
	"database/sql"
	"mimi/djq/dao"
	"mimi/djq/dao/arg"
	"mimi/djq/db/mysql"
	"mimi/djq/model"
	"mimi/djq/util"
)

type Shop struct {
}

func (service *Shop) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.Shop{conn}
}

func (service *Shop) FindByShopClassificationId(argObj *arg.Shop, shopClassificationId string) (result *util.ResultVO) {
	conn, err := mysql.Get()
	if err != nil {
		err = checkErr(err)
		result = util.BuildFailResult(err.Error())
		return
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.Shop)
	if shopClassificationId != "" {
		ids, err := daoObj.ListShopIdsByShopClassificationId(shopClassificationId)
		if err != nil {
			rollback = true
			err = checkErr(err)
			result = util.BuildFailResult(err.Error())
			return
		}
		argObj.IdsIn = ids
	}
	total, err := dao.Count(daoObj, argObj)
	if err != nil {
		rollback = true
		err = checkErr(err)
		result = util.BuildFailResult(err.Error())
		return
	}
	list, err := dao.Find(daoObj, argObj)
	if err != nil {
		rollback = true
		err = checkErr(err)
		result = util.BuildFailResult(err.Error())
		return
	}
	result = util.BuildSuccessPageResult(argObj.TargetPage, argObj.PageSize, total, list)
	return
}

func (service *Shop) Get(id string) (*model.Shop, error) {
	conn, err := mysql.Get()
	if err != nil {
		return nil, checkErr(err)
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.Shop)
	obj, err := dao.Get(daoObj, id)
	if err != nil {
		rollback = true
		return nil, checkErr(err)
	}

	shopClassificationDao := &dao.ShopClassification{conn}
	shopClassificationIds, err := shopClassificationDao.ListShopClassificationIdsByShopId(id)
	if err != nil {
		rollback = true
		return nil, checkErr(err)
	}

	if shopClassificationIds != nil && len(shopClassificationIds) != 0 {
		argShopClassification := shopClassificationDao.GetArgInstance().(*arg.ShopClassification)
		argShopClassification.SetIdsIn(shopClassificationIds)
		shopClassifications, err := dao.Find(shopClassificationDao, argShopClassification)
		if err != nil {
			rollback = true
			return nil, checkErr(err)
		}
		obj.(*model.Shop).SetShopClassificationListFromInterfaceArr(shopClassifications)
	}
	return obj.(*model.Shop), nil
}

func (service *Shop) Add(obj *model.Shop) (*model.Shop, error) {
	conn, err := mysql.Get()
	if err != nil {
		return nil, checkErr(err)
	}

	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.Shop)
	_, err = dao.Add(daoObj, obj)
	if err != nil {
		rollback = true
		return nil, checkErr(err)
	}

	if obj.ShopClassificationList != nil && len(obj.ShopClassificationList) != 0 {
		toAddShopClassificationIds := make([]string, len(obj.ShopClassificationList), len(obj.ShopClassificationList))
		for i, shopClassification := range obj.ShopClassificationList {
			toAddShopClassificationIds[i] = shopClassification.GetId()
		}
		if !util.IsStringArrEmpty(toAddShopClassificationIds) {
			for _, shopClassificationId := range toAddShopClassificationIds {
				err := daoObj.AddRelationshipWithShopClassification(obj.GetId(), shopClassificationId)
				if err != nil {
					rollback = true
					return nil, err
				}
			}
		}
	}
	return obj, nil
}

func (service *Shop) Update(obj *model.Shop) (*model.Shop, error) {
	conn, err := mysql.Get()
	if err != nil {
		return nil, checkErr(err)
	}

	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn)

	_, err = dao.Update(daoObj, obj, "name", "titleFirst", "titleSecond", "phoneNumber", "logo", "preImage", "totalCashCouponNumber", "totalCashCouponPrice", "introduction", "address", "priority", "hide")
	if err != nil {
		rollback = true
		return nil, checkErr(err)
	}
	err = service.refreshRelationshipWithShopClassification(conn, obj)
	if err != nil {
		rollback = true
		return nil, checkErr(err)
	}
	return obj, checkErr(err)
}

func (service *Shop) refreshRelationshipWithShopClassification(conn *sql.Tx, obj *model.Shop) error {
	daoObj := service.GetDaoInstance(conn).(*dao.Shop)

	if obj.ShopClassificationList == nil && len(obj.ShopClassificationList) == 0 {
		_, err := daoObj.DeleteRelationshipWithShopClassificationByShopId(obj.GetId())
		if err != nil {
			return err
		}
		return nil
	}

	toUpdateShopClassificationIds := make([]string, len(obj.ShopClassificationList), len(obj.ShopClassificationList))
	for i, shopClassification := range obj.ShopClassificationList {
		toUpdateShopClassificationIds[i] = shopClassification.GetId()
	}

	shopClassificationDao := &dao.ShopClassification{conn}

	existShopClassificationIds, err := shopClassificationDao.ListShopClassificationIdsByShopId(obj.GetId())
	if err != nil {
		return err
	}

	var toAddShopClassificationIds []string
	var toDeleteShopClassificationIds []string
	if existShopClassificationIds == nil || len(existShopClassificationIds) == 0 {
		toAddShopClassificationIds = util.StringArrCopy(toUpdateShopClassificationIds)
		toDeleteShopClassificationIds = nil
	} else {
		toAddShopClassificationIds = make([]string, 0, 10)
		toDeleteShopClassificationIds = util.StringArrCopy(existShopClassificationIds)
		for _, toUpdateShopClassificationId := range toUpdateShopClassificationIds {
			exist := false
			for _, existShopClassificationId := range existShopClassificationIds {
				if toUpdateShopClassificationId == existShopClassificationId {
					exist = true
					toDeleteShopClassificationIds = util.StringArrDelete(toDeleteShopClassificationIds, toUpdateShopClassificationId)
				}
			}
			if !exist {
				toAddShopClassificationIds = append(toAddShopClassificationIds, toUpdateShopClassificationId)
			}
		}
	}
	if !util.IsStringArrEmpty(toAddShopClassificationIds) {
		for _, shopClassificationId := range toAddShopClassificationIds {
			err := daoObj.AddRelationshipWithShopClassification(obj.GetId(), shopClassificationId)
			if err != nil {
				return err
			}
		}
	}
	if !util.IsStringArrEmpty(toDeleteShopClassificationIds) {
		_, err := shopClassificationDao.DeleteRelationshipWithShop(toDeleteShopClassificationIds...)
		if err != nil {
			return err
		}
	}

	return nil
}

func (service *Shop) Delete(ids ...string) (int64, error) {
	if ids == nil || len(ids) == 0 {
		return 0, nil
	}
	conn, err := mysql.Get()
	if err != nil {
		return 0, checkErr(err)
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.Shop)
	argObj := daoObj.GetArgInstance().(*arg.Shop)
	argObj.SetIdsIn(ids)
	count, err := dao.LogicalDelete(daoObj, argObj)
	if err != nil {
		rollback = true
		return 0, checkErr(err)
	}
	_, err = daoObj.DeleteRelationshipWithShopClassificationByShopId(ids...)
	if err != nil {
		rollback = true
		return 0, checkErr(err)
	}
	return count, checkErr(err)

}

func (service *Shop) check(obj *model.Shop) error {
	if obj == nil {
		return ErrObjectEmpty
	}
	if err := util.MatchLenWithErr(obj.Name, 2, 32, "名称"); err != nil {
		return err
	}
	if err := util.MatchLenWithErr(obj.Address, 0, 200, "地址"); err != nil {
		return err
	}
	if err := util.MatchLenWithErr(obj.Introduction, 0, 200, "介绍"); err != nil {
		return err
	}
	if err := util.MatchLenWithErr(obj.Logo, 0, 200, "商标"); err != nil {
		return err
	}
	if err := util.MatchLenWithErr(obj.PreImage, 0, 200, "预览图"); err != nil {
		return err
	}
	if err := util.MatchPriority(obj.Priority); err != nil {
		return err
	}
	return nil
}

func (service *Shop) CheckUpdate(obj model.BaseModelInterface) error {
	if obj != nil && obj.GetId() == "" {
		return ErrIdEmpty
	}
	return service.check(obj.(*model.Shop))
}

func (service *Shop) CheckAdd(obj model.BaseModelInterface) error {
	return service.check(obj.(*model.Shop))
}
