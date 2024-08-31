package main

import (
	"assignment/routes"
	"assignment/sharable"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	port := os.Getenv("PORT")

	allowedOrigins := []string{
		sharable.AllowedOrigins1,
		sharable.AllowedOrigins2,
	}

	corsHandler := cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool {
			for _, o := range allowedOrigins {
				if origin == o {
					return true
				}
			}
			return false
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "X-Custom-Header", "Another-Header"},
		ExposedHeaders:   []string{"Set-Cookie"},
		AllowCredentials: true,
	}).Handler
	log.Println("Starting server on port " + sharable.PORT)

	routes.RouteFunctions(mux)

	handler := corsHandler(mux)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
