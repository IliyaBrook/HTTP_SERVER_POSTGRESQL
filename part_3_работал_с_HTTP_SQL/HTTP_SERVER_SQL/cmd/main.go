package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/config"
	"main/internal/db"
	"main/internal/routes"
	"main/internal/utils"
	"strings"
)

// @title HTTP TEST SERVER
// @description This is a test server for HTTP requests
// @contact.name  Iliya Brook
// @contact.email iliyabrook1987@gmail.com
func main() {
	utils.LoadEnvs()
	if config.CfgApp.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	var trustedProxyList []string
	database := db.InitDataBase()
	defer database.Close()
	r := gin.Default()
	// set trusted proxies from TRUSTED_PROXIES env
	if config.CfgApp.TrustedProxies != "" {
		trustedProxyList = strings.Split(config.CfgApp.TrustedProxies, ",")
		err := r.SetTrustedProxies(trustedProxyList)
		if err != nil {
			log.Fatalf("Failed to set trusted proxies: %v", err)
		}
	}
	// register routes
	routes.RegisterUserRoutes(r)
	routes.RegisterProductsRoutes(r)
	// register cors
	utils.EnableCORS(r)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
