package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type baseModel struct {
	gorm.Model
}

var db *gorm.DB

/**
建立Mysql连接器
*/
func Connect() *gorm.DB {

	dsn := "root:123456@tcp(127.0.0.1:3306)/gin_cms?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
