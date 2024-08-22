package utils

import (
	"encoding/json"
)

func SliceToJson[T any](slice []T) []byte {
	jsonBytes, _ := json.Marshal(slice)
	return jsonBytes
}
