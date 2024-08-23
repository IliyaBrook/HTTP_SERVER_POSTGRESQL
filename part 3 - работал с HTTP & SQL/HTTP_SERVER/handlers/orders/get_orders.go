package orders

import (
	"HTTP_SERVER/handlers"
	"encoding/json"
	"net/http"
)

func GetOrders(w http.ResponseWriter, _ *http.Request) {
	if err := handlers.DbInst.ReadDatabase(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	resp, err := json.Marshal(handlers.DbInst.Orders)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to marshal orders data"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
