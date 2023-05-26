package main

import (
	"hmdp/config"
	"hmdp/global"
	"hmdp/logger"
	"hmdp/routes"
	"hmdp/services"
	"hmdp/tools"
	_ "net/http/pprof"
)

func main() {

	// 初始化配置
	config.InitializeConfig()

	// 初始化日志
	logger.InitLog()

	// 初始化redis
	global.App.Redis = tools.InitializeRedis()

	// 业务服务层初始化
	services.Init()

	routes.RunServer()

}
