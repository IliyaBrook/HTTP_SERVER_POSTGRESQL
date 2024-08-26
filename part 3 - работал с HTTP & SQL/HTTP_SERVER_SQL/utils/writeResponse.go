package utils

import (
	"fmt"
	"net/http"
)

func ResponseErrorText(err error, resWriter http.ResponseWriter, message string) {
	if message == "" {
		message = "An error occurred"
	}

	resWriter.WriteHeader(http.StatusInternalServerError)
	resWriter.Write([]byte(message))
	fmt.Println(message)
	if err != nil {
		resWriter.Write([]byte(": "))
		resWriter.Write([]byte(err.Error()))
		fmt.Printf("Error: %s\n", err)
	}
}

func ResponseSuccessText(resWriter http.ResponseWriter, message string) {
	resWriter.WriteHeader(http.StatusOK)
	resWriter.Write([]byte(message))
	fmt.Printf("Success: %s\n", message)
}
