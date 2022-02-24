package logger

import (
	"log"
	"os"
	"time"
)

// go 系统自带的
func WriteInLog(msg string) {

	dir := "./logs/"
	filename := time.Now().Format("2006-01-02") + ".log"

	logFile, err := os.OpenFile(dir+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic("日志写入失败, err:" + err.Error())
	}

	log.SetOutput(logFile)

	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println(msg)
}
