package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("111")

		// c.String(403, "dsadas")
		// c.Abort()

		c.Next()

		fmt.Println("222")

	}
}
