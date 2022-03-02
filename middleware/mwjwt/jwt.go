package mwjwt

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("jwt middleware")
		// 验证通过
		c.Next()

	}
}
