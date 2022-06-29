package helpers

import (
	"strings"

	"github.com/gin-mvc/app/model"
	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func HandlerErr(err error) {
	if err != nil {
		panic(err)
	}
}

func FirstToUpper(str string) string {

	if str == "" {
		return ""
	}

	return strings.ToUpper(str[:1]) + str[1:]
}

func LoadIni() Config {

	var config Config

	cfg, err := ini.Load(IniPath)

	HandlerErr(err)

	cfg.NameMapper = ini.TitleUnderscore

	err = cfg.MapTo(&config)

	HandlerErr(err)

	return config
}

// 加载数据库连接池
func LoadDB(cfg Config) {

	dsn := strings.Join([]string{cfg.Database.DbUsername, ":", cfg.Database.DbPassword, "@tcp(", cfg.Database.DbHost, ":", cfg.Database.DbPort, ")/", cfg.Database.DbDatabase, "?charset=utf8&parseTime=true"}, "")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	model.DB = db
}

// 加载缓存
func LoadCache(cfg Config) {

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.RedisHost,
		Password: cfg.Redis.RedisPassword,
	})

	_, err := client.Ping().Result()

	// 可以不报错，邮件提醒
	HandlerErr(err)

	RedisClient = client
}
