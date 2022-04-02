package test

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ReloadOne(c *gin.Context) {

	count := "2"

	start := time.Now().Format("2006-01-02 15:04:05")

	fmt.Println("start==", count, "===", start)

	time.Sleep(10 * time.Second)

	end := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("end=", count, "===", end)

	c.String(http.StatusOK, start+"===reload"+count+"。。。。==="+end)
}
