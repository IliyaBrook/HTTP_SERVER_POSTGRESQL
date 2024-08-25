package utils

import (
	"fmt"
	"log"
	"net/http"
)

// HandleServerError handles server errors by logging the error, returning an appropriate HTTP response,
// The third parameter is a function behavior that can use fmt.Printf if you pass "log" as this argument, by default use log.Fatal
func HandleServerError(err error, resWriter http.ResponseWriter, args ...string) {
	message := "An error occurred"
	behavior := "fatal"
	if len(args) > 0 {
		message = args[0]
	}
	if len(args) > 1 {
		behavior = args[1]
	}
	if err != nil {
		resWriter.WriteHeader(http.StatusInternalServerError)
		resWriter.Write([]byte(message))
		if behavior == "" || behavior == "fatal" {
			log.Fatal(err)
		} else {
			fmt.Printf("%s\nInternal error: %s\n", message, err)
		}
	}
}
