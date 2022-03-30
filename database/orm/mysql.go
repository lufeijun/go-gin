package orm

import (
	"gin/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlOrm *gorm.DB

// 初始化连接数据库代码。连接断开后，会自动重连，但不会再走 init 函数了
func init() {

	// 判断是否需要 mysql 链接
	if !config.DbIsOpen {
		return
	}

	var err error
	sqlStr := config.DbUser + ":" + config.DbPass + "@tcp(" + config.DbHost + ":" + config.DbPort + ")/" + config.DbDB + "?charset=utf8mb4&parseTime=True&loc=Local"
	MysqlOrm, err = gorm.Open(mysql.Open(sqlStr), &gorm.Config{}) //配置项中预设了连接池 ConnPool
	if err != nil {
		panic("mysql 数据库连接出现了问题：" + err.Error())
	}

	if MysqlOrm.Error != nil {
		panic("mysql 数据库错误：" + MysqlOrm.Error.Error())
	}

	sqlDB, _ := MysqlOrm.DB()
	sqlDB.SetMaxIdleConns(config.DbMaxIdleConns)
	sqlDB.SetMaxOpenConns(config.DbMaxOpenConns)

	// 这里不能有这句话，否则链接都被关闭了
	// defer sqlDB.Close()

}

func Paginate(page int64, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize

		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
