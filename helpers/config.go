package helpers

const IniPath = "./config/app.ini"

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

type Redis struct {
	RedisHost     string
	RedisPassword string
	RedisPort     string
}

type Config struct {
	Application Application
	Database    Database
	Redis       Redis
}
