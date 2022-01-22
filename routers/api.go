package routers

import (
	"gin/config"
	controllerArticleV1 "gin/controller/article/v1"
	"gin/middleware"

	controllerKafkaV1 "gin/controller/kafka/v1"
	controllerManager "gin/controller/manager/v1"
	controllerRedis "gin/controller/redis/v1"
	controllerSession "gin/controller/session/v1"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func LoadApi(e *gin.Engine) {

	// redis session
	store, _ := redis.NewStoreWithDB(config.SESSION_REDIS_CONNECTIONS, "tcp", config.SESSION_REDIS_ADDR, config.SESSION_REDIS_PASS, config.SESSION_REDIS_DB, []byte("secret"))
	e.Use(sessions.Sessions(config.SESSION_KEY, store))

	api := e.Group("api")

	// 登录接口
	e.POST("api/login", controllerManager.Login)
	e.POST("api/logout", controllerManager.Logout)

	// 中间件
	api.Use(middleware.Login())

	// 管理岗部分
	managerV1 := api.Group("manager/v1")
	{
		managerV1.POST("add", controllerManager.Add)
	}

	articleV1 := api.Group("article/v1")
	{
		articleV1.GET("list", controllerArticleV1.List)
		articleV1.GET("detail", controllerArticleV1.Detail)
		articleV1.POST("add", controllerArticleV1.Add)
		articleV1.POST("update", controllerArticleV1.Update)
	}

	// kafka 部分
	kafka := api.Group("kafka/v1")
	{
		kafka.GET("index", controllerKafkaV1.Index)
		kafka.GET("producer", controllerKafkaV1.Producer)
		kafka.GET("consumer", controllerKafkaV1.Consumer)
	}

	// redis
	redis := api.Group("redis/v1")
	{
		redis.GET("index", controllerRedis.Index)
		redis.GET("get", controllerRedis.Get)
		redis.GET("setnx", controllerRedis.Setnx)
	}

	// session
	session := api.Group("session/v1")
	{
		session.GET("cookie/set", controllerSession.CookieSet)
		session.GET("cookie/get", controllerSession.CookieGet)
		session.GET("session/set", controllerSession.SessionSet)
		session.GET("session/get", controllerSession.SessionGet)
	}

}
