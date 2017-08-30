package model

type ShopClassification struct {
	Id       string `json:"id" db:"id" desc:"id"`
	Name     string `json:"name" db:"name" desc:"名称"`
	Description string `json:"description" db:"description" desc:"描述"`
}
