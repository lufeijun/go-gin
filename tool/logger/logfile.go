package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

func WriteInLog(msg string) {

	dir := "./logs/"
	filename := time.Now().Format("2006-01-02") + ".log"

	logFile, err := os.OpenFile(dir+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}

	log.SetOutput(logFile)

	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println(msg)
}
