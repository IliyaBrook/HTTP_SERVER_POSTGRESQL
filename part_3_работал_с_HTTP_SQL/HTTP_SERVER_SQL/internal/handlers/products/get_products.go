package products

import (
	"github.com/gin-gonic/gin"
	"main/internal/db"
	"main/internal/utils"
	"net/http"
)

func GetProducts(c *gin.Context) {
	productId := c.Query("id")
	var ordersData []db.ProductStruct
	var err error

	if productId != "" {
		err = db.DB.Select(&ordersData, "SELECT * FROM products WHERE id = $1", productId)
		if err != nil {
			utils.ResponseErrorText(c, err, "Error to get product by id")
			return
		}
	} else {
		err = db.DB.Select(&ordersData, "SELECT * FROM products")
		if err != nil {
			utils.ResponseErrorText(c, err, "Error to get products")
			return
		}
	}

	c.JSON(http.StatusOK, ordersData)
}
