package test

import (
	"gin/database/orm"
	"gin/models/clickhouse/v1/manager"
	"gin/structs/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ClickhouseOne(c *gin.Context) {

	res := response.GetResponse()

	manager := manager.Manager{
		ID:      2,
		Name:    "姓名一",
		Address: "北京一",
	}

	result := orm.ClickhouseOrm.Create(&manager)

	if result.Error != nil {
		res.SetMessage(result.Error.Error())
	} else {
		res.SetData(manager)
	}

	c.JSON(http.StatusOK, res)
	return
}
