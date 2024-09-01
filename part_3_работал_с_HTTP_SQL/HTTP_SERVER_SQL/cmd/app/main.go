package main

import (
	"github.com/rs/cors"
	"log"
	"main/internal/db"
	"main/internal/routes"
	"net/http"
)

func main() {
	database := db.InitDataBase()
	//goland:noinspection ALL
	defer database.Close()

	mux := http.NewServeMux()

	// cors
	corsHandler := cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Set-Cookie"},
		AllowCredentials: true,
	}).Handler

	// users

	//mux.HandleFunc("/users",
	//	middlewares.AuthMiddleware(middlewares.LoggerMiddleware(handlers.HandleUsers)),
	//)
	//mux.HandleFunc("/userProducts",
	//	middlewares.AuthMiddleware(middlewares.LoggerMiddleware(handlers.HandleUserProducts)),
	//)
	//// orders
	//mux.HandleFunc("/products",
	//	middlewares.AuthMiddleware(middlewares.LoggerMiddleware(handlers.HandleProducts)),
	//)

	// register routes
	routes.RegisterUserRoutes(mux)
	routes.RegisterProductRoutes(mux)

	handler := corsHandler(mux)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
