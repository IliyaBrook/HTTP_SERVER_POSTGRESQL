package utils

import (
	"assignment/sharable"
	"encoding/json"
	"io"
	"net/http"
)

func QueryGet(w http.ResponseWriter, pathSegment string) error {
	var isError = false

	url := sharable.ApiUrl + pathSegment
	client, err := http.DefaultClient.Get(url)
	if err != nil {
		ResponseErrorText(err, w, "Error while getting"+pathSegment)
		isError = true
	}
	defer client.Body.Close()
	body, readAllErr := io.ReadAll(client.Body)
	if readAllErr != nil {
		ResponseErrorText(err, w, "Error while getting"+pathSegment+"from body")
		isError = true
	}
	err = json.Unmarshal(body, &sharable.Dogs)
	if err != nil {
		ResponseErrorText(err, w, "Failed to unmarshal"+pathSegment)
		isError = true
	}

	return isError
}
