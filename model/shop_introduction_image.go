package model

import "github.com/pkg/errors"

type ShopIntroductionImage struct {
	Id         string `form:"id" json:"id" db:"id" desc:"id"`
	ShopId     string `form:"shopId" json:"shopId" db:"shop_id" desc:"shopId"`
	Priority   int    `form:"priority" json:"priority" db:"priority" desc:"优先权重"`
	ContentUrl string `form:"contentUrl" json:"contentUrl" db:"content_url" desc:"内容路径"`
	Hide       bool   `form:"hide" json:"hide" db:"hide" desc:"隐藏"`
}

func (obj *ShopIntroductionImage) GetId() string {
	return obj.Id
}

func (obj *ShopIntroductionImage) SetId(id string) {
	obj.Id = id
}

func (obj *ShopIntroductionImage) GetTableName() string {
	return "tbl_shop_introduction_image"
}

func (obj *ShopIntroductionImage) GetDBNames() []string {
	return []string{"id", "shop_id", "priority", "content_url", "hide"}
}

func (obj *ShopIntroductionImage) GetMapNames() []string {
	return []string{"id", "shopId", "priority", "contentUrl", "hide"}
}

func (obj *ShopIntroductionImage) GetValue4Map(name string) interface{} {
	switch name {
	case "id":
		return obj.Id
	case "shopId":
		return obj.ShopId
	case "priority":
		return obj.Priority
	case "contentUrl":
		return obj.ContentUrl
	case "hide":
		return obj.Hide
	}
	panic(errors.New("对象shopIntroductionImage属性[" + name + "]不存在"))
}

func (obj *ShopIntroductionImage) GetDBFromMapName(name string) string {
	str := GetDBFromMapName(obj, name)
	if str != "" {
		return str
	}
	panic(errors.New("对象shopIntroductionImage属性[" + name + "]不存在"))
}

func (obj *ShopIntroductionImage) GetPointer4DB(name string) interface{} {
	switch name {
	case "id":
		return &obj.Id
	case "shop_id":
		return &obj.ShopId
	case "priority":
		return &obj.Priority
	case "content_url":
		return &obj.ContentUrl
	case "hide":
		return &obj.Hide
	}
	panic(errors.New("对象shopIntroductionImage属性[" + name + "]不存在"))
}

func (obj *ShopIntroductionImage) GetValue4DB(name string) interface{} {
	switch name {
	case "id":
		return obj.Id
	case "shop_id":
		return obj.ShopId
	case "priority":
		return obj.Priority
	case "content_url":
		return obj.ContentUrl
	case "hide":
		return obj.Hide
	}
	panic(errors.New("对象shopIntroductionImage属性[" + name + "]不存在"))
}
