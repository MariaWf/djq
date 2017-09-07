package service

import (
	"mimi/djq/dao/arg"
	"github.com/pkg/errors"
	"mimi/uservice/rpc/reply"
)
var ErrLoginMismatch = errors.New("service: login mismatch")

type User int

func (t *User) Find(args *arg.User, reply *reply.Objects) error {
	//conn, _ := mysql.Get()
	//defer mysql.Close(conn)
	//userDao := &dao.UserDao{conn}
	//users, err := userDao.Find(args)
	////reply.Objects = users
	//reply.Err = err
	return nil
}

//func (t *User) FindWithoutPassword(args *arg.User, reply *reply.Objects) error {
//	conn, _ := mysql.Get()
//	defer mysql.Close(conn)
//	userDao := &dao.UserDao{conn}
//	user := new(model.User)
//	user.NickName = "mimi"
//	user.Mobile = "11111111111"
//	user.Password = "123123"
//	return nil
//}
//
//func (t *User) getWithoutPassword(args *Args, reply *Reply) error {
//	reply.C = args.A * args.B
//	return nil
//}

func (t *User) Login(args *arg.User, reply *reply.Normal) error {
	//conn, _ := mysql.Get()
	//defer mysql.Close(conn)
	//userDao := &dao.UserDao{conn}
	//users, err := userDao.Find(args)
	//if err != nil {
	//	reply.Err = errors.Wrap(err, "userDao.Find")
	//}else if len(users) ==0{
	//	reply.Err = ErrLoginMismatch
	//}else{
	//	reply.Ok = true
	//}
	return nil
}