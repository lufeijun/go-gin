package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	APP_NAME string
	DbHost   string
	DbPort   string
	DbDB     string
	DbUser   string
	DbPass   string
	DbConf   string
	DbPath   = DbHost + DbPort

	KafkaBroker string
)

func init() {

	cfg, err := ini.Load(".env")
	if err != nil {
		fmt.Println("读取ini文件失败")
		//   log.Fatalln(err)
	}

	APP_NAME = cfg.Section("").Key("app_mode").String()
	DbHost = cfg.Section("mysql").Key("host").String()
	DbPort = cfg.Section("mysql").Key("port").String()
	DbDB = cfg.Section("mysql").Key("dbname").String()
	DbUser = cfg.Section("mysql").Key("username").String()
	DbPass = cfg.Section("mysql").Key("password").String()
	DbConf = cfg.Section("mysql").Key("conf").String()

	KafkaBroker = cfg.Section("kafka").Key("broker").String()

}
