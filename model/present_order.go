package model

import (
	"github.com/pkg/errors"
)

type PresentOrder struct {
	Id        string        `form:"id" json:"id" db:"id" desc:"id"`
	PresentId string        `form:"presentId" json:"presentId" db:"present_id" desc:"礼品ID"`
	UserId    string        `form:"userId" json:"userId" db:"user_id" desc:"用户ID"`
	Number    string        `form:"number" json:"number" db:"number" desc:"编码"`
	Status    int           `form:"status" json:"status" db:"status" desc:"状态"`
}

func (obj *PresentOrder) GetId() string {
	return obj.Id
}

func (obj *PresentOrder) SetId(id string) {
	obj.Id = id
}

func (obj *PresentOrder) GetTableName() string {
	return "tbl_present_order"
}

func (obj *PresentOrder) GetDBNames() []string {
	return []string{"id", "present_id", "user_id", "number", "status"}
}

func (obj *PresentOrder) GetMapNames() []string {
	return []string{"id", "presentId", "userId", "number", "status"}
}

func (obj *PresentOrder) GetValue4Map(name string) interface{} {
	switch name {
	case "id": return obj.Id
	case "presentId": return obj.PresentId
	case "userId": return obj.UserId
	case "number": return obj.Number
	case "status": return obj.Status
	}
	panic(errors.New("对象presentOrder属性[" + name + "]不存在"))
}

func (obj *PresentOrder) GetDBFromMapName(name string) string {
	str := GetDBFromMapName(obj, name)
	if str != "" {
		return str
	}
	panic(errors.New("对象presentOrder属性[" + name + "]不存在"))
}

func (obj *PresentOrder) GetPointer4DB(name string) interface{} {
	switch name {
	case "id": return &obj.Id
	case "present_id": return &obj.PresentId
	case "user_id": return &obj.UserId
	case "number": return &obj.Number
	case "status": return &obj.Status
	}
	panic(errors.New("对象presentOrder属性[" + name + "]不存在"))
}

func (obj *PresentOrder) GetValue4DB(name string) interface{} {
	switch name {
	case "id": return obj.Id
	case "present_id": return obj.PresentId
	case "user_id": return obj.UserId
	case "number": return obj.Number
	case "status": return obj.Status
	}
	panic(errors.New("对象presentOrder属性[" + name + "]不存在"))
}