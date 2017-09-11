package model

import "github.com/pkg/errors"

type Advertisement struct {
	Id          string `form:"id" json:"id" db:"id" desc:"id"`
	Name        string `form:"name" json:"name" db:"name" desc:"名称"`
	Image       string `form:"image" json:"image" db:"image" desc:"图片"`
	Link        string `form:"link" json:"link" db:"name" desc:"超链接"`
	Priority    int    `form:"priority" json:"priority" db:"priority" desc:"优先权重"`
	Hide        bool   `form:"hide" json:"hide" db:"hide" desc:"隐藏"`
	Description string `form:"description" json:"description" db:"description" desc:"描述"`
}

func (obj *Advertisement) GetId() string {
	return obj.Id
}

func (obj *Advertisement) SetId(id string) {
	obj.Id = id
}

func (obj *Advertisement) GetTableName() string {
	return "tbl_advertisement"
}

func (obj *Advertisement) GetDBNames() []string {
	return []string{"id", "name", "image", "link", "priority", "hide", "description"}
}

func (obj *Advertisement) GetMapNames() []string {
	return []string{"id", "name", "image", "link", "priority", "hide", "description"}
}

func (obj *Advertisement) GetValue4Map(name string) interface{} {
	switch name {
	case "id": return obj.Id
	case "name": return obj.Name
	case "image": return obj.Image
	case "link": return obj.Link
	case "priority": return obj.Priority
	case "hide": return obj.Hide
	case "description": return obj.Description
	}
	panic(errors.New("对象advertisement属性[" + name + "]不存在"))
}

func (obj *Advertisement) GetDBFromMapName(name string) string {
	str := GetDBFromMapName(obj, name)
	if str != "" {
		return str
	}
	panic(errors.New("对象advertisement属性[" + name + "]不存在"))
}

func (obj *Advertisement) GetPointer4DB(name string) interface{} {
	switch name {
	case "id": return &obj.Id
	case "name": return &obj.Name
	case "image": return &obj.Image
	case "link": return &obj.Link
	case "priority": return &obj.Priority
	case "hide": return &obj.Hide
	case "description": return &obj.Description
	}
	panic(errors.New("对象advertisement属性[" + name + "]不存在"))
}

func (obj *Advertisement) GetValue4DB(name string) interface{} {
	switch name {
	case "id": return obj.Id
	case "name": return obj.Name
	case "image": return obj.Image
	case "link": return obj.Link
	case "priority": return obj.Priority
	case "hide": return obj.Hide
	case "description": return obj.Description
	}
	panic(errors.New("对象advertisement属性[" + name + "]不存在"))
}