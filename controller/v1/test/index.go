package test

import (
	"context"
	"fmt"
	"gin/database/redis"
	"gin/structs/response"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	res := response.GetResponse()

	res.SetData("v1")

	c.JSON(http.StatusOK, res)

	return
}

func Jwt(c *gin.Context) {
	res := response.GetResponse()

	res.SetData("jwt")

	c.JSON(http.StatusOK, res)

	return
}

// 测试异常发生邮件
func Panic(c *gin.Context) {
	a := 1
	b, _ := strconv.Atoi(c.PostForm("b"))

	d := a / b

	fmt.Println(d)

	res := response.GetResponse()
	c.JSON(http.StatusOK, res)
}

// redis

func RedisSet(c *gin.Context) {
	res := response.GetResponse()
	ctx := context.Background()

	redis := redis.GetRedisHelper()

	name, err := redis.Set(ctx, "name", "吉鹏123", 10*time.Minute).Result()

	if err != nil {
		res.SetMessage("写入失败")
	} else {
		res.SetData(name)
	}

	c.JSON(http.StatusOK, res)
	return
}

func RedisGet(c *gin.Context) {
	res := response.GetResponse()

	ctx := context.Background()

	redisTool := redis.GetRedisHelper()

	name, err := redisTool.Get(ctx, "name").Result()

	if err != nil {
		res.SetMessage(err.Error())
	} else {
		res.SetData(name)
	}
	c.JSON(http.StatusOK, res)
}

// 切片
func SliceOne(c *gin.Context) {

	intslice := make([]int, 5)

	fmt.Println("长度===", len(intslice))

	intslice = append(intslice, 10)

	fmt.Println("长度===", len(intslice))

	// intslice[11] = 100  报错
	// fmt.Println("长度===", len(intslice))

	for index, v := range intslice {
		fmt.Println(index, "===", v)
	}

	c.String(http.StatusOK, "ok")
}
