package testv2

import (
	"gin/structs/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	res := response.GetResponse()

	res.SetData("v2")
	c.JSON(http.StatusOK, res)

	return
}
