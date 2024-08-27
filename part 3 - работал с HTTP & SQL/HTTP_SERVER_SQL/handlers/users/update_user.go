package users

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
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

	if err != nil {
		utils.ResponseErrorText(err, w, "failed to decode request body")
		return
	}

	var queryParts []string
	var args []interface{}

	argID := 1
	for key, value := range updatedUserData {
		queryParts = append(queryParts, fmt.Sprintf("%s = $%d", key, argID))
		args = append(args, value)
		argID++
	}

	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d", strings.Join(queryParts, ", "), argID)
	args = append(args, updateUserId)

	_, err = data.DB.Exec(query, args...)

	if err != nil {
		log.Printf("Failed to update user: %v", err)
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	utils.ResponseSuccessText(w, "user successfully updated")
}
