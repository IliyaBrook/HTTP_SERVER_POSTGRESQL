package handlers

import (
	"HTTP_SERVER/data"
	"encoding/json"
	"io"
	"net/http"
)

var db = &data.DB{}
var user = &data.User{}

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		createUser(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func getUsers(w http.ResponseWriter, _ *http.Request) {
	if err := db.ReadDatabase(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to load database"))
		return
	}

	resp, err := json.Marshal(db.Users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to marshal users data"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser data.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := db.ReadDatabase(); err != nil {
		http.Error(w, "Failed to load database", http.StatusInternalServerError)
	}

	user.ID = len(db.Users) + 1
	db.Users = append(db.Users, newUser)

	if err := db.SaveDatabase(); err != nil {
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

func deleteUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var requestData struct {
		ID int `json:"ID"`
	}

	_ = json.Unmarshal(body, &requestData)

	if err := db.ReadDatabase(); err != nil {
		http.Error(w, "Failed to load database", http.StatusInternalServerError)
	}

	for i, user := range db.Users {
		if user.ID == requestData.ID {
			db.Users = append(db.Users[:i], db.Users[i+1:]...)
			break
		}
	}

	if err := db.SaveDatabase(); err != nil {
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
