package service

import (
	"database/sql"
	"github.com/pkg/errors"
	"mimi/djq/dao"
	"mimi/djq/model"
	"mimi/djq/db/mysql"
	"strings"
	"mimi/djq/constant"
	"mimi/djq/dao/arg"
	"math/rand"
)

type PresentOrder struct {
}

func (service *PresentOrder) GetDaoInstance(conn *sql.Tx) dao.BaseDaoInterface {
	return &dao.PresentOrder{conn}
}
func (service *PresentOrder) Complete( id string) (err error) {
	if id == "" {
		err = errors.New("未知礼品订单")
	}

	conn, err := mysql.Get()
	if err != nil {
		err = checkErr(err)
		return
	}
	rollback := false
	defer mysql.Close(conn, &rollback)

	daoObj := service.GetDaoInstance(conn).(*dao.PresentOrder)
	presentOrderO, err := dao.Get(daoObj, id)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	presentOrder := presentOrderO.(*model.PresentOrder)
	if presentOrder.Status != constant.PresentOrderStatusWaiting2Receive {
		rollback = true
		err = errors.New("礼品订单状态异常")
		return
	}
	daoPresent := &dao.Present{conn}
	presentO, err := dao.Get(daoPresent, presentOrder.PresentId)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	present := presentO.(*model.Present)
	if present.Stock < 1 || present.Requirement < 1 {
		rollback = true
		err = errors.New("礼品存货异常")
		return
	}
	present.Stock = present.Stock - 1
	present.Requirement = present.Requirement - 1
	_, err = dao.Update(daoPresent, present, "stock", "requirement")
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	presentOrder.Status = constant.PresentOrderStatusReceived
	_, err = dao.Update(daoObj, presentOrder, "status")
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	return
}
func (service *PresentOrder) Random(userId, presentIds string) (id string, err error) {
	if userId == "" {
		err = errors.New("未知用户ID")
		return
	}
	if presentIds == "" {
		err = errors.New("未知礼品")
	}

	conn, err := mysql.Get()
	if err != nil {
		err = checkErr(err)
		return
	}
	rollback := false
	defer mysql.Close(conn, &rollback)

	daoUser := &dao.User{conn}
	userO, err := dao.Get(daoUser, userId)
	user := userO.(*model.User)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	if user == nil || user.Id == "" {
		rollback = true
		err = errors.New("未知用户")
		return
	}
	if user.Locked {
		rollback = true
		err = errors.New("用户已冻结")
		return
	}
	if user.PresentChance < 1 {
		rollback = true
		err = errors.New("抽奖机会为0")
		return
	}
	user.PresentChance = user.PresentChance - 1
	_, err = dao.Update(daoUser, user, "presentChance")
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	ids := strings.Split(presentIds, constant.Split4Id)
	daoPresent := &dao.Present{conn}
	argPresent := &arg.Present{}
	argPresent.IdsIn = ids
	argPresent.NotIncludeHide = true
	argPresent.BeforeExpiryDate = true
	presentList, err := dao.Find(daoPresent, argPresent)
	if err != nil {
		rollback = true
		err = checkErr(err)
		return
	}
	weightTotal := 0
	stock := 0
	for _, v := range presentList {
		p := v.(*model.Present)
		weightTotal += p.Weight
		stock += (p.Stock - p.Requirement)
	}
	if weightTotal == 0 || stock == 0 {
		rollback = true
		err = errors.New("抽奖异常，请于管理员联系")
		return
	}
	randomWeight := rand.Intn(weightTotal)
	for _, v := range presentList {
		p := v.(*model.Present)
		randomWeight -= p.Weight
		if randomWeight < 0 {
			if (p.Stock - p.Requirement > 0) {
				p.Requirement = p.Requirement + 1
				_, err = dao.Update(daoPresent, p, "requirement")
				if err != nil {
					rollback = true
					err = checkErr(err)
					return
				}
				id = p.Id
				break
			}
		}
	}
	if id == "" {
		for _, v := range presentList {
			p := v.(*model.Present)
			if (p.Stock - p.Requirement > 0) {
				p.Requirement = p.Requirement + 1
				_, err = dao.Update(daoPresent, p, "requirement")
				if err != nil {
					rollback = true
					err = checkErr(err)
					return
				}
				id = p.Id
				break
			}
		}
	}
	return
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