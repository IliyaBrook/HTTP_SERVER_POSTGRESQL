package users

import (
	"HTTP_SERVER/handlers"
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	if err := handlers.DbInst.ReadDatabase(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to load database"))
		return
	}

	resp, err := json.Marshal(handlers.DbInst)
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
