package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/config"
	"main/internal/cors"
	"main/internal/db"
	"main/internal/routes"
	"main/pkg"
	"strings"
)

func main() {
	pkg.LoadEnvs()
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
	cors.EnableCORS(r)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
