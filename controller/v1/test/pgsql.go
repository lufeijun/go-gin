package test

import (
	"gin/database/orm"
	"gin/models"
	"gin/models/v1/test"
	"gin/structs/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PgsqlOne(c *gin.Context) {

	res := response.GetResponse()
	test := test.TestOne{}

	orm.PgsqlOrm.First(&test)

	res.SetData(test)

	c.JSON(http.StatusOK, res)
}

func PgsqlTwo(c *gin.Context) {

	res := response.GetResponse()

	str, _ := time.Parse("2006-01-02 15:04:05", "2022-03-25 12:12:12")

	test := test.TestOne{
		ID:   3,
		Date: models.GormTime{Time: str},
	}

	orm.PgsqlOrm.Create(&test)

	res.SetData(test)

	c.JSON(http.StatusOK, res)
}
