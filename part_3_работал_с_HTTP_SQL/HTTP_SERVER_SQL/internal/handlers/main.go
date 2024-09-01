package handlers

import (
	products2 "main/internal/handlers/products"
	users2 "main/internal/handlers/users"
	"net/http"
)

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users2.GetUsers(w, r)
	case http.MethodPatch:
		users2.UpdateUser(w, r)
	case http.MethodPost:
		users2.CreateUser(w, r)
	case http.MethodDelete:
		users2.DeleteUser(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func HandleUserProducts(w http.ResponseWriter, r *http.Request) {
	if http.MethodGet == r.Method {
		users2.GetUserProducts(w, r)
	}
}

func HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		products2.GetProducts(w, r)
	case http.MethodPost:
		products2.AddProduct(w, r)
	case http.MethodPatch:
		products2.UpdateProduct(w, r)
	case http.MethodDelete:
		products2.DeleteProduct(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}
