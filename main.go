package main

import (
	"easy_curd/conf"
	"easy_curd/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	_ = r.Run(":3000")
}
