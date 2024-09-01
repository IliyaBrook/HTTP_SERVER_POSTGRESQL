package utils

import (
	"dog_breeds_search/sharable"
	"encoding/json"
	"io"
	"net/http"
)

func QueryGet[T any](w http.ResponseWriter, pathSegment string, target T) bool {
	url := sharable.ApiUrl + pathSegment
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ResponseErrorText(err, w, "Error while creating request for "+pathSegment)
		return false
	}
	req.Header.Set("x-api-key", sharable.ApiKey)
	client := &http.Client{}
	resp, errClientDo := client.Do(req)
	if err != nil {
		ResponseErrorText(errClientDo, w, "Error while getting "+pathSegment)
		return false
	}
	defer resp.Body.Close()
	body, readAllErr := io.ReadAll(resp.Body)

	if readAllErr != nil {
		ResponseErrorText(err, w, "Error while getting "+pathSegment+" from body")
		return false
	}
	err = json.Unmarshal(body, &target)
	if err != nil {
		ResponseErrorText(err, w, "Failed to unmarshal "+pathSegment)
		return false
	}

	return true
}
