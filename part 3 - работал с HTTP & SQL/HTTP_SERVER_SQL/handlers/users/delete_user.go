package users

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	var requestData struct {
		ID int `json:"id" db:"id"`
	}

	marshalErr := json.NewDecoder(r.Body).Decode(&requestData)
	if marshalErr != nil {
		utils.ResponseErrorText(marshalErr, w, "Failed to marshal body")
	}

	result, err := data.DB.NamedExec("DELETE FROM users WHERE ID=:id", &requestData)
	if err != nil {
		fmt.Println("Delete user failed", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		utils.ResponseErrorText(err, w, "Failed to check rows affected")
		return
	}

	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}
