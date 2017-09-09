package mysql

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestClose(t *testing.T) {
	err := test1()
	if err != nil {
		t.Error(err)
	} else {
		t.Log("ok")
	}
}

func test1() error {
	conn, err := Get()
	if err != nil {
		return err
	}
	rollback := false
	defer tempClose(conn, &rollback)
	row, err := conn.Query("select * from tbl_role")
	if err != nil {
		return err
	}
	defer row.Close()
	//rollback = true
	//Rollback(conn)
	//panic(errors.New("test"))
	////Rollback(conn)
	//defer panic(errors.New("test2"))
	return errors.New("test")
}

func tempClose(tx *sql.Tx, rollback *bool) {
	if err := recover(); err != nil {
		tx.Rollback()
		fmt.Println("1")
		panic(err)
	} else if *rollback {
		tx.Rollback()
		fmt.Println("2")
	} else {
		tx.Commit()
		fmt.Println("3")
		//panic(errors.New("test3"))
	}
}

func BenchmarkBuildPassword4DB(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		test1()
	}
}
