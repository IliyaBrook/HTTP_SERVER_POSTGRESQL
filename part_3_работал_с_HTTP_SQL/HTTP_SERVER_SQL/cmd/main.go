package main

import (
	"log"
	"main/internal/cors"
	"main/internal/db"
	"main/internal/env"
	"main/internal/routes"
	"net/http"
)

func main() {
	// load environment variables from .env
	env.LoadEnvs()
	// init postgres sql db
	database := db.InitDataBase()
	defer database.Close()
	mux := http.NewServeMux()
	// register routes
	routes.RegisterUserRoutes(mux)
	routes.RegisterProductRoutes(mux)
	// register cors
	handler := cors.CORSHandler(mux)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
