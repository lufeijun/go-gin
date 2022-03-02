package test

import (
	"gin/structs/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	res := response.GetResponse()

	res.SetData("v1")

	c.JSON(http.StatusOK, res)

	return
}

func Jwt(c *gin.Context) {
	res := response.GetResponse()

	res.SetData("jwt")

	c.JSON(http.StatusOK, res)

	return
}
