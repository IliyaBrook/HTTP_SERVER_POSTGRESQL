package users

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/sharable"
	"HTTP_SERVER/utils"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser data.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	instUser := &data.User{}
	utils.HandleServerError(err, w, "Failed to marshal orders data")

	instUser.ID = len(sharable.DbInst.Users) + 1
	sharable.DbInst.Users = append(sharable.DbInst.Users, newUser)

	err = sharable.DbInst.SaveDatabase()
	utils.HandleServerError(err, w, "Failed to save database")

	w.WriteHeader(http.StatusCreated)
	resp, err := json.Marshal(&newUser)
	utils.HandleServerError(err, w, "Failed to marshal user data")

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
