package model

import (
	"mimi/djq/util"
)

type CashCoupon struct {
	Id         string        `json:"id" db:"id" desc:"id"`
	Name       string        `json:"name" db:"name" desc:"名称"`
	PreImage   string        `json:"preImage" db:"pre_image" desc:"预览图"`
	ExpiryDate util.JSONTime `json:"expiryDate" db:"expiry_date" desc:"有效日期" time_format:"2006-01-02" time_utc:"1"`
}
