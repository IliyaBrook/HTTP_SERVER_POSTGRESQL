package users

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/utils"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser data.UserStruct
	err := json.NewDecoder(r.Body).Decode(&newUser)
	defer r.Body.Close()
	if err != nil {
		utils.ResponseErrorText(err, w, "Failed to marshal orders data")
	}

	rows, errInsert := data.DB.NamedQuery(
		"INSERT INTO users (name, email, password) VALUES (:name, :email, :password) RETURNING id",
		map[string]interface{}{
			"name":     newUser.Name,
			"email":    newUser.Email,
			"password": newUser.Password,
		},
	)
	defer rows.Close()

	if errInsert != nil {
		utils.ResponseErrorText(err, w, "Failed to create user")
		return
	}

	w.Write([]byte("User created successfully."))
}
