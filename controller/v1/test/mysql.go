package test

import (
	"fmt"
	"gin/database/orm"
	"gin/models"
	"gin/models/v1/test"
	"gin/structs/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 数据库操作测试

// 时间字段的
func MysqlOne(c *gin.Context) {
	res := response.GetResponse()

	str, err := time.Parse("2006-01-02 15:04:05", "2022-03-25 12:12:12")

	if err != nil {
		res.SetMessage(err.Error())
		c.JSON(http.StatusOK, res)
		return
	}

	test := test.TestOne{
		// Date: models.GormTime{time.Now()},
		Date: models.GormTime{str},
	}

	result := orm.MysqlOrm.Create(&test)

	if result.Error != nil {
		fmt.Println("报错信息")
		fmt.Println(result.Error)
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	res.SetData(test)

	c.JSON(http.StatusOK, res)

}

// 事务
func Mysql2(c *gin.Context) {
	res := response.GetResponse()

	tx := orm.MysqlOrm.Statement.DB.Begin()

	test := test.TestOne{
		Date: models.GormTime{time.Now()},
	}

	tx.Create(&test)

	tx.Commit()

	fmt.Println("11111")

	c.JSON(http.StatusOK, res)
	return
}
