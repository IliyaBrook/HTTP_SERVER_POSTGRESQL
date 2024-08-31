package sharable

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	ApiKey          = os.Getenv("API_KEY")
	PORT            = os.Getenv("PORT")
	AllowedOrigins1 = os.Getenv("ALLOWED_ORIGINS_1")
	AllowedOrigins2 = os.Getenv("ALLOWED_ORIGINS_2")
	ApiUrl          = os.Getenv("API_URL")
)

func LoadEnvironmentVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		log.Fatal(err)
	}
	ApiKey = os.Getenv("API_KEY")
	PORT = os.Getenv("PORT")
	AllowedOrigins1 = os.Getenv("ALLOWED_ORIGINS_1")
	AllowedOrigins2 = os.Getenv("ALLOWED_ORIGINS_2")
	ApiUrl = os.Getenv("API_URL")
}
