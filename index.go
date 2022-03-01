package main

import (
	"gin/config"
	"gin/cron"
	"gin/routers"
	"gin/tool/logger"

	"github.com/gin-gonic/gin"
)

// 初始化一写东西
func init() {

	// 日志函数
	logger.InitZapLogger()

	// 定时任务
	if config.APP_IS_CRON {
		cron.InitCron()
	}
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
	err := r.Run(":" + config.APP_PORT)

	if err != nil {

		panic("启动失败：err=" + err.Error())
	}
}
