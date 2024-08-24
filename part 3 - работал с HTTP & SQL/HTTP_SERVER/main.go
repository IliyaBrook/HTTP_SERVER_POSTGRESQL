package main

import (
	"HTTP_SERVER/handlers"
	"HTTP_SERVER/middlewares"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Connected to database")

	// users
	http.HandleFunc("/users",
		middlewares.AuthMiddleware(middlewares.LoggerMiddleware(handlers.HandleUsers)),
	)
	// orders
	http.HandleFunc("/orders",
		middlewares.AuthMiddleware(middlewares.LoggerMiddleware(handlers.HandleOrders)),
	)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
