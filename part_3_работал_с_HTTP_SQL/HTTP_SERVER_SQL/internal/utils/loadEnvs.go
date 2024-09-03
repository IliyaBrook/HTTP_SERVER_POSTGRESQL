package utils

import (
	log "github.com/sirupsen/logrus"
	"main/config"
)

func LoadEnvs() {
	configPath := AbsPathJoin("config")
	if err := config.New[config.DbConfig](configPath, "db", &config.CfgDb); err != nil {
		log.Error("Error reading .env file for DB", err)
		log.Fatal(err)
	}
	if err := config.New[config.AppConfig](configPath, "app", &config.CfgApp); err != nil {
		log.Error("Error reading .env file for APP", err)
		log.Fatal(err)
	}
}
