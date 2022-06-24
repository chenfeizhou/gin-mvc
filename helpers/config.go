package helpers

import (
	"bufio"
	"encoding/json"
	"os"
)

type Application struct {
	AppName  string `json:"app_name"`
	AppMode  string `json:"app_mode"`
	HttpHost string `json:"http_host"`
	HttpPort string `json:"http_port"`
}

type Database struct {
	Connection string `json:"db_connection"`
	Host       string `json:"db_host"`
	Port       string `json:"db_port"`
	Database   string `json:"db_database"`
	Username   string `json:"db_username"`
	Password   string `json:"db_password"`
}

type Config struct {
	Application
	Database
}

const ConfigPath = "./config/app.ini"

var cfg *Config = nil

func LoadConfig() (*Config, error) {

	file, err := os.Open(ConfigPath)

	handlerErr(err)

	defer file.Close()

	reader := bufio.NewReader(file)

	decoder := json.NewDecoder(reader)

	if err = decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
