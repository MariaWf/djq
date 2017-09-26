package model

import (
	"github.com/pkg/errors"
)

type Present struct {
	Id          string `form:"id" json:"id" db:"id" desc:"id"`
	Name        string `form:"name" json:"name" db:"name" desc:"名称"`
	Image       string `form:"image" json:"image" db:"image" desc:"地址"`
	Address     string `form:"address" json:"address" db:"address" desc:"地址"`
	Stock       int    `form:"stock" json:"stock" db:"stock" desc:"库存"`
	Requirement int    `form:"requirement" json:"requirement" db:"requirement" desc:"需求"`
	Weight      int    `form:"weight" json:"weight" db:"weight" desc:"抽奖权重"`
	ExpiryDate  string `form:"expiryDate" json:"expiryDate" db:"expiry_date" desc:"有效时间"`
	Expired     bool   `form:"expired" json:"expired" db:"expired" desc:"已过期"`
	Hide        bool   `form:"hide" json:"hide" db:"hide" desc:"隐藏"`
}

func (obj *Present) GetId() string {
	return obj.Id
}

func (obj *Present) SetId(id string) {
	obj.Id = id
}

func (obj *Present) GetTableName() string {
	return "tbl_present"
}

func (obj *Present) GetDBNames() []string {
	return []string{"id", "name", "image", "address", "stock", "requirement", "weight", "expiry_date", "expired", "hide"}
}

func (obj *Present) GetMapNames() []string {
	return []string{"id", "name", "image", "address", "stock", "requirement", "weight", "expiryDate", "expired", "hide"}
}

func (obj *Present) GetValue4Map(name string) interface{} {
	switch name {
	case "id":
		return obj.Id
	case "name":
		return obj.Name
	case "image":
		return obj.Image
	case "address":
		return obj.Address
	case "stock":
		return obj.Stock
	case "requirement":
		return obj.Requirement
	case "weight":
		return obj.Weight
	case "expiryDate":
		return obj.ExpiryDate
	case "expired":
		return obj.Expired
	case "hide":
		return obj.Hide
	}
	panic(errors.New("对象present属性[" + name + "]不存在"))
}

func (obj *Present) GetDBFromMapName(name string) string {
	str := GetDBFromMapName(obj, name)
	if str != "" {
		return str
	}
	panic(errors.New("对象present属性[" + name + "]不存在"))
}

func (obj *Present) GetPointer4DB(name string) interface{} {
	switch name {
	case "id":
		return &obj.Id
	case "name":
		return &obj.Name
	case "image":
		return &obj.Image
	case "address":
		return &obj.Address
	case "stock":
		return &obj.Stock
	case "requirement":
		return &obj.Requirement
	case "weight":
		return &obj.Weight
	case "expiry_date":
		return &obj.ExpiryDate
	case "expired":
		return &obj.Expired
	case "hide":
		return &obj.Hide
	}
	panic(errors.New("对象present属性[" + name + "]不存在"))
}

func (obj *Present) GetValue4DB(name string) interface{} {
	switch name {
	case "id":
		return obj.Id
	case "name":
		return obj.Name
	case "image":
		return obj.Image
	case "address":
		return obj.Address
	case "stock":
		return obj.Stock
	case "requirement":
		return obj.Requirement
	case "weight":
		return obj.Weight
	case "expiry_date":
		return obj.ExpiryDate
	case "expired":
		return obj.Expired
	case "hide":
		return obj.Hide
	}
	panic(errors.New("对象present属性[" + name + "]不存在"))
}
