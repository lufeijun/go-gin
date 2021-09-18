package database

import (
	"fmt"
	"gin/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

// 初始化连接数据库代码。连接断开后，会自动重连，但不会再走 init 函数了
func init() {
	// fmt.Println("sql init")
	var err error
	sqlStr := config.DbUser + ":" + config.DbPass + "@tcp(" + config.DbHost + ":" + config.DbPort + ")/" + config.DbDB + "?charset=utf8mb4&parseTime=true&loc=Local"
	GormDB, err = gorm.Open(mysql.Open(sqlStr), &gorm.Config{}) //配置项中预设了连接池 ConnPool
	if err != nil {
		fmt.Println("数据库连接出现了问题：", err)
		return
	}

	if GormDB.Error != nil {
		fmt.Printf("database error %v", GormDB.Error)
	}

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
