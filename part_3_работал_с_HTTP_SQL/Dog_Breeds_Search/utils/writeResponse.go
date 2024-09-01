package utils

import (
	"dog_breeds_search/sharable"
	"encoding/json"
	"fmt"
	"net/http"
)

func ResponseErrorText(err error, resWriter http.ResponseWriter, message string) {
	if message == "" {
		message = "An error occurred"
	}

	errorResponse := sharable.ErrorResponse{
		Message: message,
	}

	if err != nil {
		errorResponse.Error = err.Error()
	}

	resWriter.Header().Set("Content-Type", "application/json")
	resWriter.WriteHeader(http.StatusInternalServerError)
	jsonResponse, jsonErr := json.Marshal(errorResponse)
	if jsonErr != nil {
		fmt.Printf("Error marshalling error response: %s\n", jsonErr)
		resWriter.Write([]byte(`{"message": "Internal Server Error"}`))
		return
	}

	resWriter.Write(jsonResponse)
	fmt.Printf("Error: %s\n", err)
}
