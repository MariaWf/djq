package model

import "github.com/pkg/errors"

type ShopAccount struct {
	Id          string `form:"id" json:"id" db:"id" desc:"id"`
	ShopId      string `form:"shopId" json:"shopId" db:"shop_id" desc:"shopId"`
	Name        string `form:"name" json:"name" db:"name" desc:"名称"`
	Password    string `form:"password" json:"password" db:"password" desc:"密码"`
	Description string `form:"description" json:"description" db:"description" desc:"描述"`
	MoneyChance int    `form:"moneyChance" json:"moneyChance" db:"money_chance" desc:"优先权重"`
	TotalMoney  int    `form:"totalMoney" json:"totalMoney" db:"total_money" desc:"优先权重"`
	Locked      bool   `form:"locked" json:"locked" db:"locked" desc:"隐藏"`
}

func (obj *ShopAccount) GetId() string {
	return obj.Id
}

func (obj *ShopAccount) SetId(id string) {
	obj.Id = id
}

func (obj *ShopAccount) GetTableName() string {
	return "tbl_shop_account"
}

func (obj *ShopAccount) GetDBNames() []string {
	return []string{"id", "shop_id", "name", "password", "description", "money_chance", "total_money", "locked"}
}

func (obj *ShopAccount) GetMapNames() []string {
	return []string{"id", "shopId", "name", "password", "description", "moneyChance", "totalMoney", "locked"}
}

func (obj *ShopAccount) GetValue4Map(name string) interface{} {
	switch name {
	case "id":
		return obj.Id
	case "shopId":
		return obj.ShopId
	case "name":
		return obj.Name
	case "password":
		return obj.Password
	case "description":
		return obj.Description
	case "moneyChance":
		return obj.MoneyChance
	case "totalMoney":
		return obj.TotalMoney
	case "locked":
		return obj.Locked
	}
	panic(errors.New("对象shopAccount属性[" + name + "]不存在"))
}

func (obj *ShopAccount) GetDBFromMapName(name string) string {
	str := GetDBFromMapName(obj, name)
	if str != "" {
		return str
	}
	panic(errors.New("对象shopAccount属性[" + name + "]不存在"))
}

func (obj *ShopAccount) GetPointer4DB(name string) interface{} {
	switch name {
	case "id":
		return &obj.Id
	case "shop_id":
		return &obj.ShopId
	case "name":
		return &obj.Name
	case "password":
		return &obj.Password
	case "description":
		return &obj.Description
	case "money_chance":
		return &obj.MoneyChance
	case "total_money":
		return &obj.TotalMoney
	case "locked":
		return &obj.Locked
	}
	panic(errors.New("对象shopAccount属性[" + name + "]不存在"))
}

func (obj *ShopAccount) GetValue4DB(name string) interface{} {
	switch name {
	case "id":
		return obj.Id
	case "shop_id":
		return obj.ShopId
	case "name":
		return obj.Name
	case "password":
		return obj.Password
	case "description":
		return obj.Description
	case "money_chance":
		return obj.MoneyChance
	case "total_money":
		return obj.TotalMoney
	case "locked":
		return obj.Locked
	}
	panic(errors.New("对象shopAccount属性[" + name + "]不存在"))
}
