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

	if err != nil {
		resWriter.Write([]byte(": "))
		resWriter.Write([]byte(err.Error()))
	}
}

func ResponseSuccessText(resWriter http.ResponseWriter, message string) {
	resWriter.WriteHeader(http.StatusOK)
	resWriter.Write([]byte(message))
	fmt.Printf("Success: %s\n", message)
}
