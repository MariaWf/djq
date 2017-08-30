package model

import "github.com/pkg/errors"

type User struct {
	//base
	Id       string `json:"id" db:"id" desc:"id"`
	Name     string `json:"name" db:"name" desc:"姓名"`
	NickName string `json:"nickName" db:"nick_name" desc:"昵称"`
	Mobile   string `json:"mobile" db:"mobile" desc:"手机"`
	Password string `json:"password" db:"password" desc:"密码"`
}

func (u *User) Convert2Map() map[string]interface{} {
	m := make(map[string]interface{})
	m["id"] = u.Id
	m["name"] = u.Name
	m["nickName"] = u.NickName
	m["mobile"] = u.Mobile
	m["password"] = u.Password
	return m
}

func (u *User) ConvertFromMap(m map[string]interface{}) {
	if m["id"] != nil {
		u.Id = m["id"].(string)
	}
	if m["name"] != nil {
		u.Name = m["name"].(string)
	}
	if m["mobile"] != nil {
		u.Mobile = m["mobile"].(string)
	}
	if m["nickName"] != nil {
		u.NickName = m["nickName"].(string)
	}
	if m["password"] != nil {
		u.Password = m["password"].(string)
	}
}

func (u *User) GetPointer4DB(name string) interface{} {
	switch name {
	case "id":return &u.Id;
	case "name":return &u.Name
	case "nick_name":return &u.NickName
	case "mobile":return &u.Mobile
	case "password":return &u.Password
	}
	panic(errors.New("属性不存在"))
}

func (u *User) GetPointers4DB(names []string) []interface{} {
	pointers := make([]interface{}, 0, 5)
	for _, name := range names {
		pointers = append(pointers, u.GetPointer4DB(name))
	}
	return pointers
}

func (u *User) GetValue4DB(name string) interface{} {
	switch name {
	case "id":return u.Id;
	case "name":return u.Name
	case "nick_name":return u.NickName
	case "mobile":return u.Mobile
	case "password":return u.Password
	}
	panic(errors.New("属性不存在"))
}

func (u *User) GetValues4DB(names []string) []interface{} {
	values := make([]interface{}, 0, 5)
	for _, name := range names {
		values = append(values, u.GetValue4DB(name))
	}
	return values
}