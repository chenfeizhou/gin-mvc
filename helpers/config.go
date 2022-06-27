package helpers

import (
	"gopkg.in/ini.v1"
)

type Application struct {
	AppName  string
	AppMode  string
	HttpHost string
	HttpPort string
}

type Database struct {
	DbConnection string
	DbHost       string
	DbPort       string
	DbDatabase   string
	DbUsername   string
	DbPassword   string
}

type Config struct {
	Application Application
	Database    Database
}

const IniPath = "./config/app.ini"

func (config *Config) InitConfig() *Config {

	cfg, err := ini.Load(IniPath)
	HandlerErr(err)

	cfg.NameMapper = ini.TitleUnderscore

	err = cfg.Section("application").MapTo(&config.Application)
	HandlerErr(err)

	err = cfg.Section("database").MapTo(&config.Database)
	HandlerErr(err)

	return config
}
