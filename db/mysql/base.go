package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"mimi/djq/config"
	"sync"
)

var db *sql.DB
var once sync.Once

func Get() (*sql.Tx, error) {
	var err error
	once.Do(func() {
		db, err = sql.Open("mysql", config.Get("mysqlDataSourceName"))
		if err != nil {
			panic(err)
		}
	})
	return db.Begin()
}

func Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}

func Close(tx *sql.Tx, rollback *bool) {
	if err := recover(); err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			log.Println(err2)
		}
		panic(err)
	} else if *rollback {
		checkErr(tx.Rollback())
	} else {
		checkErr(tx.Commit())
	}
}

func GetStatus() sql.DBStats {
	return db.Stats()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
