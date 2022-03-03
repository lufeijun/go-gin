package manager

import (
	"gin/structs/response"
	"gin/ultility/v1/manager"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	res := response.GetResponse()

	email := c.PostForm("email")
	pwd := c.PostForm("pwd")

	if email == "" {
		res.SetMessage("email 不能为空")
	} else if pwd == "" {
		res.SetMessage("pwd 不能为空")
	} else {
		err := manager.Login(email, pwd, c)

		if err != nil {
			res.SetMessage(err.Error())
		}
	}

	c.JSON(http.StatusOK, res)
	return
}

func Logout(c *gin.Context) {

	res := response.GetResponse()

	manager.Logut(c)

	c.JSON(http.StatusOK, res)
	return
}

// 添加管理员
func Add(c *gin.Context) {
	res := response.GetResponse()

	name := c.PostForm("name")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	pwd := c.PostForm("pwd")

	id, err := manager.Add(name, phone, email, pwd)
	if err != nil {
		res.SetMessage("插入失败")
	}

	// aa, _ := manager.GetManagerSession(c)

	res.SetData(id)

	c.JSON(http.StatusOK, res)
	return
}
