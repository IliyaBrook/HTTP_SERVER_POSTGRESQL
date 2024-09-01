package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DbUser string
	DbPass string
)

func LoadEnvs() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DbUser = os.Getenv("DB_USER")
	DbPass = os.Getenv("DB_PASSWORD")
}
