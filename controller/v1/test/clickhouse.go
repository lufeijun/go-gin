package test

import (
	"fmt"
	"gin/database/orm"
	"gin/models/clickhouse/v1/manager"
	"gin/models/clickhouse/v1/test"
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

func Wechat_user_sessions_list(c *gin.Context) {
	res := response.GetResponse()

	// var mysqls []test.Wechatusersessions
	// orm.ClickhouseOrm.Where("msgtime >= ?", "2022-04-01 00:00:00").Where("msgtime <= ?", "2022-04-06 00:00:00").Distinct("roomid").Find(&mysqls)

	var mysqls []string
	var wechat test.Wechatusersessions
	// result := orm.ClickhouseOrm.Model(&wechat).Where("msgtime >= ?", "2022-04-01 00:00:00").Where("msgtime <= ?", "2022-04-06 00:00:00").Distinct().Pluck("roomid", &mysqls)
	result := orm.ClickhouseOrm.Model(&wechat).Where("created_at >= ?", "2022-04-01 00:00:00").Where("created_at <= ?", "2022-04-06 00:00:00").Distinct().Pluck("roomid", &mysqls)
	// orm.MysqlOrm.Model(&wechat).Where("msgtime >= ?", "2022-04-01 00:00:00").Where("msgtime <= ?", "2022-04-06 00:00:00").Distinct().Pluck("roomid", &mysqls)

	fmt.Println(result.RowsAffected)

	res.SetData(mysqls)

	c.JSON(http.StatusOK, res)
	return
}

func Wechat_user_sessions(c *gin.Context) {
	res := response.GetResponse()

	// var msg test.Wechatusersessions

	id := 20110001
	limit := 100000

	for j := 0; j < 20; j++ {
		go Wechat_user_sessions_go(id, limit)
		id += limit
	}

	c.JSON(http.StatusOK, res)
	return
}

func Wechat_user_sessions_go(id, limit int) {
	fmt.Println("start...", id)

	var mysqls []test.Wechatusersessions

	result := orm.MysqlOrm.Where("id > ?", id).Limit(limit).Find(&mysqls)

	if result.Error != nil {
		fmt.Println("出错" + result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		fmt.Println(id, "==0")
		return
	}

	orm.ClickhouseOrm.Create(&mysqls)

	fmt.Println("end...", id)
}
