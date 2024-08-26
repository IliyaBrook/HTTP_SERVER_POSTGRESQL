package data

import (
	"HTTP_SERVER/utils"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var filePath = utils.ResolvePath("data/database.json")

var DB *sqlx.DB

func InitDataBase() *sqlx.DB {
	// load environment file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// env from .env
	DB, err = sqlx.Open("postgres", "host=localhost port=5432 user="+os.Getenv("DB_USER")+" password="+os.Getenv("DB_PASSWORD")+" dbname=postgres sslmode=disable search_path=golang_course")
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	return DB
}

func (d *DbStruct) ReadDatabase() error {

	file, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	if err := json.Unmarshal(file, d); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return err
	}

	return nil
}

func (d *DbStruct) SaveDatabase() error {
	file, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filePath, file, 0644); err != nil {
		return err
	}

	return nil
}

var DbInst = &DbStruct{}
