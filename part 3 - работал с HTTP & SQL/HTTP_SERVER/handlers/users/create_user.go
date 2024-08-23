package users

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/handlers"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser data.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := handlers.DbInst.ReadDatabase(); err != nil {
		http.Error(w, "Failed to load database", http.StatusInternalServerError)
	}

	handlers.UserInst.ID = len(handlers.DbInst.Users) + 1
	handlers.DbInst.Users = append(handlers.DbInst.Users, newUser)

	if err := handlers.DbInst.SaveDatabase(); err != nil {
		http.Error(w, "Failed to save database", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	resp, err := json.Marshal(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to marshal user data"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
