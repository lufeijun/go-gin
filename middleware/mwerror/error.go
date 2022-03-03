package mwerror

import (
	"gin/structs/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// err 处理中间件，
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 先调用c.Next()执行后面的中间件

		// 所有中间件及router处理完毕后从这里开始执行
		// 检查c.Errors中是否有错误
		for _, e := range c.Errors {
			err := e.Err
			res := response.GetResponse()
			res.SetMessage(err.Error())

			c.JSON(http.StatusOK, res)
			return // 检查一个错误就行
		}
	}
}
