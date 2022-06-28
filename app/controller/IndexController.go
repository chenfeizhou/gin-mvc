package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-mvc/app/repository"
)

type IndexController struct {
}

func (index *IndexController) Router(r *gin.Engine) {
	r.GET("/index", index.Index)
}

func (index *IndexController) Index(c *gin.Context) {

	username := c.DefaultQuery("user_name", "")
	pageSize := c.DefaultQuery("page_size", "20")
	pageNum := c.DefaultQuery("page", "1")

	var data interface{}
	var total int64

	data, total = new(repository.User).GetUsers(username, pageNum, pageSize)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  data,
			"count": total,
		},
	})

}
