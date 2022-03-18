package apis

import (
	"gin/controller/v1/article"
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

		// redis
		testApi.GET("redis/set", test.RedisSet)
		testApi.GET("redis/get", test.RedisGet)
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

	// 文章部分
	articleV1 := api.Group("article/v1")
	{
		articleV1.POST("list", article.List)
		articleV1.POST("detail/:id", article.Detail)
		articleV1.POST("add", article.Add)
		articleV1.POST("update/:id", article.Update)

		// 类目
		articleCategoryV1 := articleV1.Group("category")
		{
			articleCategoryV1.POST("list", article.CategoryList)
			articleCategoryV1.POST("add", article.CategoryAdd)
			articleCategoryV1.POST("update/:id", article.CategoryUpdate)
			articleCategoryV1.POST("detail/:id", article.CategoryDetail)
		}

	}

}