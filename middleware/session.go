package middleware

import (
	"gin/config"
	"gin/structs/response"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		// fmt.Println("===中间件==start===")

		// 查看 session 登录信息
		sessionBytes := session.Get(config.SESSION_KEY)
		if sessionBytes == nil {
			res := response.GetResponse()
			res.SetMessage("请登录")
			c.JSON(http.StatusOK, res)
			c.Abort()
		}

		// 验证通过
		c.Next()

	}
}
