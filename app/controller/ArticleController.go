package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-mvc/app/model"
	"github.com/gin-mvc/app/repository"
	"github.com/gin-mvc/helpers"
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

	helpers.Success(c, gin.H{"list": data, "count": total}, "success")
}

func (articleController *ArticleController) Create(c *gin.Context) {
	var creatArticleRequest model.CreateArticleRequest

	if c.ShouldBind(&creatArticleRequest) == nil {
		return
	}

	article := new(repository.ArticleRepository).CreateArticle(creatArticleRequest)

	helpers.Success(c, gin.H{"data": article}, "success")
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

	helpers.Success(c, gin.H{"data": article}, "success")
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

	helpers.Success(c, gin.H{"data": nil}, "success")
}
