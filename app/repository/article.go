package repository

import (
	"strconv"

	"github.com/gin-mvc/app/model"
)

type ArticleRepository struct {
}

func (articleRepository *ArticleRepository) GetArticles(title string, pageNum string, pageSize string) ([]model.Article, int64) {

	var articles []model.Article
	var total int64

	page, _ := strconv.Atoi(pageNum)
	page_size, _ := strconv.Atoi(pageSize)

	articles, total = model.GetArticles(title, page, page_size)

	return articles, total
}

func (articleRepository *ArticleRepository) GetArticleDetail(id int) model.Article {
	return model.GetArticleDetail(id)
}

func (articleRepository *ArticleRepository) CreateArticle(createArticleRequest model.CreateArticleRequest) model.Article {
	return model.CreateArticle(createArticleRequest)
}

func (articleRepository *ArticleRepository) DeleteArticle(id int) {
	model.DB.Delete(&model.Article{}, id)
}
