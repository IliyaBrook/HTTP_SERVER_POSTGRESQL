package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"main/config"
)

func InitDataBase() *sqlx.DB {
	var err error
	DB, err = sqlx.Open("postgres", "host=localhost port=5432 user="+config.CfgDb.DBUser+" password="+config.CfgDb.DBPassword+" dbname=postgres sslmode=disable search_path=golang_course")
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	return DB
}
