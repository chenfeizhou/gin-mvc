package model

import "strconv"

type Article struct {
	BaseModel
	UserId     uint
	User       User `gorm:"foreignkey:UserId"`
	CategoryId uint
	Category   Category `gorm:"foreignkey:CategoryId"`
	Title      string   `gorm:"type:varchar(100);not null" `
	Content    string   `gorm:"type:longtext" `
}

func (*Article) TableName() string {
	return "article"
}

//创建用户的请求
type CreateArticleRequest struct {
	ID         string `form:"id" json:"id"`
	Title      string `form:"title" json:"title" binding:"required,min=2,max=500"`
	Content    string `form:"content" json:"content"`
	CategoryId string `form:"category_id" json:"category_id" binding:"required,max=1000"`
}

func GetArticles(title string, pageNum int, pageSize int) ([]Article, int64) {

	var articles []Article
	var total int64

	if title != "" {
		DB.Select("id,title,cid,created_at").Where(
			"title like ?", title+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles)

		DB.Model(&articles).Where(
			"title Like ?", title+"%",
		).Count(&total)

		return articles, total
	}

	DB.Preload("Category").Preload("User").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles)
	DB.Model(&articles).Count(&total)

	return articles, total
}

func GetArticleDetail(id int) Article {

	var article Article
	DB.First(&article, "id = ?", id)
	return article
}

func CreateArticle(creatArticleRequest CreateArticleRequest) Article {

	DB.Create(&creatArticleRequest)

	var article Article

	id, _ := strconv.Atoi(creatArticleRequest.ID)

	article = GetArticleDetail(id)

	return article
}
