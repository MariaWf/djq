package service

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao"
	"mimi/djq/model"
	"mimi/djq/util"
	"log"
	"mimi/djq/dao/arg"
	"mimi/djq/db/mysql"
	"math/rand"
)

type ShopAccount struct {
}

func (service *ShopAccount) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.ShopAccount{conn}
}

func (service *ShopAccount) GetMoney(shopAccountId string) (money int, err error) {
	if shopAccountId == "" {
		err = errors.New("未知商户")
	}
	conn, err := mysql.Get()
	if err != nil {
		err = checkErr(err)
		return
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.ShopAccount)
	shopAccountO, err := dao.Get(daoObj, shopAccountId)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	shopAccount := shopAccountO.(*model.ShopAccount)
	if shopAccount.MoneyChance < 1 {
		rollback = true
		err = errors.New("抽红包机会为0")
		return
	}
	money = rand.Intn(10)
	shopAccount.MoneyChance = shopAccount.MoneyChance - 1
	shopAccount.TotalMoney = shopAccount.TotalMoney + money
	_, err = dao.Update(daoObj, shopAccount, "moneyChance", "totalMoney")
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	return
}

func (service *ShopAccount) CheckLogin(obj *model.ShopAccount) (*model.ShopAccount, error) {
	if obj == nil {
		return nil, ErrObjectEmpty
	}
	if obj.Name == "" {
		return nil, errors.New("登录名为空")
	}
	var err error
	obj.Password, err = util.DecryptPassword(obj.Password)
	if err != nil {
		log.Println(errors.Wrap(err, "未能识别登录密码"))
		return nil, errors.New("未能识别登录密码")
	}
	obj.Password = util.BuildPassword4DB(obj.Password)
	argObj := &arg.ShopAccount{}
	argObj.NameEqual = obj.Name
	argObj.PasswordEqual = obj.Password
	argObj.PageSize = 1
	conn, err := mysql.Get()
	if err != nil {
		return nil, checkErr(err)
	}
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.ShopAccount)
	list, err := dao.Find(daoObj, argObj)
	if err != nil {
		rollback = true
		return nil, checkErr(err)
	}
	if list == nil || len(list) == 0 {
		rollback = true
		return nil, errors.New("登录名或登录密码不正确")
	}
	newObj := list[0].(*model.ShopAccount)
	if newObj.Locked {
		rollback = true
		return nil, errors.New("账号已被冻结")
	}
	return newObj, nil
}

func (service *ShopAccount) Add(obj *model.ShopAccount) (*model.ShopAccount, error) {
	sourcePassword, err := util.DecryptPassword(obj.Password)
	if err != nil {
		return nil, checkErr(err)
	}
	obj.Password = sourcePassword

	err = service.CheckAdd(obj)
	if err != nil {
		return nil, err
	}
	obj.Password = util.BuildPassword4DB(obj.Password)

	conn, err := mysql.Get()
	if err != nil {
		return nil, checkErr(err)
	}

	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn).(*dao.ShopAccount)

	argObj := &arg.ShopAccount{}
	argObj.NameEqual = obj.Name
	count, err := dao.Count(daoObj, argObj)
	if err != nil {
		rollback = true
		return nil, checkErr(err)
	}
	if count > 0 {
		rollback = true
		return nil, errors.New("用户名已存在")
	}

	_, err = dao.Add(daoObj, obj)
	if err != nil {
		rollback = true
		return nil, checkErr(err)
	}
	return obj, nil
}

func (service *ShopAccount) Update(obj *model.ShopAccount) (*model.ShopAccount, error) {
	sourcePassword, err := util.DecryptPassword(obj.Password)
	if err != nil {
		return nil, checkErr(err)
	}
	obj.Password = sourcePassword

	err = service.CheckUpdate(obj)
	if err != nil {
		return nil, err
	}
	if obj.Password != "" {
		obj.Password = util.BuildPassword4DB(obj.Password)
	}

	conn, err := mysql.Get()
	if err != nil {
		return nil, checkErr(err)
	}

	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := service.GetDaoInstance(conn)

	argObj := &arg.ShopAccount{}
	argObj.NameEqual = obj.Name
	list, err := dao.Find(daoObj, argObj)
	if err != nil {
		rollback = true
		return nil, checkErr(err)
	}
	if len(list) > 1 || (len(list) > 0 && list[0].(*model.ShopAccount).GetId() != obj.GetId()) {
		rollback = true
		return nil, errors.New("用户名已存在")
	}

	if obj.Password != "" {
		_, err = dao.Update(daoObj, obj, "name", "description", "moneyChance", "totalMoney", "locked")
	} else {
		_, err = dao.Update(daoObj, obj, "name", "password", "description", "moneyChance", "totalMoney", "locked")
	}

	if err != nil {
		rollback = true
		return nil, checkErr(err)
	}
	obj.Password = ""
	if err != nil {
		rollback = true
		return nil, checkErr(err)
	}
	return obj, checkErr(err)
}

func (service *ShopAccount) check(obj *model.ShopAccount) error {
	if obj == nil {
		return ErrObjectEmpty
	}
	if obj.ShopId == "" {
		return errors.New("店铺ID为空")
	}
	if err := util.MatchLenWithErr(obj.Name, 2, 32, "名称"); err != nil {
		return err
	}
	if err := util.MatchLenWithErr(obj.Description, 0, 200, "描述"); err != nil {
		return err
	}
	return nil
}

func (service *ShopAccount) CheckUpdate(obj model.BaseModelInterface) error {
	if obj != nil && obj.GetId() == "" {
		return ErrIdEmpty
	}
	return service.check(obj.(*model.ShopAccount))
}

func (service *ShopAccount) CheckAdd(obj model.BaseModelInterface) error {
	if obj != nil && obj.(*model.ShopAccount).Password == "" {
		return util.ErrPasswordFormat
	}
	return service.check(obj.(*model.ShopAccount))
}
