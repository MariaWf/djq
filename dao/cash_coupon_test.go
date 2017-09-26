package dao

import (
	"encoding/json"
	"mimi/djq/dao/arg"
	"mimi/djq/db/mysql"
	"mimi/djq/model"
	"mimi/djq/util"
	"testing"
	"time"
)

func TestCashCoupon_Add(t *testing.T) {
	conn, _ := mysql.Get()
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := &CashCoupon{conn}
	obj := &model.CashCoupon{}
	obj.ExpiryDate = util.StringTime4DB(time.Now())
	//2017-09-10 22:07:29
	_, err := Add(daoObj, obj)
	if err != nil {
		t.Error(err)
	}
	st, _ := json.Marshal(obj)

	t.Log(string(st))
}

func TestCashCoupon_List(t *testing.T) {
	conn, _ := mysql.Get()
	rollback := false
	defer mysql.Close(conn, &rollback)
	daoObj := &CashCoupon{conn}
	argObj := &arg.CashCoupon{}
	argObj.OverExpiryDate = true
	//argObj.BeforeExpiryDate= true
	list, err := Find(daoObj, argObj)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(len(list), list[0])
	}

}
