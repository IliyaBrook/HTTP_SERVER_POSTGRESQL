package handlers

import (
	"HTTP_SERVER/handlers/orders"
	"HTTP_SERVER/handlers/users"
	"net/http"
)

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users.GetUsers(w, r)
	case http.MethodPost:
		users.CreateUser(w, r)
	case http.MethodDelete:
		users.DeleteUser(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func HandleOrders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		orders.GetOrders(w, r)
	case http.MethodPost:
		// todo
	case http.MethodDelete:
		// todo
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}
