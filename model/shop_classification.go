package model

import "github.com/pkg/errors"

type ShopClassification struct {
	Id          string `form:"id" json:"id" db:"id" desc:"id"`
	Name        string `form:"name" json:"name" db:"name" desc:"名称"`
	Priority    int    `form:"priority" json:"priority" db:"priority" desc:"优先权重"`
	Description string `form:"description" json:"description" db:"description" desc:"描述"`
	Hide        bool   `form:"hide" json:"hide" db:"hide" desc:"隐藏"`
}

func (obj *ShopClassification) GetId() string {
	return obj.Id
}

func (obj *ShopClassification) SetId(id string) {
	obj.Id = id
}

func (obj *ShopClassification) GetPointer4DB(name string) interface{} {
	switch name {
	case "id":
		return &obj.Id
	case "name":
		return &obj.Name
	case "priority": return &obj.Priority
	case "description":
		return &obj.Description
	case "hide":
		return &obj.Hide
	}
	panic(errors.New("对象shopClassification属性[" + name + "]不存在"))
}

func (obj *ShopClassification) GetValue4DB(name string) interface{} {
	switch name {
	case "id":
		return obj.Id
	case "name":
		return obj.Name
	case "priority": return obj.Priority
	case "description":
		return obj.Description
	case "hide":
		return obj.Hide
	}
	panic(errors.New("对象shopClassification属性[" + name + "]不存在"))
}

func (obj *ShopClassification) GetValue4Map(name string) interface{} {
	switch name {
	case "id":
		return obj.Id
	case "name":
		return obj.Name
	case "priority": return obj.Priority
	case "description":
		return obj.Description
	case "hide":
		return obj.Hide
	}
	panic(errors.New("对象shopClassification属性[" + name + "]不存在"))
}

