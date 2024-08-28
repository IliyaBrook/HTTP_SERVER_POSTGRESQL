package products

import "net/http"

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	updatedUserId := r.URL.Query().Get("id")

	if updatedUserId == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
}
