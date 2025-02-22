package users

import (
	"HTTP_SERVER/data"
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

	for i, user := range data.DbInst.Users {
		if user.ID == requestData.ID {
			data.DbInst.Users = append(data.DbInst.Users[:i], data.DbInst.Users[i+1:]...)
			break
		}
	}

	saveDbErr := data.DbInst.SaveDatabase()
	utils.HandleServerError(saveDbErr, w, "Failed to save database")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("UserStruct deleted successfully"))
}
