package users

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/sharable"
	"HTTP_SERVER/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("ID")
	readDbErr := sharable.DbInst.ReadDatabase()
	utils.HandleServerError(readDbErr, w, "Failed to load database")

	var resp []byte
	var err error

	if userId == "" {
		resp, err = json.Marshal(sharable.DbInst.Orders)
	} else {
		id, _ := strconv.Atoi(userId)
		//var userFound data.User
		var userFound []data.User
		for _, u := range sharable.DbInst.Users {
			if u.ID == id {
				userFound = append(userFound, u)
			}
		}
		resp, err = json.Marshal(userFound)
	}

	utils.HandleServerError(err, w, "Failed to marshal users data")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, writeUserErr := w.Write(resp)
	utils.HandleServerError(writeUserErr, w, "Failed to write response")
}
