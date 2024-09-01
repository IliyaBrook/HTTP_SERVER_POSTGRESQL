package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"main/internal/env"
)

var DB *sqlx.DB

func InitDataBase() *sqlx.DB {
	// load environment file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// env from .env
	DB, err = sqlx.Open("postgres", "host=localhost port=5432 user="+env.DbUser+" password="+env.DbPass+" dbname=postgres sslmode=disable search_path=golang_course")
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	return DB
}
