package main

import (
	"assignment/routes"
	"assignment/sharable"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	sharable.LoadEnvironmentVariables()
	mux := http.NewServeMux()

	port := sharable.PORT

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

	routes.RouteFunctions(mux)

	handler := corsHandler(mux)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
