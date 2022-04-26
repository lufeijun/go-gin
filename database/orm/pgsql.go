package orm

import (
	"gin/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PgsqlOrm *gorm.DB

// 初始化连接数据库代码。连接断开后，会自动重连，但不会再走 init 函数了
func init() {

	// 判断是否需要 mysql 链接
	if !config.Pgsql_Is_Open {
		return
	}

	var err error

	dsn := "host=" + config.Pgsql_Host + " user=" + config.Pgsql_User_Name + " password=" + config.Pgsql_Password + " dbname=" + config.Pgsql_Db_Name + " port=" + config.Pgsql_Port + " sslmode=" + config.Pgsql_Sslmode + " TimeZone=" + config.Pgsql_Time_Zone
	PgsqlOrm, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("pgsql 数据库连接出现了问题：" + err.Error())
	}

	if PgsqlOrm.Error != nil {
		panic("pgsql 数据库错误：" + PgsqlOrm.Error.Error())
	}

	sqlDB, _ := PgsqlOrm.DB()
	sqlDB.SetMaxIdleConns(config.Pgsql_Max_Idle_Conns)
	sqlDB.SetMaxOpenConns(config.Pgsql_Max_Open_Conns)

	// 这里不能有这句话，否则链接都被关闭了
	// defer sqlDB.Close()

}
