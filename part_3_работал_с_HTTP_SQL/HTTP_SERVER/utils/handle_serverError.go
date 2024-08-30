package utils

import "net/http"

func HandleServerError(err error, resWriter http.ResponseWriter, message string) {
	if err != nil {
		resWriter.WriteHeader(http.StatusInternalServerError)
		resWriter.Write([]byte(message))
		panic(err)
	}
}
