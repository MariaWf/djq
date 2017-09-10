package service

import (
	"database/sql"
	"mimi/djq/dao"
	"mimi/djq/dao/arg"
	"mimi/djq/db/mysql"
	"mimi/djq/model"
	"mimi/djq/util"
)

type ShopClassification struct {
}

func (service *ShopClassification) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.ShopClassification{conn}
}

//func (service *ShopClassification) Get(id string) (*model.ShopClassification, error) {
//	conn, err := mysql.Get()
//	if err != nil {
//		return nil, checkErr(err)
//	}
//	rollback := false
//	defer mysql.Close(conn, &rollback)
//	daoObj := service.GetDaoInstance(conn)
//	obj, err := dao.Get(daoObj, id)
//	if err != nil {
//		rollback = true
//		return nil, checkErr(err)
//	}
//	return obj.(*model.ShopClassification), nil
//}
//
//func (service *ShopClassification) Add(obj *model.ShopClassification) (*model.ShopClassification, error) {
//	err := service.CheckAdd(obj)
//	if err != nil {
//		return nil, err
//	}
//	conn, err := mysql.Get()
//	if err != nil {
//		return nil, checkErr(err)
//	}
//	rollback := false
//	defer mysql.Close(conn, &rollback)
//	daoObj := service.GetDaoInstance(conn)
//	_, err = dao.Add(daoObj, obj)
//	if err != nil {
//		rollback = true
//		return obj, checkErr(err)
//	}
//	return obj, nil
//}
//
//func (service *ShopClassification) Update(obj *model.ShopClassification) (*model.ShopClassification, error) {
//	err := service.CheckUpdate(obj)
//	if err != nil {
//		return nil, err
//	}
//	conn, err := mysql.Get()
//	if err != nil {
//		return nil, checkErr(err)
//	}
//	rollback := false
//	defer mysql.Close(conn, &rollback)
//	daoObj := service.GetDaoInstance(conn)
//	_, err = dao.Update(daoObj, obj, "name", "description", "permissionListStr")
//	if err != nil {
//		rollback = true
//		return obj, checkErr(err)
//	}
//	return obj, nil
//}

func (service *ShopClassification) Delete(shopClassificationIds ...string) (int64, error) {
	if shopClassificationIds == nil || len(shopClassificationIds) == 0 {
		return 0, ErrIdEmpty
	}
	conn, err := mysql.Get()
	if err != nil {
		return 0, checkErr(err)
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.ShopClassification)
	argObj := &arg.ShopClassification{}
	argObj.SetIdsIn(shopClassificationIds)
	count, err := dao.LogicalDelete(daoObj, argObj)
	if err != nil {
		rollback = true
		return 0, checkErr(err)
	}
	_, err = daoObj.DeleteRelationshipWithShop(shopClassificationIds...)
	if err != nil {
		rollback = true
		return 0, checkErr(err)
	}
	return count, nil
}

func (service *ShopClassification) check(obj *model.ShopClassification) error {
	if obj == nil {
		return ErrObjectEmpty
	}
	if err := util.MatchLenWithErr(obj.Name, 2, 32, "名称"); err != nil {
		return err
	}
	if err := util.MatchLenWithErr(obj.Description, 0, 200, "描述"); err != nil {
		return err
	}
	if err := util.MatchPriority(obj.Priority); err != nil {
		return err
	}
	return nil
}

func (service *ShopClassification) CheckUpdate(obj model.BaseModelInterface) error {
	if obj != nil && obj.GetId() == "" {
		return ErrIdEmpty
	}
	return service.check(obj.(*model.ShopClassification))
}

func (service *ShopClassification) CheckAdd(obj model.BaseModelInterface) error {
	return service.check(obj.(*model.ShopClassification))
}
