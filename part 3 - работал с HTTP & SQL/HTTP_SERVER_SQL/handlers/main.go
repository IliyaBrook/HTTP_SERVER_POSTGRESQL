package handlers

import (
	"HTTP_SERVER/handlers/products"
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

func HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		products.GetProducts(w, r)
	case http.MethodPost:
		products.AddProduct(w, r)
	case http.MethodDelete:
		products.DeleteProduct(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}
