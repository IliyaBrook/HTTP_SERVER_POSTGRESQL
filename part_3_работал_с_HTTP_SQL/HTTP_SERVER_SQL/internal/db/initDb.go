package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"main/config"
)

func InitDataBase() *sqlx.DB {
	var err error
	dbString := fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=postgres sslmode=disable search_path=golang_course", config.CfgDb.DBUser, config.CfgDb.DBPassword)
	DB, err = sqlx.Open("postgres", dbString)

	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	return DB
}
