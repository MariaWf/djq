package dao

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

type interface1 interface {
	Method1()
	Method2()
}

type struct2 struct {
	interface1
	Arg1 int
}

type struct1 struct {
	struct2
}

func (obj *struct1) Method1() {
	fmt.Println("Method1")
}

func (obj *struct1) Method2() {
	fmt.Println("Method2")
}

func outp(obj interface1) {
	obj.Method1()
}

func TestAdminDao_Find(t *testing.T) {
	//obj := &struct1{}
	//outp(obj)

	//url1 := "http://www.51zxiu.cn"
	url2 := "http://www.51zxiu.cn/decoration"
	//url3 := "http://www.h0758.net/decoration"
	rep, err := http.Get(url2)
	if err != nil {
		t.Error(err)
	}
	//log.Println()
	fmt.Println(rep.Status)
	rep.Write(os.Stdout)
	//conn, err := mysql.Get()
	//if err != nil {
	//	t.Error(err)
	//}
	//defer mysql.Rollback(conn)
	//dao := &AdminDao{conn}
	//for i := 0; i < 10; i++ {
	//	admin := &model.Admin{}
	//	admin.Name = "name" + strconv.Itoa(i)
	//	admin.Password = "password" + strconv.Itoa(i)
	//	admin.Locked = i % 2 == 0
	//	_, err := dao.Add(admin)
	//	if err != nil {
	//		t.Error(err)
	//	}
	//}
	//arg := &arg.Admin{}
	//arg.LockedOnly = true
	//list, err := dao.Find(arg)
	//if err != nil {
	//	t.Error(err)
	//}
	//for _, admin := range list {
	//	t.Log(admin.Name)
	//}
}
