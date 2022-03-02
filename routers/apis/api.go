package apis

import (
	"gin/controller/v1/manager"
	"gin/controller/v1/test"
	"gin/controller/v2/testv2"

	"github.com/gin-gonic/gin"
)

// 实际写路由的地方
func LoadApi(api *gin.RouterGroup) {
	// api := e.Group("api")

	// 测试部分 v1 版
	testApi := api.Group("test/v1")
	{
		testApi.GET("index", test.Index)
	}
	// v2 版
	testApiV2 := api.Group("test/v2")
	{
		testApiV2.GET("index", testv2.Index)
	}
	// 登录接口
	// e.POST("api/login", controllerManager.Login)
	// e.POST("api/logout", controllerManager.Logout)

	// 管理岗部分
	managerV1 := api.Group("manager/v1")
	{
		managerV1.POST("api/login", manager.Login)
		// managerV1.POST("add", controllerManager.Add)
	}
}
