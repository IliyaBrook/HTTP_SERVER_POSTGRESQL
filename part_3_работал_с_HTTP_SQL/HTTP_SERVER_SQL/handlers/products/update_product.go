package products

import (
	"encoding/json"
	"errors"
	"main/data"
	"main/pkg"
	"net/http"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	updatedProductId := r.URL.Query().Get("id")
	var updatedProductData map[string]interface{}
	var err error

	if updatedProductId == "" {
		noIdErr := errors.New("id not found")
		pkg.ResponseErrorText(noIdErr, w, "Missing id parameter")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&updatedProductData)
	defer r.Body.Close()
	if err != nil {
		pkg.ResponseErrorText(err, w, "Invalid request body")
		return
	}

	query, args, queryErr := pkg.BuildSQLDynamic("UPDATE", "products", updatedProductData, "id = $1", updatedProductId)
	if err != nil {
		pkg.ResponseErrorText(queryErr, w, "Failed to build update query")
		return
	}

	_, err = data.DB.Exec(query, args...)
	if err != nil {
		pkg.ResponseErrorText(err, w, "Failed to execute update query")
		return
	}

	pkg.ResponseSuccessText(w, "Product updated successfully")
}
