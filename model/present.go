package model

import "mimi/djq/util"

//id varchar(64) not null primary key,
//name varchar(32) not null default '',
//address varchar(200) not null default '',
//stock int(8) not null default 0,
//requirement int(8) not null default 0,
//ExpiryDate timestamp DEFAULT CURRENT_TIMESTAMP,
//expired tinyint(1) default false,
//locked tinyint(1) default false,
//del_flag tinyint(1) default false

type Present struct {
	Id          string        `json:"id" db:"id" desc:"id"`
	Name        string        `json:"name" db:"name" desc:"名称"`
	Address     string        `json:"address" db:"address" desc:"地址"`
	Stock       int           `json:"stock" db:"stock" desc:"库存"`
	Requirement int           `json:"requirement" db:"requirement" desc:"需求"`
	ExpiryDate  util.JSONTime `json:"expiryDate" db:"expiryDate" desc:"有效时间"`
	Expired     bool          `json:"expired" db:"expired" desc:"已过期"`
	Locked      bool          `json:"locked" db:"locked" desc:"锁定"`
}
