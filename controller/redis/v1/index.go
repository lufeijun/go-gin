package v1

import (
	"context"
	"gin/database"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	ctx := context.Background()

	redisTool := database.GetRedisHelper()

	user := "test"

	name, err := redisTool.Set(ctx, "go", user, 10*time.Minute).Result()

	if err != nil {
		panic("缓存数据失败")
	}

	c.JSON(200, name)
}

func Get(c *gin.Context) {
	ctx := context.Background()

	redisTool := database.GetRedisHelper()

	name, err := redisTool.Get(ctx, "go").Result()

	// fmt.Println(name)
	// fmt.Println(err)

	if err != nil {
		panic("读取数据失败")
	}
	c.JSON(200, name)
}

func Setnx(c *gin.Context) {
	ctx := context.Background()

	redisTool := database.GetRedisHelper()

	name, err := redisTool.SetNX(ctx, "lock", 1, 10*time.Minute).Result()
	if err != nil {
		panic("读取数据失败")
	}

	c.JSON(200, name)

}
