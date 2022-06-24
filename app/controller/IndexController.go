package controller

import "github.com/gin-gonic/gin"

type IndexController struct {
}

func (index *IndexController) Router(r *gin.Engine) {
	r.GET("/index", index.Index)
}

func (index *IndexController) Index(c *gin.Context) {

}
