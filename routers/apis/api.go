package apis

import (
	"gin/controller/v1/manager"
	"gin/controller/v1/test"
	"gin/controller/v2/testv2"

	"github.com/gin-gonic/gin"
)

// 实际写路由的地方
func LoadApi(api *gin.RouterGroup) {

	// 测试部分 v1 版
	testApi := api.Group("test/v1")
	{
		testApi.GET("index", test.Index)
		testApi.GET("panic", test.Panic)
	}
	// v2 版
	testApiV2 := api.Group("test/v2")
	{
		testApiV2.GET("index", testv2.Index)
	}

	// 管理岗部分
	managerV1 := api.Group("manager/v1")
	{
		managerV1.POST("login", manager.Login)
		managerV1.POST("logout", manager.Logout)
		managerV1.POST("add", manager.Add)
	}
}
