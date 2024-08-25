package users

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/utils"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser data.UserStruct
	var resp []byte
	var rows *sqlx.Rows
	err := json.NewDecoder(r.Body).Decode(&newUser)
	utils.HandleServerError(err, w, "Failed to marshal orders data", "log")
	rows, err = data.DB.NamedQuery(
		"INSERT INTO users (name, email, password) VALUES (:name, :email, :password) RETURNING id",
		map[string]interface{}{
			"name":     newUser.Name,
			"email":    newUser.Email,
			"password": newUser.Password,
		},
	)

	fmt.Println("create user rows:", rows)
	utils.HandleServerError(err, w, "Failed to save database", "log")

	w.WriteHeader(http.StatusCreated)
	resp, err = json.Marshal(&newUser)
	utils.HandleServerError(err, w, "Failed to marshal user data", "log")

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
