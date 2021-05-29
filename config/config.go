package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
}

func GetConfig() Configuration {
	conf := Configuration{}
	gonfig.GetConf("configs/config.json", &conf)

	return conf
}
