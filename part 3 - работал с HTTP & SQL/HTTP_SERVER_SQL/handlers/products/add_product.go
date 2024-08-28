package products

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("ID").(string)

	var newProdId int
	var newProductData data.ProductStruct

	err := json.NewDecoder(r.Body).Decode(&newProductData)
	defer r.Body.Close()
	if err != nil {
		utils.ResponseErrorText(err, w, "failed to decode request body product")
		return
	}

	tx, err := data.DB.Beginx()
	if err != nil {
		utils.ResponseErrorText(err, w, "failed to begin transaction")
		return
	}

	err = tx.QueryRowx(
		`INSERT INTO products (Name, Quantity, Price, Description) 
		VALUES ($1, $2, $3, $4) RETURNING id`,
		newProductData.Name, newProductData.Quantity, newProductData.Price, newProductData.Description,
	).Scan(&newProdId)

	if err != nil {
		_ = tx.Rollback()
		utils.ResponseErrorText(err, w, "error to add product")
		return
	}

	if newProdId == 0 {
		_ = tx.Rollback()
		utils.ResponseErrorText(fmt.Errorf("no ID returned"), w, "error to add product")
		return
	}

	_, err = tx.Queryx(
		`INSERT INTO user_orders (user_id, product_id) VALUES ($1, $2)`,
		userId, newProdId,
	)

	if err != nil {
		_ = tx.Rollback()
		utils.ResponseErrorText(err, w, "error to add product")
		return
	}

	if err = tx.Commit(); err != nil {
		utils.ResponseErrorText(err, w, "failed to commit transaction")
		return
	}

	utils.ResponseSuccessText(w, "Product added successfully")
}
