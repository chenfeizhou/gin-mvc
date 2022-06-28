package test

import (
	"fmt"
	"testing"

	"github.com/gin-mvc/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGormTest(t *testing.T) {

	dsn := "root:123456@tcp(127.0.0.1:3306)/gin_cms?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	data := make([]*model.User, 0)

	err = db.Find(&data).Error
	if err != nil {
		panic(err)
	}

	for _, v := range data {
		fmt.Println(v)
	}

}
