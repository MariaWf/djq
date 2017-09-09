package model

type ShopIntroductionImage struct {
	Id         string `json:"id" db:"id" desc:"id"`
	Priority   int    `json:"priority" db:"priority" desc:"优先权重"`
	ContentUrl string `json:"contentUrl" db:"content_url" desc:"内容路径"`
}
