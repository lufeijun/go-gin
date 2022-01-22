package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func CookieSet(c *gin.Context) {

	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, 1)

	// 方法一
	cookie := http.Cookie{Name: "user-name", Value: "lufeijun1234", Expires: expiration}

	http.SetCookie(c.Writer, &cookie)
	// c.SetCookie()

	// 方法二
	c.SetCookie("user-name-2", "cookievalue", 3600, "/", "", false, true)

	c.JSON(200, 124)
}

func CookieGet(c *gin.Context) {
	value, _ := c.Cookie("user-name-2")

	for _, cookie := range c.Request.Cookies() {
		fmt.Println(cookie.Name, " 对应的值 ", cookie.Value)
	}

	// c.JSON(200, value)
	c.String(200, value)

}

func SessionSet(c *gin.Context) {

	//
	// store := sessions.NewCookieStore([]byte("session-key"))

	store, _ := redis.NewStoreWithDB(10, "tcp", "127.0.0.1:6379", "123456", "10", []byte("session-key"))

	session, _ := store.Get(c.Request, "session-name123")

	session.Values["name"] = "吉鹏"
	session.Values["age"] = "12"

	session.Options.MaxAge = 200

	// session.Options.Path

	err := session.Save(c.Request, c.Writer)
	if err != nil {
		panic("出错啦")
	}

	c.String(200, "sdasds")
}

func SessionGet(c *gin.Context) {
	// store := sessions.NewCookieStore([]byte("session-key"))

	store, _ := redis.NewStoreWithDB(10, "tcp", "127.0.0.1:6379", "123456", "10", []byte("session-key"))

	session, _ := store.Get(c.Request, "session-name123")

	fmt.Println(session)

	name := session.Values["name"]

	if name == nil {
		c.String(200, "过期了")
		return
	}

	age := session.Values["age"]

	// var aaa map[string]interface{}
	// aaa = make(map[string]interface{})
	// aaa["name"] = name
	// aaa["age"] = age

	var aaa map[string]string
	aaa = make(map[string]string)

	aaa["name"] = name.(string)
	aaa["age"] = age.(string)

	c.JSON(200, aaa)
}
