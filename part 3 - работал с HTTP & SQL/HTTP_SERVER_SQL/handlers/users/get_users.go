package users

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("id")
	var err error
	var resp []byte

	if userId == "" {
		var users []data.UserStruct

		err = data.DB.Select(&users, "SELECT id, name, email, password, registered_at FROM users")
		if err != nil {
			utils.HandleServerError(err, w, "Failed to load users", "test")
			return
		}

		data.DbInst.Users = users
		resp, err = json.Marshal(data.DbInst.Users)
	} else {
		id, err := strconv.Atoi(userId)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		var user data.UserStruct

		err = data.DB.Get(&user, "SELECT id, name, email, password, registered_at FROM users WHERE id=$1", id)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		resp, err = json.Marshal(user)
	}

	if err != nil {
		utils.HandleServerError(err, w, "Failed to marshal users data", "test")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, writeUserErr := w.Write(resp)
	utils.HandleServerError(writeUserErr, w, "Failed to write response", "test")
}
