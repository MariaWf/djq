package main

import (
	"mimi/djq/router"
	"os"
	"log"
	"mimi/djq/config"
	"path/filepath"
	"time"
)
func main() {
	initLog()
	log.Println("------LstdFlags：" + time.Now().String())
	log.SetFlags(log.LstdFlags|log.Llongfile)
	//log.Println(1)
	//log.SetFlags(log.Ldate)
	//log.Println(2)
	//log.SetFlags(log.Llongfile)
	//log.Println(3)
	//log.SetFlags(log.Lmicroseconds)
	//log.Println(4)
	//log.SetFlags(log.Lshortfile)
	//log.Println(5)
	//log.SetFlags(log.LstdFlags)
	//log.Println(6)
	//log.SetFlags(log.Ltime)
	//log.Println(7)
	//log.SetFlags(log.LUTC)
	//log.Println(8)
	router.Begin()
}

func initLog(){
	if "false" == config.Get("output_log"){
		return
	}
	globalLogUrl := config.Get("global_log")
	if globalLogUrl == "" {
		globalLogUrl = "global.log"
	} else {
		path := filepath.Dir(globalLogUrl)
		os.MkdirAll(path, 0777)
	}
	logFile, err := os.OpenFile(globalLogUrl, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err!=nil{
		panic(err)
	}
	log.SetOutput(logFile)
	log.Println()
}