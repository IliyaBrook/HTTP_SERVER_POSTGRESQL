package handlers

import (
	"main/handlers/products"
	"main/handlers/users"
	"net/http"
)

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users.GetUsers(w, r)
	case http.MethodPatch:
		users.UpdateUser(w, r)
	case http.MethodPost:
		users.CreateUser(w, r)
	case http.MethodDelete:
		users.DeleteUser(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func HandleUserProducts(w http.ResponseWriter, r *http.Request) {
	if http.MethodGet == r.Method {
		users.GetUserProducts(w, r)
	}
}

func HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		products.GetProducts(w, r)
	case http.MethodPost:
		products.AddProduct(w, r)
	case http.MethodPatch:
		products.UpdateProduct(w, r)
	case http.MethodDelete:
		products.DeleteProduct(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}
