package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-mvc/app/controller"
	"github.com/gin-mvc/app/model"
	"github.com/gin-mvc/helpers"
)

func main() {

	// 加载配置
	helpers.InitConfig()

	// 初始化数据库
	model.Connect()

	// 运行
	app := gin.Default()

	// 设置html目录
	app.LoadHTMLGlob("./view/*")

	// 静态资源映射
	app.Static("/storage", "./storage")

	// 注册路由
	RegisterRouter(app)

	app.Run()
}

// 路由注册
func RegisterRouter(router *gin.Engine) {
	new(controller.IndexController).Router(router)
}
