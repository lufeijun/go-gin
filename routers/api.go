package routers

import (
	controllerArticleV1 "gin/controller/article/v1"

	controllerKafkaV1 "gin/controller/kafka/v1"

	"github.com/gin-gonic/gin"
)

func LoadApi(e *gin.Engine) {

	api := e.Group("api")

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

}
