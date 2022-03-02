package mwsession

import (
	"gin/config"
	"gin/structs/response"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// var nocheckUri = []string{
// 	"/api/manager/v1/api/login",
// 	"/api/manager/v1/api/logout",
// }

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// fmt.Println("===中间件==start===")

		// 是否需要 session 验证
		if isNeedCheck(c.Request.URL.String()) {
			session := sessions.Default(c)

			// 查看 session 登录信息
			sessionBytes := session.Get(config.SESSION_KEY)
			if sessionBytes == nil {
				res := response.GetResponse()
				res.SetMessage("请登录")
				c.JSON(http.StatusOK, res)
				c.Abort()
			}
		} // end 是否需要 session 验证

		// 验证通过
		c.Next()

	}
}

//
func isNeedCheck(uri string) (ok bool) {
	ok = true

	if uri == "/api/manager/v1/api/login" {
		ok = false
	} else if uri == "/api/manager/v1/api/logout" {
		ok = false
	}

	return
}
