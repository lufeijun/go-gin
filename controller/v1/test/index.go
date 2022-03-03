package test

import (
	"fmt"
	"gin/structs/response"
	"net/http"
	"strconv"

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
