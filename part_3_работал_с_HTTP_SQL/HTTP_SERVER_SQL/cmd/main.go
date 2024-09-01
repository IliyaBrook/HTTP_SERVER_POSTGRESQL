package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/internal/cors"
	"main/internal/db"
	"main/internal/env"
	"main/internal/routes"
	"strings"
)

func main() {
	// load environment variables from .env
	env.LoadEnvs()
	if env.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	var trustedProxyList []string
	// init postgres sql db
	database := db.InitDataBase()
	defer database.Close()
	r := gin.Default()
	// set trusted proxies from TRUSTED_PROXIES env
	if env.TrustedProxies != "" {
		trustedProxyList = strings.Split(env.TrustedProxies, ",")
		err := r.SetTrustedProxies(trustedProxyList)
		if err != nil {
			log.Fatalf("Failed to set trusted proxies: %v", err)
		}
	}
	// register routes
	routes.RegisterUserRoutes(r)
	routes.RegisterProductsRoutes(r)
	// register cors
	cors.EnableCORS(r)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
