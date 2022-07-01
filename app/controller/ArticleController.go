package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-mvc/app/model"
	"github.com/gin-mvc/app/repository"
)

type ArticleController struct {
}

func (articleController *ArticleController) Router(r *gin.Engine) {
	r.GET("/articles", articleController.Index)
	r.GET("/articles/:id", articleController.Show)
	r.POST("/articles", articleController.Create)
	r.DELETE("/articles/:id", articleController.Destory)
}

func (articleController *ArticleController) Index(c *gin.Context) {

	title := c.DefaultQuery("title", "")
	pageSize := c.DefaultQuery("page_size", "20")
	pageNum := c.DefaultQuery("page", "1")

	var data interface{}
	var total int64

	data, total = new(repository.ArticleRepository).GetArticles(title, pageNum, pageSize)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  data,
			"count": total,
		},
	})
}

func (articleController *ArticleController) Create(c *gin.Context) {
	var creatArticleRequest model.CreateArticleRequest

	if c.ShouldBind(&creatArticleRequest) == nil {
		return
	}

	article := new(repository.ArticleRepository).CreateArticle(creatArticleRequest)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": article,
	})
}

func (articleController *ArticleController) Show(c *gin.Context) {

	type ArticleRequest struct {
		ID string `uri:"id"`
	}

	var request ArticleRequest

	if err := c.ShouldBindUri(&request); err != nil {
		return
	}

	articleId, _ := strconv.Atoi(request.ID)

	article := new(repository.ArticleRepository).GetArticleDetail(articleId)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": article,
	})
}

func (articleController *ArticleController) Destory(c *gin.Context) {

	type ArticleRequest struct {
		ID string `uri:"id"`
	}

	var request ArticleRequest

	if err := c.ShouldBindUri(&request); err != nil {
		return
	}

	articleId, _ := strconv.Atoi(request.ID)

	new(repository.ArticleRepository).DeleteArticle(articleId)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": nil,
	})
}
