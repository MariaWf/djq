package model

import "github.com/pkg/errors"

type PromotionalPartner struct {
	Id          string `form:"id" json:"id" db:"id" desc:"id"`
	Name        string `form:"name" json:"name" db:"name" desc:"名称"`
	Description string `form:"description" json:"description" db:"description" desc:"描述"`
	TotalUser   int `form:"totalUser" json:"totalUser" db:"image" desc:"用户数"`
	TotalPrice  int `form:"totalPrice" json:"totalPrice" db:"name" desc:"收益金额"`
	TotalPay    int    `form:"totalPay" json:"totalPay" db:"totalPay" desc:"已提现金额"`
}

func (obj *PromotionalPartner) GetId() string {
	return obj.Id
}

func (obj *PromotionalPartner) SetId(id string) {
	obj.Id = id
}

func (obj *PromotionalPartner) GetTableName() string {
	return "tbl_promotional_partner"
}

func (obj *PromotionalPartner) GetDBNames() []string {
	return []string{"id", "name", "description", "total_user", "total_price", "total_pay"}
}

func (obj *PromotionalPartner) GetMapNames() []string {
	return []string{"id", "name", "description", "totalUser", "totalPrice", "totalPay"}
}

func (obj *PromotionalPartner) GetValue4Map(name string) interface{} {
	switch name {
	case "id": return obj.Id
	case "name": return obj.Name
	case "description": return obj.Description
	case "totalUser": return obj.TotalUser
	case "totalPrice": return obj.TotalPrice
	case "totalPay": return obj.TotalPay
	}
	panic(errors.New("对象promotionalPartner属性[" + name + "]不存在"))
}

func (obj *PromotionalPartner) GetDBFromMapName(name string) string {
	str := GetDBFromMapName(obj, name)
	if str != "" {
		return str
	}
	panic(errors.New("对象promotionalPartner属性[" + name + "]不存在"))
}

func (obj *PromotionalPartner) GetPointer4DB(name string) interface{} {
	switch name {
	case "id": return &obj.Id
	case "name": return &obj.Name
	case "description": return &obj.Description
	case "total_user": return &obj.TotalUser
	case "total_price": return &obj.TotalPrice
	case "total_pay": return &obj.TotalPay
	}
	panic(errors.New("对象promotionalPartner属性[" + name + "]不存在"))
}

func (obj *PromotionalPartner) GetValue4DB(name string) interface{} {
	switch name {
	case "id": return obj.Id
	case "name": return obj.Name
	case "description": return obj.Description
	case "total_user": return obj.TotalUser
	case "total_price": return obj.TotalPrice
	case "total_pay": return obj.TotalPay
	}
	panic(errors.New("对象promotionalPartner属性[" + name + "]不存在"))
}
