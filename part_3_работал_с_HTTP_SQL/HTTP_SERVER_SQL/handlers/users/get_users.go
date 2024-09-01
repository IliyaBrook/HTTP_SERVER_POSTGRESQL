package users

import (
	"database/sql"
	"encoding/json"
	"errors"
	"main/data"
	"main/pkg"
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
			if errors.Is(err, sql.ErrNoRows) {
				pkg.ResponseErrorText(err, w, "No rows")
				return
			}
			pkg.ResponseErrorText(err, w, "Failed to load users")
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
		pkg.ResponseErrorText(err, w, "Failed to marshal users data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, writeUserErr := w.Write(resp)
	if writeUserErr != nil {
		pkg.ResponseErrorText(writeUserErr, w, "Failed to write response")
	}
}
