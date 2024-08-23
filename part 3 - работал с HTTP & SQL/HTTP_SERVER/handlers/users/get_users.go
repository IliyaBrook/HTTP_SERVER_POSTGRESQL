package users

import (
	"HTTP_SERVER/sharable"
	"HTTP_SERVER/utils"
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	readDbErr := sharable.DbInst.ReadDatabase()
	utils.HandleServerError(readDbErr, w, "Failed to load database")

	resp, marshalUserErr := json.Marshal(sharable.DbInst)
	utils.HandleServerError(marshalUserErr, w, "Failed to marshal users data")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, writeUserErr := w.Write(resp)
	utils.HandleServerError(writeUserErr, w, "Failed to write response")
}
