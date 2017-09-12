package model

import "github.com/pkg/errors"

type RefundReason struct {
	Id          string `form:"id" json:"id" db:"id" desc:"id"`
	Priority    int    `form:"priority" json:"priority" db:"priority" desc:"优先权重"`
	Hide        bool   `form:"hide" json:"hide" db:"hide" desc:"隐藏"`
	Description string `form:"description" json:"description" db:"description" desc:"描述"`
}

func (obj *RefundReason) GetId() string {
	return obj.Id
}

func (obj *RefundReason) SetId(id string) {
	obj.Id = id
}

func (obj *RefundReason) GetTableName() string {
	return "tbl_refund_reason"
}

func (obj *RefundReason) GetDBNames() []string {
	return []string{"id","priority", "hide", "description"}
}

func (obj *RefundReason) GetMapNames() []string {
	return []string{"id","priority", "hide", "description"}
}

func (obj *RefundReason) GetValue4Map(name string) interface{} {
	switch name {
	case "id": return obj.Id
	case "priority": return obj.Priority
	case "hide": return obj.Hide
	case "description": return obj.Description
	}
	panic(errors.New("对象refundReason属性[" + name + "]不存在"))
}

func (obj *RefundReason) GetDBFromMapName(name string) string {
	str := GetDBFromMapName(obj, name)
	if str != "" {
		return str
	}
	panic(errors.New("对象refundReason属性[" + name + "]不存在"))
}

func (obj *RefundReason) GetPointer4DB(name string) interface{} {
	switch name {
	case "id": return &obj.Id
	case "priority": return &obj.Priority
	case "hide": return &obj.Hide
	case "description": return &obj.Description
	}
	panic(errors.New("对象refundReason属性[" + name + "]不存在"))
}

func (obj *RefundReason) GetValue4DB(name string) interface{} {
	switch name {
	case "id": return obj.Id
	case "priority": return obj.Priority
	case "hide": return obj.Hide
	case "description": return obj.Description
	}
	panic(errors.New("对象refundReason属性[" + name + "]不存在"))
}