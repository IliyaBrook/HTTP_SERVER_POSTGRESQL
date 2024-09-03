package main

import (
	"dog_breeds_search/docs"
	"dog_breeds_search/routes"
	"dog_breeds_search/sharable"
	"github.com/rs/cors"
	_ "github.com/swaggo/files"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title Dog Breeds API
// @version 1.0
// @description Dog Breeds API project from golang ninja course
// @BasePath
// @contact.name   Iliya Brook
// @contact.email  iliyabrook1987@gmail.com
func main() {
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	sharable.LoadEnvironmentVariables()
	mux := http.NewServeMux()
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

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
