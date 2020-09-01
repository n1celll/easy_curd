package conf

import (
	"os"
	"easy_curd/cache"
	"easy_curd/model"
	"easy_curd/util"
	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	model.Database(os.Getenv("DATABASE"), os.Getenv("DATABASE_DSN"))
	cache.Redis()
}
