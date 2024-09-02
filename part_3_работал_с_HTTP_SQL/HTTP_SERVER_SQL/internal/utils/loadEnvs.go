package utils

import (
	"fmt"
	"log"
	"main/config"
)

func LoadEnvs() {
	if err := config.New[config.DbConfig]("../config", "db", &config.CfgDb); err != nil {
		fmt.Println("Error reading .env file for DB", err)
		log.Fatal(err)
	}
	if err := config.New[config.AppConfig]("../config", "app", &config.CfgApp); err != nil {
		fmt.Println("Error reading .env file for APP", err)
		log.Fatal(err)
	}
}
