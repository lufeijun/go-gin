package test

import (
	"fmt"
	"gin/structs/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 协程部分
// 问题
// 1、开启协程后，如果协程抛出异常，会导致整个服务停掉
// 2、异常处理中间件不会 handle 协程里边的异常。需要在携程函数里边处理异常

// 简单开启协程
func AsyncOne(c *gin.Context) {

	res := response.GetResponse()
	c.JSON(http.StatusOK, res)

	fmt.Println(c.Request.PostForm)
	fmt.Println(c.PostForm("a"))
	fmt.Println(c.Request.PostForm)

	// 同步
	// time.Sleep(5 * time.Second)
	// fmt.Println("done in path " + c.Request.URL.Path)

	// 异步
	cCp := c.Copy()
	// fmt.Println(c.PostForm("a"))
	go func() {

		defer func() {
			if err := recover(); err != nil {
				// 打印异常，关闭资源，退出此函数
				fmt.Println(err)
			}
		}()

		time.Sleep(5 * time.Second)
		fmt.Println("done in path " + cCp.PostForm("a"))
		panic("测试")
	}()

	return
}
