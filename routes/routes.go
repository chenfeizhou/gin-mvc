package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-mvc/app/controller"
	"github.com/gin-mvc/middleware"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	// 设置html目录
	router.LoadHTMLGlob("./view/*")

	// 静态资源映射
	router.Static("/storage", "./storage")

	router.Use(middleware.Cors())

	// 注册路由
	RegisterRouter(router)

	return router
}

// 路由注册
func RegisterRouter(router *gin.Engine) {
	new(controller.IndexController).Router(router)
	new(controller.UserController).Router(router)
}
