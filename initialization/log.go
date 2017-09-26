package initialization

import (
	"path/filepath"
	"os"
	"mimi/djq/config"
	"log"
)

func InitGlobalLog() {
	if "false" == config.Get("output_log") {
		return
	}
	globalLogUrl := config.Get("global_log")
	if globalLogUrl == "" {
		globalLogUrl = "logs/global.log"
	}
	path := filepath.Dir(globalLogUrl)
	os.MkdirAll(path, 0777)
	logFile, err := os.OpenFile(globalLogUrl, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.Println()
}
