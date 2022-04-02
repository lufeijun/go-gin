package test

import (
	"fmt"
	"gin/database/orm"
	"gin/models/clickhouse/v1/manager"
	"gin/structs/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ClickhouseOne(c *gin.Context) {

	res := response.GetResponse()

	clickhouseInsertBatch()
	// clickhouseInsertOne()

	c.JSON(http.StatusOK, res)
	return
}

func clickhouseInsertOne() {
	manager := manager.Manager{
		ID:      2,
		Name:    "姓名一",
		Address: "北京一",
	}
	orm.ClickhouseOrm.Create(&manager)
}

func clickhouseInsertBatch() {
	for j := 0; j < 10; j++ {
		go clickhouseInsertBatchgo(j)
	}

}

// 协程开启
func clickhouseInsertBatchgo(index int) {
	fmt.Println("start...", index)
	length := 1000000
	var managers [1000000]manager.Manager

	for i := 0; i < length; i++ {
		managers[i] = manager.Manager{
			ID:      uint64(i + 1),
			Name:    "姓名：" + strconv.Itoa(i%50),
			Address: "北京：" + strconv.Itoa(i),
		}
	}
	orm.ClickhouseOrm.Create(&managers)
	fmt.Println("end...", index)
}

// 获取列表
func ClickhouseTwo(c *gin.Context) {

	res := response.GetResponse()

	var list []manager.Manager

	orm.ClickhouseOrm.Where("name= ?", "姓名：24").Find(&list)

	res.SetData(list)

	c.JSON(http.StatusOK, res)
	return
}
