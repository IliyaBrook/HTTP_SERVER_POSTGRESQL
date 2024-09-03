package products

import (
	"github.com/gin-gonic/gin"
	"main/internal/db"
	"main/internal/utils"
	"net/http"
)

// @Summary Get products
// @Description get products. If no id is provided, returns all products.
// @Tags products
// @Accept  json
// @Produce  json
// @Param id query string false "Product ID"
// @Param x-id header string true "X-ID" default(1)
// @Success 200 {array} db.ProductStruct "Products"
// @Failure 400 {string} string "Error to get product by id"
// @Failure 500 {string} string "Error to get products"
// @Router /products [get]
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
