package orm

import (
	"gin/config"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

var ClickhouseOrm *gorm.DB

func init() {

	// 检测是否启用
	if !config.Clickhouse_Is_Open {
		return
	}

	var err error
	dsn := "tcp://" + config.Clickhouse_Host + ":" + config.Clickhouse_Port + "?database=" + config.Clickhouse_Db_Name + "&username=" + config.Clickhouse_User_Name + "&password=" + config.Clickhouse_Password + "&read_timeout=" + config.Clickhouse_Read_Timeout + "&write_timeout=" + config.Clickhouse_Write_Timeout

	ClickhouseOrm, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("clickhouse 数据库连接出现了问题：" + err.Error())
	}

	if ClickhouseOrm.Error != nil {
		panic("clickhouse 数据库错误：" + ClickhouseOrm.Error.Error())
	}

	sqlDB, _ := ClickhouseOrm.DB()
	sqlDB.SetMaxIdleConns(config.Clickhouse_Max_Idle_Conns)
	sqlDB.SetMaxOpenConns(config.Clickhouse_Max_Open_Conns)

	// 这里不能有这句话，否则链接都被关闭了
	// defer sqlDB.Close()

}
