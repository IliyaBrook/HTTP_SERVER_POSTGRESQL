package products

import (
	"errors"
	"github.com/gin-gonic/gin"
	"main/internal/db"
	"main/internal/utils"
)

// @Summary Update product
// @Description update product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id query string true "Product ID"
// @Param x-id header string true "X-ID" default(1)
// @Success 200 {string} string "Product updated successfully"
// @Failure 400 {string} string "Missing id parameter"
// @Failure 500 {string} string "Failed to execute update query"
// @Router /products [patch]
func UpdateProduct(c *gin.Context) {
	updatedProductId := c.Query("id")
	var updatedProductData map[string]interface{}
	var err error

	if updatedProductId == "" {
		noIdErr := errors.New("id not found")
		utils.ResponseErrorText(c, noIdErr, "Missing id parameter")
		return
	}

	err = c.ShouldBindJSON(&updatedProductData)

	query, args, queryErr := utils.BuildSQLDynamic("UPDATE", "products", updatedProductData, "id = $1", updatedProductId)
	if err != nil {
		utils.ResponseErrorText(c, queryErr, "Failed to build update query")
		return
	}

	_, err = db.DB.Exec(query, args...)
	if err != nil {
		utils.ResponseErrorText(c, err, "Failed to execute update query")
		return
	}

	utils.ResponseSuccessText(c, "Product updated successfully")
}
