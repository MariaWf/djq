package model

import "errors"

type User struct {
	Id                   string `form:"id" json:"id" db:"id" desc:"id"`
	PromotionalPartnerId string `form:"promotionalPartnerId" json:"promotionalPartnerId" db:"promotional_partner_id" desc:"推荐人ID"`
	Mobile               string `form:"mobile" json:"mobile" db:"mobile" desc:"手机"`
	PresentChance        int    `form:"presentChance" json:"presentChance" db:"present_chance" desc:"抽奖次数"`
	Shared               bool   `form:"shared" json:"shared" db:"shared" desc:"已分享"`
	Locked               bool   `form:"locked" json:"locked" db:"locked" desc:"锁定"`
}

func (obj *User) GetId() string {
	return obj.Id
}

func (obj *User) SetId(id string) {
	obj.Id = id
}

func (obj *User) GetTableName() string {
	return "tbl_user"
}

func (obj *User) GetDBNames() []string {
	return []string{"id", "promotional_partner_id", "mobile", "present_chance", "shared", "locked"}
}

func (obj *User) GetMapNames() []string {
	return []string{"id", "promotionalPartnerId", "mobile", "presentChance", "shared", "locked"}
}

func (obj *User) GetValue4Map(name string) interface{} {
	switch name {
	case "id":
		return obj.Id
	case "promotionalPartnerId":
		return obj.PromotionalPartnerId
	case "mobile":
		return obj.Mobile
	case "presentChance":
		return obj.PresentChance
	case "shared":
		return obj.Shared
	case "locked":
		return obj.Locked
	}
	panic(errors.New("对象user属性[" + name + "]不存在"))
}

func (obj *User) GetDBFromMapName(name string) string {
	str := GetDBFromMapName(obj, name)
	if str != "" {
		return str
	}
	panic(errors.New("对象user属性[" + name + "]不存在"))
}

func (obj *User) GetPointer4DB(name string) interface{} {
	switch name {
	case "id":
		return &obj.Id
	case "promotional_partner_id":
		return &obj.PromotionalPartnerId
	case "mobile":
		return &obj.Mobile
	case "present_chance":
		return &obj.PresentChance
	case "shared":
		return &obj.Shared
	case "locked":
		return &obj.Locked
	}
	panic(errors.New("对象user属性[" + name + "]不存在"))
}

func (obj *User) GetValue4DB(name string) interface{} {
	switch name {
	case "id":
		return obj.Id
	case "promotional_partner_id":
		return obj.PromotionalPartnerId
	case "mobile":
		return obj.Mobile
	case "present_chance":
		return obj.PresentChance
	case "shared":
		return obj.Shared
	case "locked":
		return obj.Locked
	}
	panic(errors.New("对象user属性[" + name + "]不存在"))
}
