package users

import (
	"HTTP_SERVER/sharable"
	"HTTP_SERVER/utils"
	"encoding/json"
	"io"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	body, readBodyErr := io.ReadAll(r.Body)
	utils.HandleServerError(readBodyErr, w, "Failed to read request body")

	var requestData struct {
		ID int `json:"ID"`
	}

	marshallBodyErr := json.Unmarshal(body, &requestData)
	utils.HandleServerError(marshallBodyErr, w, "Failed to marshal body")

	for i, user := range sharable.DbInst.Users {
		if user.ID == requestData.ID {
			sharable.DbInst.Users = append(sharable.DbInst.Users[:i], sharable.DbInst.Users[i+1:]...)
			break
		}
	}

	saveDbErr := sharable.DbInst.SaveDatabase()
	utils.HandleServerError(saveDbErr, w, "Failed to save database")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}
