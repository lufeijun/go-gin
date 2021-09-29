package main

import (
	"gin/routers"

	mylog "gin/tool/logger"

	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {

	// 1.创建路由
	// r := gin.Default()
	r := gin.New()

	routers.LoadApi(r)

	// 3.监听端口，默认在8080
	err := r.Run(":8008")

	if err != nil {

		mylog.WriteInLog("启动失败：err=" + err.Error())
		return
	}
}
