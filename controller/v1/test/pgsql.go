package test

import (
	"gin/database/orm"
	"gin/models"
	"gin/models/pgsql"
	"gin/models/v1/test"
	"gin/structs"
	"gin/structs/response"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PgsqlOne(c *gin.Context) {

	res := response.GetResponse()

	var tests []test.TestOne
	sql := orm.PgsqlOrm

	var pagestruct structs.PageStruct
	var page int64
	var pagesize int64
	var total int64

	page = 1
	pagesize = 3

	sql.Model(&test.TestOne{}).Count(&total)
	pagestruct.LastPage = int64(math.Ceil(float64(total) / float64(pagesize)))

	result := sql.Scopes(orm.Paginate(page, pagesize)).Order("id desc").Find(&tests)
	if result.Error != nil {
		panic(result.Error)
	}

	// 赋值
	pagestruct.Total = total
	pagestruct.Page = page
	pagestruct.Size = pagesize
	pagestruct.Data = tests

	res.SetData(pagestruct)

	c.JSON(http.StatusOK, res)
}

func PgsqlTwo(c *gin.Context) {

	res := response.GetResponse()

	// str, _ := time.Now()

	test := test.TestOne{
		Date: models.GormTime{Time: time.Now()},
	}

	orm.PgsqlOrm.Create(&test)

	res.SetData(test)

	c.JSON(http.StatusOK, res)
}

func PgsqlUserInsert(c *gin.Context) {
	res := response.GetResponse()
	user := pgsql.User{
		// Id:      1,
		Name:    "吉鹏",
		Age:     10,
		Address: "测试😊",
	}

	result := orm.PgsqlOrm.Create(&user)

	if result.Error != nil {
		res.SetMessage(result.Error.Error())
		c.JSON(http.StatusOK, res)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func PgsqlUserGet(c *gin.Context) {
	res := response.GetResponse()

	var user pgsql.User

	orm.PgsqlOrm.Last(&user)

	res.SetData(user)

	c.JSON(http.StatusOK, res)
	return
}
