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

		kafkatest := testApi.Group("kafka")
		{
			kafkatest.GET("producer", test.KafkaProducer)
			kafkatest.GET("consumer", test.KafkaConsumer)
		}

		// 协程
		asynctest := testApi.Group("async")
		{
			asynctest.POST("one", test.AsyncOne)
		}

		// 数据库测试
		mysqltest := testApi.Group("mysql")
		{
			mysqltest.POST("one", test.MysqlOne)
			mysqltest.POST("two", test.Mysql2)
			mysqltest.POST("three", test.Mysql3)
		}

		clickhouseTest := testApi.Group("clickhouse")
		{
			clickhouseTest.POST("one", test.ClickhouseOne)
			clickhouseTest.POST("two", test.ClickhouseTwo)
			clickhouseTest.POST("wechat_user_sessions", test.Wechat_user_sessions)
			clickhouseTest.POST("wechat_user_sessions_list", test.Wechat_user_sessions_list)
		}

		// 爬虫
		anjuke := testApi.Group("anjuke")
		{
			anjuke.GET("index", test.Anjuke)
		}

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
