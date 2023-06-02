package config

import "sync"

type AppConfig struct {
	SERVER_PORT		int
	DB_DRIVER		string
	DB_USERNAME		string
	DB_PASSWORD		string
	DB_HOST			string
	DB_PORT			string
	DB_NAME			string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Lock()

	if appConfig == nil {
		appConfig = initConfig()
	}
	return appConfig
}

func initConfig() *AppConfig {
	return *&appConfig
}