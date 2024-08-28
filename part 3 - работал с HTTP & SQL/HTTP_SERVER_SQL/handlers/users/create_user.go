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
	_, err = data.DB.NamedQuery(
		"INSERT INTO users (name, email, password) VALUES (:name, :email, :password) RETURNING id",
		map[string]interface{}{
			"name":     newUser.Name,
			"email":    newUser.Email,
			"password": newUser.Password,
		},
	)
	if err != nil {
		utils.ResponseErrorText(err, w, "Failed to save database")
		return
	}

	w.Write([]byte("User created successfully."))
}
