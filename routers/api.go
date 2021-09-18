package routers

import (
	"github.com/gin-gonic/gin"

	controllerArticleV1 "gin/controller/article/v1"
)

func LoadApi(e *gin.Engine) {

	api := e.Group("api")

	articleV1 := api.Group("article/v1")
	{
		articleV1.GET("list", controllerArticleV1.List)
		articleV1.POST("add", controllerArticleV1.Add)
		articleV1.POST("update", controllerArticleV1.Update)
	}

}
