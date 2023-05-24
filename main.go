package main

import (
	"hmdp/config"
	"hmdp/global"
	"hmdp/logger"
	"hmdp/routes"
	"hmdp/services"
	"hmdp/tools"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

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
