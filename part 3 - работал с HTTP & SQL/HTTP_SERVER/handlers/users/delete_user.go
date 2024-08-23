package users

import (
	"HTTP_SERVER/handlers"
	"encoding/json"
	"io"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var requestData struct {
		ID int `json:"ID"`
	}

	_ = json.Unmarshal(body, &requestData)

	if err := handlers.DbInst.ReadDatabase(); err != nil {
		http.Error(w, "Failed to load database", http.StatusInternalServerError)
	}

	for i, user := range handlers.DbInst.Users {
		if user.ID == requestData.ID {
			handlers.DbInst.Users = append(handlers.DbInst.Users[:i], handlers.DbInst.Users[i+1:]...)
			break
		}
	}

	if err := handlers.DbInst.SaveDatabase(); err != nil {
		http.Error(w, "Failed to save database", http.StatusInternalServerError)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to delete user from the database"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}
