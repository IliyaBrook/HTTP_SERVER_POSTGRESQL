package users

import (
	"encoding/json"
	"main/internal/db"
	"main/pkg"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser db.UserStruct
	err := json.NewDecoder(r.Body).Decode(&newUser)
	defer r.Body.Close()
	if err != nil {
		pkg.ResponseErrorText(err, w, "Failed to marshal orders data")
	}

	rows, errInsert := db.DB.NamedQuery(
		"INSERT INTO users (name, email, password) VALUES (:name, :email, :password) RETURNING id",
		map[string]interface{}{
			"name":     newUser.Name,
			"email":    newUser.Email,
			"password": newUser.Password,
		},
	)
	defer rows.Close()

	if errInsert != nil {
		pkg.ResponseErrorText(err, w, "Failed to create user")
		return
	}

	w.Write([]byte("User created successfully."))
}
