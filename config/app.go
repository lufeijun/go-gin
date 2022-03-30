package config

import (
	"gopkg.in/ini.v1"
)

var (
	APP_NAME    string
	APP_MODE    string
	APP_PORT    string
	APP_DEBUG   bool
	APP_IS_CRON bool

	// email
	MAIL_DRIVER   string
	MAIL_HOST     string
	MAIL_PORT     int
	MAIL_FROM     string
	MAIL_PASSWORD string
	MAIL_NAME     string
	MAIL_IS_SEND  bool

	// session
	SESSION_KEY        string
	SESSION_COOKIE_KEY string
	SESSION_MAX_AGE    int
	// session redis
	SESSION_REDIS_ADDR        string
	SESSION_REDIS_PASS        string
	SESSION_REDIS_DB          string
	SESSION_REDIS_CONNECTIONS int

	// redis
	REDIS_ADDR string
	REDIS_PASS string
	REDIS_DB   int

	// 数据库

	// mysql
	DbIsOpen       bool
	DbHost         string
	DbPort         string
	DbDB           string
	DbUser         string
	DbPass         string
	DbConf         string
	DbPath         = DbHost + DbPort
	DbMaxIdleConns int
	DbMaxOpenConns int

	// clickhouse
	Clickhouse_Is_Open        bool
	Clickhouse_Host           string
	Clickhouse_Port           string
	Clickhouse_Db_Name        string
	Clickhouse_User_Name      string
	Clickhouse_Password       string
	Clickhouse_Read_Timeout   string
	Clickhouse_Write_Timeout  string
	Clickhouse_Max_Idle_Conns int
	Clickhouse_Max_Open_Conns int

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
	APP_DEBUG, err = cfg.Section("app").Key("debug").Bool()
	if err != nil {
		APP_DEBUG = false
	}

	APP_IS_CRON, err = cfg.Section("app").Key("is_cron").Bool()
	if err != nil {
		APP_IS_CRON = false
	}

	// 邮箱
	MAIL_DRIVER = cfg.Section("email").Key("drive").String()
	MAIL_HOST = cfg.Section("email").Key("host").String()
	MAIL_PORT, _ = cfg.Section("email").Key("port").Int()
	MAIL_FROM = cfg.Section("email").Key("from").String()
	MAIL_PASSWORD = cfg.Section("email").Key("password").String()
	MAIL_NAME = cfg.Section("email").Key("name").String()
	MAIL_IS_SEND, err = cfg.Section("email").Key("is_send").Bool()
	if err != nil {
		MAIL_IS_SEND = false
	}

	// mysql
	// Bbis_open
	DbIsOpen, err = cfg.Section("mysql").Key("is_open").Bool()
	if err != nil {
		DbIsOpen = false
	}
	DbHost = cfg.Section("mysql").Key("host").String()
	DbPort = cfg.Section("mysql").Key("port").String()
	DbDB = cfg.Section("mysql").Key("dbname").String()
	DbUser = cfg.Section("mysql").Key("username").String()
	DbPass = cfg.Section("mysql").Key("password").String()
	DbConf = cfg.Section("mysql").Key("conf").String()

	DbMaxIdleConns, err = cfg.Section("mysql").Key("max_idle_conns").Int()

	if err != nil {
		DbMaxIdleConns = 10
	}

	DbMaxOpenConns, err = cfg.Section("mysql").Key("max_open_conns").Int()
	if err != nil {
		DbMaxOpenConns = 100
	}

	DbConf = cfg.Section("mysql").Key("conf").String()

	// clickhouse
	Clickhouse_Is_Open, err = cfg.Section("clickhouse").Key("is_open").Bool()
	if err != nil {
		Clickhouse_Is_Open = false
	}
	Clickhouse_Host = cfg.Section("clickhouse").Key("host").String()
	Clickhouse_Port = cfg.Section("clickhouse").Key("port").String()
	Clickhouse_Db_Name = cfg.Section("clickhouse").Key("port").String()
	Clickhouse_Db_Name = cfg.Section("clickhouse").Key("dbname").String()
	Clickhouse_User_Name = cfg.Section("clickhouse").Key("username").String()
	Clickhouse_Password = cfg.Section("clickhouse").Key("password").String()
	Clickhouse_Read_Timeout = cfg.Section("clickhouse").Key("read_timeout").String()
	Clickhouse_Write_Timeout = cfg.Section("clickhouse").Key("write_timeout").String()
	Clickhouse_Max_Idle_Conns, err = cfg.Section("clickhouse").Key("max_idle_conns").Int()
	if err != nil {
		Clickhouse_Max_Idle_Conns = 3
	}

	Clickhouse_Max_Open_Conns, err = cfg.Section("clickhouse").Key("max_open_conns").Int()
	if err != nil {
		Clickhouse_Max_Open_Conns = 100
	}

	// redis
	REDIS_ADDR = cfg.Section("redis").Key("addr").String()
	REDIS_PASS = cfg.Section("redis").Key("password").String()
	REDIS_DB, err = cfg.Section("redis").Key("db").Int()
	if err != nil {
		REDIS_DB = 0
	}

	// session
	SESSION_KEY = cfg.Section("session").Key("cookie_key").String()
	SESSION_COOKIE_KEY = cfg.Section("session").Key("cookie_key").String()
	SESSION_MAX_AGE, err = cfg.Section("session").Key("session_max_age").Int()
	if err != nil {
		SESSION_MAX_AGE = 7200
	}

	SESSION_REDIS_ADDR = cfg.Section("session").Key("redis_addr").String()
	SESSION_REDIS_PASS = cfg.Section("session").Key("redis_password").String()
	SESSION_REDIS_DB = cfg.Section("session").Key("redis_db").String()
	SESSION_REDIS_CONNECTIONS, err = cfg.Section("session").Key("redis_connections").Int()
	if err != nil {
		SESSION_REDIS_CONNECTIONS = 10
	}

	// kafka
	KafkaBroker = cfg.Section("kafka").Key("broker").String()

}
