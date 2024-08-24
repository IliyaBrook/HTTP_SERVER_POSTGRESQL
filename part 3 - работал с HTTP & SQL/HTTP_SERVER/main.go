package main

import (
	"HTTP_SERVER/handlers"
	"HTTP_SERVER/middlewares"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// load environment file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// env from .env
	//DB_USER
	//DB_PASSWORD
	//sql.Open("postgres", "host=localhost port=5432 user=yourusername password=yourpassword dbname=postgres sslmode=disable search_path=golang_course")
	db, err := sql.Open("postgres", "host=localhost port=5432 user="+os.Getenv("DB_USER")+" password="+os.Getenv("DB_PASSWORD")+" dbname=postgres sslmode=disable search_path=golang_course")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		closeDbErr := db.Close()
		if closeDbErr != nil {
			fmt.Printf("Error closing db connection: %v", closeDbErr)
		}
	}(db)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")

	// users
	http.HandleFunc("/users",
		middlewares.AuthMiddleware(middlewares.LoggerMiddleware(handlers.HandleUsers)),
	)
	// orders
	http.HandleFunc("/orders",
		middlewares.AuthMiddleware(middlewares.LoggerMiddleware(handlers.HandleOrders)),
	)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
