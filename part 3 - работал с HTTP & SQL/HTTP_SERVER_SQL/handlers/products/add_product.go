package products

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type newProductDataT struct {
	Name        string  `json:"name" db:"name"`
	Quantity    int     `json:"quantity" db:"quantity"`
	Price       float64 `json:"price" db:"price"`
	Description string  `json:"description" db:"description"`
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("ID").(string)

	fmt.Println("user id:", userId)

	var newProdId int
	var newProductData newProductDataT

	err := json.NewDecoder(r.Body).Decode(&newProductData)
	if err != nil {
		utils.ResponseErrorText(err, w, "failed to decode request body product")
		return
	}
	fmt.Println("before insert")

	insertErr := data.DB.QueryRowx(
		`INSERT INTO products (Name, Quantity, Price, Description) 
		VALUES ($1, $2, $3, $4) RETURNING id`,
		newProductData.Name, newProductData.Quantity, newProductData.Price, newProductData.Description,
	).Scan(&newProdId)

	if insertErr != nil {
		utils.ResponseErrorText(insertErr, w, "error to add product")
		return
	}

	if newProdId == 0 {
		utils.ResponseErrorText(fmt.Errorf("no ID returned"), w, "error to add product")
		return
	}

	utils.ResponseSuccessText(w, "Product added successfully")
}
