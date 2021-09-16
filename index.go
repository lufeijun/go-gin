package main

import (
	"fmt"
	"gin/config"
	"gin/routers"

	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {

	fmt.Println("======")
	fmt.Println(config.DbHost)
	fmt.Println("======")
	// 1.创建路由
	r := gin.Default()

	routers.LoadApi(r)

	// 3.监听端口，默认在8080
	err := r.Run(":8000")

	if err != nil {
		fmt.Println("启动失败：err=%v\n", err)
	}
}
