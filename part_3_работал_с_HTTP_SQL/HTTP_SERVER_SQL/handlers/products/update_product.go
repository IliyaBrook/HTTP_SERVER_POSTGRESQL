package products

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/utils"
	"encoding/json"
	"errors"
	"net/http"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	updatedProductId := r.URL.Query().Get("id")
	var updatedProductData map[string]interface{}
	var err error

	if updatedProductId == "" {
		noIdErr := errors.New("id not found")
		utils.ResponseErrorText(noIdErr, w, "Missing id parameter")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&updatedProductData)
	defer r.Body.Close()
	if err != nil {
		utils.ResponseErrorText(err, w, "Invalid request body")
		return
	}

	query, args, queryErr := utils.BuildSQLDynamic("UPDATE", "products", updatedProductData, "id = $1", updatedProductId)
	if err != nil {
		utils.ResponseErrorText(queryErr, w, "Failed to build update query")
		return
	}

	_, err = data.DB.Exec(query, args...)
	if err != nil {
		utils.ResponseErrorText(err, w, "Failed to execute update query")
		return
	}

	utils.ResponseSuccessText(w, "Product updated successfully")
}
