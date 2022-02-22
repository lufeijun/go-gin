package middleware

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"gin/config"
	"gin/structs/response"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
)

func RecoverHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录日志
				sendToLog(c, fmt.Errorf("%v", err).Error())

				// 发送错误邮件
				sendToEmail(c, fmt.Errorf("%v", err).Error())

				res := response.GetResponse()
				res.SetMessage("系统错误")
				c.JSON(http.StatusOK, res)
				c.Abort()
			}
		}()

		c.Next()
	}
}

// 记录日志
func sendToLog(c *gin.Context, msg string) {

}

//
func sendToEmail(c *gin.Context, msg string) {

	// 加个开关
	if !config.MAIL_IS_SEND {
		return
	}

	// 构建邮箱内容
	var body string
	body += "<h1>错误信息：" + msg + "</h1>"

	body += "<div>路由=" + c.Request.URL.String() + "</div>"
	body += "<div>方法=" + c.Request.Method + "</div>"
	parameter, _ := json.Marshal(c.Request.PostForm)
	body += "<div>参数= " + string(parameter) + "</div>"

	body += "<h2> 错误栈： </h2>"
	for _, v := range strings.Split(string(debug.Stack()), "\n") {
		body += v + "<br>"
	}

	// 实际发送邮件代码
	m := gomail.NewMessage()
	m.SetAddressHeader("From", config.MAIL_FROM, config.MAIL_NAME) //发送者

	m.SetHeader("To", "jipeng@zhufaner.com") //接受者可以有多个

	// 标题
	subject := "【" + config.APP_NAME + "bug】"
	m.SetHeader("Subject", subject) // 邮件标题

	// 发邮件的内容
	m.SetBody("text/html", body)

	d := gomail.NewDialer(config.MAIL_HOST, config.MAIL_PORT, config.MAIL_FROM, config.MAIL_PASSWORD)
	//这里第一个参数为服务器地址，第二个为端口号，第三个为发送者邮箱号
	//第四个如果是qq邮箱为授权玛而其他邮箱是密码
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		// fmt.Printf("***%s\n", )
		panic(err.Error())
	}

}
