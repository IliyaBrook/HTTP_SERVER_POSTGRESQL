package products

import (
	"encoding/json"
	"fmt"
	"main/data"
	"main/pkg"
	"net/http"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("ID").(string)

	var newProdId int
	var newProductData data.ProductStruct

	err := json.NewDecoder(r.Body).Decode(&newProductData)
	defer r.Body.Close()
	if err != nil {
		pkg.ResponseErrorText(err, w, "failed to decode request body product")
		return
	}

	tx, errTxBegin := data.DB.Beginx()
	defer tx.Rollback()
	if errTxBegin != nil {
		pkg.ResponseErrorText(err, w, "failed to begin transaction")
		return
	}

	err = tx.QueryRowx(
		`INSERT INTO products (Name, Quantity, Price, Description) 
		VALUES ($1, $2, $3, $4) RETURNING id`,
		newProductData.Name, newProductData.Quantity, newProductData.Price, newProductData.Description,
	).Scan(&newProdId)

	if err != nil {
		pkg.ResponseErrorText(err, w, "error to add product")
		return
	}

	if newProdId == 0 {
		pkg.ResponseErrorText(fmt.Errorf("no ID returned"), w, "error to add product")
		return
	}

	rows, errAddProd := tx.Queryx(
		`INSERT INTO user_orders (user_id, product_id) VALUES ($1, $2)`,
		userId, newProdId,
	)
	defer rows.Close()

	if errAddProd != nil {
		pkg.ResponseErrorText(err, w, "error to add product")
		return
	}

	if err = tx.Commit(); err != nil {
		pkg.ResponseErrorText(err, w, "failed to commit transaction")
		return
	}

	pkg.ResponseSuccessText(w, "Product added successfully")
}
