package products

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/utils"
	"encoding/json"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	var ordersData []data.ProductStruct

	if productId != "" {
		err := data.DB.Select(&ordersData, "SELECT * FROM products WHERE id = $1", productId)
		utils.ResponseErrorText(err, w, "Error to get product by id")
	} else {
		err := data.DB.Select(&ordersData, "SELECT * FROM products")
		utils.ResponseErrorText(err, w, "Error to get products")
	}
	jsonData, err := json.Marshal(&ordersData)
	if err != nil {
		utils.ResponseErrorText(err, w, "marshall data failed:")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
