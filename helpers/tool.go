package helpers

import (
	"strings"

	"gopkg.in/ini.v1"
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
