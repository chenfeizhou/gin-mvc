package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-mvc/app/model"
)

func main() {

	// 加载配置

	// 初始化数据库
	model.Connect()

	// 加载路由

	// 运行
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin mvc Framework",
		})
	})

	r.Run()
}
