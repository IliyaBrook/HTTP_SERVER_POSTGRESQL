package main

import (
	"HTTP_SERVER/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", handlers.HandleUsers)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
