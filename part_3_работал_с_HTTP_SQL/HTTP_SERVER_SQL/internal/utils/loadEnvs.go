package utils

import (
	"fmt"
	"log"
	"main/config"
	"os"
	"path/filepath"
)

func LoadEnvs() {
	rootPath, _ := os.Getwd()
	configPath := filepath.Join(rootPath, "config")
	if err := config.New[config.DbConfig](configPath, "db", &config.CfgDb); err != nil {
		fmt.Println("Error reading .env file for DB", err)
		log.Fatal(err)
	}
	if err := config.New[config.AppConfig]("../config", "app", &config.CfgApp); err != nil {
		fmt.Println("Error reading .env file for APP", err)
		log.Fatal(err)
	}
}
