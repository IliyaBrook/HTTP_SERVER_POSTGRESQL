package users

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("id")
	var resp []byte
	var rows *sql.Rows
	var err error
	var users []data.UserStruct
	var user data.UserStruct

	if userId == "" {
		rows, err = data.DB.Query("SELECT * FROM users")
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RegisteredAt)
			if err != nil {
				fmt.Printf("Error scanning row users: %v\n", err)
			}
			users = append(users, user)
		}
		data.DbInst.Users = users
		utils.HandleServerError(err, w, "Failed to load database")
		resp, err = json.Marshal(data.DbInst.Users)
	} else {
		id, _ := strconv.Atoi(userId)
		rows, err = data.DB.Query("SELECT * FROM users WHERE id=$1", id)

		if err != nil {
			fmt.Println("Get users error:", err)
		}
		defer rows.Close()

		if !rows.Next() {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("User not found"))
			return
		} else {
			err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RegisteredAt)
			if err != nil {
				fmt.Printf("Error scanning row users: %v\n", err)
			}
			fmt.Println("log found user:", user)
			resp, err = json.Marshal(user)
		}
	}

	utils.HandleServerError(err, w, "Failed to marshal users data")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, writeUserErr := w.Write(resp)
	utils.HandleServerError(writeUserErr, w, "Failed to write response")
}
