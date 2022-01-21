package config

import (
	"gopkg.in/ini.v1"
)

var (
	APP_NAME string
	APP_MODE string
	APP_PORT string

	// redis
	REDIS_ADDR string
	REDIS_PASS string
	REDIS_DB   int

	// 数据库
	DbHost string
	DbPort string
	DbDB   string
	DbUser string
	DbPass string
	DbConf string
	DbPath = DbHost + DbPort

	KafkaBroker string
)

func init() {

	cfg, err := ini.Load(".env")
	if err != nil {
		panic("读取ini文件失败")
	}

	APP_NAME = cfg.Section("app").Key("name").String()
	APP_MODE = cfg.Section("app").Key("mode").String()
	APP_PORT = cfg.Section("app").Key("port").String()

	DbHost = cfg.Section("mysql").Key("host").String()
	DbPort = cfg.Section("mysql").Key("port").String()
	DbDB = cfg.Section("mysql").Key("dbname").String()
	DbUser = cfg.Section("mysql").Key("username").String()
	DbPass = cfg.Section("mysql").Key("password").String()
	DbConf = cfg.Section("mysql").Key("conf").String()

	// redis
	REDIS_ADDR = cfg.Section("redis").Key("host").String()
	REDIS_PASS = cfg.Section("redis").Key("password").String()
	REDIS_DB, err = cfg.Section("redis").Key("db").Int()
	if err != nil {
		panic("REDIS_DB 有误")
	}

	// kafka
	KafkaBroker = cfg.Section("kafka").Key("broker").String()

}
