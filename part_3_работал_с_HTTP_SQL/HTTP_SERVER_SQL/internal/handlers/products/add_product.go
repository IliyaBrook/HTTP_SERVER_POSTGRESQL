package products

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/internal/db"
	"main/pkg"
)

func AddProduct(c *gin.Context) {
	userId, exists := c.Get("ID")
	if !exists {
		pkg.ResponseErrorText(c, fmt.Errorf("user ID not found"), "user ID is missing")
		return
	}
	userIdStr, ok := userId.(string)
	if !ok {
		pkg.ResponseErrorText(c, fmt.Errorf("user ID is not a valid string"), "user ID is invalid")
		return
	}

	var newProductData db.ProductStruct
	if err := c.ShouldBindJSON(&newProductData); err != nil {
		pkg.ResponseErrorText(c, err, "failed to decode request body product")
		return
	}

	tx, err := db.DB.Beginx()
	if err != nil {
		pkg.ResponseErrorText(c, err, "failed to begin transaction")
		return
	}
	defer tx.Rollback()

	var newProdId int
	err = tx.QueryRowx(
		`INSERT INTO products (Name, Quantity, Price, Description) 
		VALUES ($1, $2, $3, $4) RETURNING id`,
		newProductData.Name, newProductData.Quantity, newProductData.Price, newProductData.Description,
	).Scan(&newProdId)

	if err != nil {
		pkg.ResponseErrorText(c, err, "error to add product")
		return
	}

	if newProdId == 0 {
		pkg.ResponseErrorText(c, fmt.Errorf("no ID returned"), "error to add product")
		return
	}

	_, err = tx.Exec(
		`INSERT INTO user_orders (user_id, product_id) VALUES ($1, $2)`,
		userIdStr, newProdId,
	)

	if err != nil {
		pkg.ResponseErrorText(c, err, "error to add product")
		return
	}

	if err = tx.Commit(); err != nil {
		pkg.ResponseErrorText(c, err, "failed to commit transaction")
		return
	}

	pkg.ResponseSuccessText(c, "Product added successfully")
}
