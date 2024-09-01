package products

import (
	"encoding/json"
	"main/internal/db"
	"main/pkg"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	var ordersData []db.ProductStruct
	var err error
	var jsonResData []byte

	if productId != "" {
		err = db.DB.Select(&ordersData, "SELECT * FROM products WHERE id = $1", productId)
		if err != nil {
			pkg.ResponseErrorText(err, w, "Error to get product by id")
			return
		}
	} else {
		err = db.DB.Select(&ordersData, "SELECT * FROM products")
		if err != nil {
			pkg.ResponseErrorText(err, w, "Error to get products")
			return
		}
	}
	jsonResData, err = json.Marshal(&ordersData)
	if err != nil {
		pkg.ResponseErrorText(err, w, "marshall data failed:")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResData)
}
