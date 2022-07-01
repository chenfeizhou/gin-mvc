package test

import (
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Article struct {
	Title      string
	CategoryId uint
	Category   Category `gorm:"foreignkey:CategoryId"`
}

type Category struct {
	model.baseModel
	Name string
}

func TestGormTest(t *testing.T) {

	dsn := "root:123456@tcp(127.0.0.1:3306)/gin_test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Article{}, &Category{})

	// category := Category{Name: "分类2"}
	// result1 := db.Create(&category) // 通过数据的指针来创建
	// fmt.Println(result1)

	// article := Article{Title: "test2", CategoryId: category.ID}
	// result2 := db.Create(&article) // 通过数据的指针来创建
	// fmt.Println(result2)
	// fmt.Println(article.ID)

	// Retrieve user list with edger loading credit card

	var articles []Article

	err = db.Model(&Article{}).Preload("Category").Find(&articles).Error

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(articles)
}
