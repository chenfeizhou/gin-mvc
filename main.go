package main

import (
	"github.com/gin-mvc/app/model"
	"github.com/gin-mvc/helpers"
	"github.com/gin-mvc/routes"
)

func main() {

	// 加载配置
	cfg := new(helpers.Config).InitConfig()

	// 初始化数据库
	model.InitDb(cfg)

	// 初始化路由
	engine := routes.InitRouter()

	engine.Run(":" + cfg.Application.HttpPort)
}
