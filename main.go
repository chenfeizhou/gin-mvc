package main

import (
	"github.com/gin-mvc/helpers"
	"github.com/gin-mvc/routes"
)

func main() {

	// 加载配置
	cfg := helpers.LoadIni()

	// 数据库
	helpers.LoadDB(cfg)

	// 初始化路由
	engine := routes.InitRouter()

	engine.Run(":" + cfg.Application.HttpPort)
}
