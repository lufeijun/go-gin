package main

import (
	"gin/config"
	"gin/routers"

	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {

	// 设置运行模式
	if config.APP_MODE == "server" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 1.创建路由
	// r := gin.Default()
	r := gin.New()

	routers.LoadApi(r)

	// 3.监听端口，默认在8080
	err := r.Run(":8008")

	if err != nil {

		panic("启动失败：err=" + err.Error())
	}
}
