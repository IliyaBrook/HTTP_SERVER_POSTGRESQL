package main

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/handlers"
	"HTTP_SERVER/middlewares"
	"log"
	"net/http"
)

func main() {
	data.InitDataBase()
	// users
	http.HandleFunc("/users",
		middlewares.AuthMiddleware(middlewares.LoggerMiddleware(handlers.HandleUsers)),
	)
	// orders
	http.HandleFunc("/products",
		middlewares.AuthMiddleware(middlewares.LoggerMiddleware(handlers.HandleOrders)),
	)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
