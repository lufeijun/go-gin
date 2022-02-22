package test

import (
	"crypto/tls"
	"errors"
	"fmt"
	"gin/structs/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
)

func Error(c *gin.Context) {
	// var a int

	// a = 1
	// b, _ := strconv.Atoi(c.PostForm("b"))

	// d := a / b

	// fmt.Println(d)

	c.Error(errors.New("测试全局错误捕获"))

}

func Panic(c *gin.Context) {
	a := 1
	b, _ := strconv.Atoi(c.PostForm("b"))

	d := a / b

	fmt.Println(d)

	res := response.GetResponse()
	c.JSON(http.StatusOK, res)
}

// 发送邮件
func SendEmail(c *gin.Context) {
	m := gomail.NewMessage()
	m.SetHeader("From", "lufeijun_1234@126.com") //发送者
	m.SetHeader("To", "jipeng@zhufaner.com")     //接受者可以有多个
	body := "这是测试，这是演习"                          //发送内容
	m.SetHeader("Subject", "演习演习演习")             // 邮件标题
	m.SetBody("text/html", body)                 // 发送邮件内容
	d := gomail.NewDialer("smtp.126.com", 465, "lufeijun_1234@126.com", "JP418955279jp")
	//这里第一个参数为服务器地址，第二个为端口号，第三个为发送者邮箱号
	//第四个如果是qq邮箱为授权玛而其他邮箱是密码
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("***%s\n", err.Error())
	}

	res := response.GetResponse()
	c.JSON(http.StatusOK, res)
}
