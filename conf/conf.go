package conf

import (
	"easy_curd/cache"
	"easy_curd/model"
	"easy_curd/util"
)

const (
	DATABASE      = "mysql"
	DATABASE_DSN  = "username:pwd@tcp(host:port)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
	REDIS_ADDR    = "127.0.0.1:6379"
	REDIS_PW      = ""
	REDIS_DB      = ""
	SERVER_SECRET = ""
	GIN_MODE      = "debug"
	LOG_LEVEL     = "debug"
	WX_APPID      = ""
	WX_APPSECRET  = ""
	WX_MCHID      = ""
	WX_APPKEY     = ""
)

// Init 初始化配置项
func Init() {

	// 设置日志级别
	util.BuildLogger(LOG_LEVEL)

	// 连接数据库
	model.Database(DATABASE, DATABASE_DSN)
	cache.Redis()
}
