package model


type Admin struct {
	//base
	Id       string `json:"id" db:"id" desc:"id"`
	Name     string `json:"name" db:"name" desc:"名称"`
	Mobile   string `json:"mobile" db:"mobile" desc:"手机"`
	Password string `json:"password" db:"password" desc:"密码"`
}
