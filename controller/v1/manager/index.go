package manager

import (
	"gin/structs/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	res := response.GetResponse()

	c.JSON(http.StatusOK, res)
	return
}
