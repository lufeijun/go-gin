package v1

import (
	"fmt"
	"gin/config"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {

	fmt.Println(config.APP_NAME)

	fmt.Println("listcontroller")
}
