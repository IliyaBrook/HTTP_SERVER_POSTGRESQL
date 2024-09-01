package users

import (
	"encoding/json"
	"errors"
	"main/internal/db"
	"main/pkg"
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	updateUserId := r.URL.Query().Get("id")
	var updatedUserData map[string]interface{}

	if updateUserId == "" {
		noIdErr := errors.New("id not found")
		pkg.ResponseErrorText(noIdErr, w, "id not found in URL query")
		return
	}

	err := json.NewDecoder(r.Body).Decode(&updatedUserData)
	defer r.Body.Close()
	if err != nil {
		pkg.ResponseErrorText(err, w, "failed to decode request body")
		return
	}

	query, args, err := pkg.BuildSQLDynamic("UPDATE", "users", updatedUserData, "id = $1", updateUserId)
	if err != nil {
		pkg.ResponseErrorText(err, w, "failed to build update query")
		return
	}

	_, err = db.DB.Exec(query, args...)
	if err != nil {
		pkg.ResponseErrorText(err, w, "Failed to update user")
		return
	}

	pkg.ResponseSuccessText(w, "user successfully updated")
}
