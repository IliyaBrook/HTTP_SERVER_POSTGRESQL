package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/internal/cors"
	"main/internal/db"
	"main/internal/env"
	"main/internal/routes"
)

func main() {
	// load environment variables from .env
	env.LoadEnvs()
	// init postgres sql db
	database := db.InitDataBase()
	defer database.Close()
	r := gin.Default()
	// register routes
	routes.RegisterUserRoutes(r)
	routes.RegisterProductsRoutes(r)
	// register cors
	cors.EnableCORS(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
