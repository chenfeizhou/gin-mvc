package model

import (
	"strings"

	"github.com/gin-mvc/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type baseModel struct {
	gorm.Model
}

var db = InitDb()

/**
建立Mysql连接器
*/
func InitDb() *gorm.DB {

	cfg := helpers.LoadIni()

	dsn := strings.Join([]string{cfg.Database.DbUsername, ":", cfg.Database.DbPassword, "@tcp(", cfg.Database.DbHost, ":", cfg.Database.DbPort, ")/", cfg.Database.DbDatabase, "?charset=utf8&parseTime=true"}, "")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
