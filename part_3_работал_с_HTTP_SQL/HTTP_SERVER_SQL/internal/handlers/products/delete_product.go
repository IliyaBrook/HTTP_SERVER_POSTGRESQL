package products

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"main/internal/db"
	"main/internal/utils"
)

func DeleteProduct(c *gin.Context) {
	var deletedProduct struct {
		ID int `json:"id" db:"id"`
	}

	tx, err := db.DB.Beginx()
	if err != nil {
		utils.ResponseErrorText(c, err, "failed to begin transaction")
		return
	}

	if err := c.ShouldBindJSON(&deletedProduct); err != nil {
		utils.ResponseErrorText(c, err, "failed to decode request body")
		return
	}
	log.Println("product id from body:", deletedProduct.ID)

	if deletedProduct.ID == 0 {
		utils.ResponseErrorText(c, fmt.Errorf("invalid product ID"), "failed to get deleted product id")
		return
	}

	_, err = tx.NamedExec("DELETE FROM products WHERE id=:id", &deletedProduct)
	if err != nil {
		utils.ResponseErrorText(c, err, "failed to delete product")
		return
	}

	_, err = tx.Exec("DELETE FROM user_orders WHERE product_id=$1", deletedProduct.ID)
	if err != nil {
		utils.ResponseErrorText(c, err, "failed to delete product from user orders")
		return
	}

	if err = tx.Commit(); err != nil {
		utils.ResponseErrorText(c, err, "failed to commit transaction")
		return
	}

	utils.ResponseSuccessText(c, "Product successfully deleted")
}
