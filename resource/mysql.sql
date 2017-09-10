drop table if exists tbl_admin;
create table tbl_admin(
id varchar(64) not null primary key,
name varchar(32) not null default '',
mobile varchar(32) not null default '',
password varchar(64) not null default '',
locked tinyint(1) default false,
del_flag tinyint(1) default false
)default charset=utf8;

drop table if exists tbl_role;
create table tbl_role(
id varchar(64) not null primary key,
name varchar(32) not null default '',
permission_list_str varchar(1000) not null default '',
description varchar(200) not null default '',
del_flag tinyint(1) default false
)default charset=utf8;

drop table if exists tr_admin_role;
create table tr_admin_role(
id varchar(64) not null primary key,
admin_id varchar(64) not null default '',
role_id varchar(64) not null default '',
del_flag tinyint(1) default false
)default charset=utf8;

drop table if exists tbl_advertisement;
create table tbl_advertisement(
id varchar(64) not null primary key,
name varchar(32) not null default '',
image varchar(200) not null default '',
link varchar(200) not null default '',
priority int(8) not null default 0,
hide tinyint(1) default false,
description varchar(200) not null default '',
del_flag tinyint(1) default false
)default charset=utf8;

drop table if exists tbl_shop;
create table tbl_shop(
id varchar(64) not null primary key,
name varchar(32) not null default '',
logo varchar(200) not null default '',
pre_image varchar(200) not null default '',
total_cash_coupon_number int(8) not null default 0,
total_cash_coupon_price int(8) not null default 0,
Introduction varchar(200) not null default '',
Address varchar(200) not null default '',
priority int(8) not null default 0,
hide tinyint(1) default false,
del_flag tinyint(1) default false
)default charset=utf8;

drop table if exists tbl_shop_account;
create table tbl_shop_account(
id varchar(64) not null primary key,
shop_id varchar(64) not null default '',
name varchar(32) not null default '',
password varchar(64) not null default '',
description varchar(200) not null default '',
money_chance int(8) not null default 0,
total_money int(8) not null default 0,
locked tinyint(1) default false,
del_flag tinyint(1) default false
)default charset=utf8;

drop table if exists tbl_shop_classification;
create table tbl_shop_classification(
id varchar(64) not null primary key,
name varchar(32) not null default '',
description varchar(200) not null default '',
hide tinyint(1) default false,
priority int(8) not null default 0,
del_flag tinyint(1) default false
)default charset=utf8;

drop table if exists tr_shop_classification_shop;
create table tr_shop_classification_shop(
id varchar(64) not null primary key,
shop_id varchar(64) not null default '',
shop_classification_id varchar(64) not null default '',
del_flag tinyint(1) default false
)default charset=utf8;

drop table if exists tbl_shop_introduction_image;
create table tbl_shop_introduction_image(
id varchar(64) not null primary key,
shop_id varchar(64) not null default '',
priority int(8) not null default 0,
content_url varchar(200) not null default '',
hide tinyint(1) default false,
del_flag tinyint(1) default false
)default charset=utf8;

drop table if exists tbl_cash_coupon;
create table tbl_cash_coupon(
id varchar(64) not null primary key,
name varchar(32) not null default '',
pre_image varchar(200) not null default '',
discount_amount int(8) not null default 0,
expiryDate timestamp DEFAULT CURRENT_TIMESTAMP,
expired tinyint(1) default false,
locked tinyint(1) default false,
del_flag tinyint(1) default false
)default charset=utf8;

drop table if exists tbl_user;
create table tbl_user(
id varchar(64) not null primary key,
promotional_partner_id varchar(64) not null default '',
mobile varchar(32) not null default '',
present_chance int(8) not null default 0,
shared tinyint(1) default false,
locked tinyint(1) default false,
del_flag tinyint(1) default false
)default charset=utf8;

drop table if exists tbl_cash_coupon_order;
create table tbl_cash_coupon_order(
id varchar(64) not null primary key,
user_id varchar(64) not null default '',
cash_coupon_id varchar(64) not null default '',
number varchar(200) not null default '',
status int(8) not null default 0,
del_flag tinyint(1) default false
)default charset=utf8;

drop table if exists tbl_refund;
create table tbl_refund(
id varchar(64) not null primary key,
cash_coupon_order_id varchar(64) not null default '',
evidence varchar(200) not null default '',
reason varchar(200) not null default '',
status int(8) not null default 0,
del_flag tinyint(1) default false
)default charset=utf8;

drop table if exists tbl_refund_reason;
create table tbl_refund_reason(
id varchar(64) not null primary key,
description varchar(200) not null default '',
del_flag tinyint(1) default false
)default charset=utf8;

drop table if exists tbl_present;
create table tbl_present(
id varchar(64) not null primary key,
name varchar(32) not null default '',
address varchar(200) not null default '',
stock int(8) not null default 0,
requirement int(8) not null default 0,
weight int(8) not null default 1,
expiryDate timestamp DEFAULT CURRENT_TIMESTAMP,
expired tinyint(1) default false,
locked tinyint(1) default false,
del_flag tinyint(1) default false
)default charset=utf8;

drop table if exists tbl_present_order;
create table tbl_present_order(
id varchar(64) not null primary key,
present_id varchar(64) not null default '',
number varchar(200) not null default '',
status int(8) not null default 0,
del_flag tinyint(1) default false
)default charset=utf8;

drop table if exists tbl_promotional_partner;
create table tbl_promotional_partner(
id varchar(64) not null primary key,
name varchar(32) not null default '',
description varchar(200) not null default '',
total_price int(8) not null default 0,
total_pay int(8) not null default 0,
del_flag tinyint(1) default false
)default charset=utf8;
