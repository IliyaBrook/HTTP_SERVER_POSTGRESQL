package users

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/utils"
	"encoding/json"
	"errors"
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	updateUserId := r.URL.Query().Get("id")
	var updatedUserData map[string]interface{}

	if updateUserId == "" {
		noIdErr := errors.New("id not found")
		utils.ResponseErrorText(noIdErr, w, "id not found in URL query")
		return
	}

	err := json.NewDecoder(r.Body).Decode(&updatedUserData)
	defer r.Body.Close()
	if err != nil {
		utils.ResponseErrorText(err, w, "failed to decode request body")
		return
	}

	query, args, err := utils.BuildSQLDynamic("UPDATE", "users", updatedUserData, "id = $1", updateUserId)
	if err != nil {
		utils.ResponseErrorText(err, w, "failed to build update query")
		return
	}

	_, err = data.DB.Exec(query, args...)
	if err != nil {
		utils.ResponseErrorText(err, w, "Failed to update user")
		return
	}

	utils.ResponseSuccessText(w, "user successfully updated")
}
