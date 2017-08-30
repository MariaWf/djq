package model

type Advertisement struct {
	Id          string `json:"id" db:"id" desc:"id"`
	Name        string `json:"name" db:"name" desc:"名称"`
	PreImage    string `json:"pre_image" db:"name" desc:"预览图"`
	Link        string `json:"link" db:"name" desc:"超链接"`
	Priority    int  `json:"priority" db:"priority" desc:"优先权重"`
	Description string `json:"description" db:"description" desc:"描述"`
}

