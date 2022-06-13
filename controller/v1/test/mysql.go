package test

import (
	"context"
	"fmt"
	"gin/database/orm"
	"gin/models"
	"gin/models/v1/test"
	"gin/structs/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		Date: models.GormTime{Time: str},
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

	// 手动事务
	// transaction1()

	// 自动事务
	transaction2()

	c.JSON(http.StatusOK, res)
	return
}

// 手动
func transaction1() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*3))

	defer cancel()

	tx := orm.MysqlOrm.Statement.DB.WithContext(ctx).Begin()

	test := test.TestOne{
		Date: models.GormTime{Time: time.Now()},
	}

	time.Sleep(time.Second * 5)

	aa := tx.Create(&test)

	if aa.Error != nil {
		fmt.Println("create 报错 ：" + aa.Error.Error())
	}

	result := tx.Commit()

	if result.Error != nil {
		fmt.Println("Commit 报错 ：" + tx.Error.Error())
	}
}

// 自动
func transaction2() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*3))

	defer cancel()

	trans := orm.MysqlOrm.Statement.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		test := test.TestOne{
			Date: models.GormTime{Time: time.Now()},
		}

		// time.Sleep(time.Second * 5)

		result := tx.Create(&test)
		if result.Error != nil {
			fmt.Println("create err: " + result.Error.Error())
			return result.Error
		}

		// 返回 nil 。提交事务
		return nil
	})

	if trans != nil {
		fmt.Println("trans err: " + trans.Error())
	}

}

// 多协程查询
func Mysql3(c *gin.Context) {
	res := response.GetResponse()

	for i := 0; i < 10; i++ {
		go asyncInsert()
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	c.JSON(http.StatusOK, res)
	return
}

// 协程插入
func asyncInsert() {
	orm.MysqlOrm.Statement.DB.Transaction(func(tx *gorm.DB) error {
		test := test.TestOne{
			Date: models.GormTime{Time: time.Now()},
		}

		time.Sleep(time.Second * 10)

		result := tx.Create(&test)
		if result.Error != nil {
			fmt.Println("create err: " + result.Error.Error())
			return result.Error
		}

		// 返回 nil 。提交事务
		return nil
	})

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}

func Mysql4(c *gin.Context) {
	res := response.GetResponse()

	for i := 0; i < 20; i++ {
		go asyncInsert4()
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	c.JSON(http.StatusOK, res)
	return
}

func asyncInsert4() {

	for i := 0; i < 100; i++ {
		test := test.TestOne{
			Date: models.GormTime{time.Now()},
		}

		if i%10 == 0 {
			time.Sleep(time.Second)
		}

		orm.MysqlOrm.Create(&test)

	}
}
