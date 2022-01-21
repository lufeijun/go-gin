package v1

import (
	"gin/structs/response"
	managerUltility "gin/ultility/manager/v1"
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
	}

	err := managerUltility.Login(email, pwd)

	if err != nil {
		res.SetMessage(err.Error())
	}

	c.JSON(http.StatusOK, res)
	return
}

// 添加
func Add(c *gin.Context) {

	res := response.GetResponse()

	name := c.PostForm("name")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	pwd := c.PostForm("pwd")

	id, err := managerUltility.Add(name, phone, email, pwd)
	if err != nil {
		res.SetMessage("插入失败")
	}

	res.SetData(id)

	c.JSON(http.StatusOK, res)
	return
}
