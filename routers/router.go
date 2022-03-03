package routers

import (
	"gin/config"
	"gin/controller/v1/test"
	"gin/middleware/mwjwt"
	"gin/middleware/mwsession"
	"gin/routers/apis"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// 将这个文件作为路由文件入口
func InitRouter(e *gin.Engine) {

	api := e.Group("api")

	//
	initSessionApi(api)

	//
	// initJwtApi(api)
}

// 说明：
//   apis.LoadApi 这个方法也可放到 InitRouter 或者 initSessionApi|initJwtApi 方法中，
//   1、如果放到 InitRouter ，在 InitRouter 中可以同时调用 initSessionApi 或者 initJwtApi 方法
//   2、如果放到 initSessionApi 和 initJwtApi 方法中，则在 InitRouter 中，initSessionApi 和 initJwtApi 只能加载一个，否则路由冲突
//
// 这里选择 2 ，apis.LoadApi 放到 init* 方法中

// api 路由部分由 session 控制
func initSessionApi(api *gin.RouterGroup) {

	// 子路由，这样就不影响其他路由的中间件设置
	sessionApi := api.Group("")

	// redis session
	store, _ := redis.NewStoreWithDB(config.SESSION_REDIS_CONNECTIONS, "tcp", config.SESSION_REDIS_ADDR, config.SESSION_REDIS_PASS, config.SESSION_REDIS_DB, []byte("secret"))
	sessionApi.Use(sessions.Sessions(config.SESSION_KEY, store))

	// 引用中间件
	sessionApi.Use(mwsession.Login())

	// 加载路由 ,
	apis.LoadApi(sessionApi)

}

// 这部分接口由 jwt 中间件控制
func initJwtApi(api *gin.RouterGroup) {

	jwtApi := api.Group("")

	// 引用中间件
	jwtApi.Use(mwjwt.Login())

	jwtApi.GET("jwt", test.Jwt)

	// 加载路由
	apis.LoadApi(jwtApi)
}
