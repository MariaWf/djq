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
	log.Println("------开启服务：" + time.Now().String())
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