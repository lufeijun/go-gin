package mwsession

import (
	"fmt"
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

	if uri == "/api/manager/v1/login" {
		ok = false
	} else if uri == "/api/manager/v1/logout" {
		ok = false
	}

	return
}

// 当前问题。session 值保存在 redis 中，按照 redis 的过期时间来检测，理论上来说，在每次请求后，需要自动将 redis 对应值的过期时间延长
// 否则，在前端在操作的时候，可能会突然间退出系统
// 原理，获取 session 的 id，依据id更新对应的 redis 的值
func resetSessionTime(session sessions.Session) {
	id := session.ID() // 这个值为 B3R3ALZ3HWRVOSBCAB2CSBWDJHQQUIFJFTRXEJ7LYBMQCPY3MDMQ
	// 对应的 redis 中 key 为 session_B3R3ALZ3HWRVOSBCAB2CSBWDJHQQUIFJFTRXEJ7LYBMQCPY3MDMQ
	// 即有一个前缀，依据这个前缀更新 redis

	fmt.Println(id)

}
