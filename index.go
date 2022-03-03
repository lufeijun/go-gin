package main

import (
	"gin/config"
	"gin/cron"
	"gin/middleware/mwerror"
	"gin/middleware/mwrecover"
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
	r := gin.New()

	// 两个基础的中间件，处理 err 和 panic
	r.Use(mwerror.ErrorHandler())
	r.Use(mwrecover.RecoverHandler())

	// 路由 登录控制由 session 来控制
	routers.InitRouter(r)

	// 3.监听端口，默认在8080
	err := r.Run(":" + config.APP_PORT)
	if err != nil {
		panic("启动失败：err=" + err.Error())
	}
}
