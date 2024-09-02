package users

import (
	"errors"
	"github.com/gin-gonic/gin"
	"main/internal/db"
	"main/internal/utils"
	"net/http"
)

func GetUserProducts(c *gin.Context) {
	var userOrders []db.ProductStruct

	var err error
	var userId string

	if userId = c.Query("id"); userId == "" {
		if userId = c.Request.Header.Get("x-id"); userId == "" {
			err = errors.New("failed to get id from body")
			utils.ResponseErrorText(c, err, "failed to get id from body")
			return
		}
	}

	err = db.DB.Select(&userOrders,
		`
		SELECT p.*
		FROM products p
		JOIN user_orders uo ON p.id = uo.product_id
		WHERE uo.user_id = $1
		`,
		userId,
	)

	if err != nil || len(userOrders) == 0 {
		utils.ResponseErrorText(c, err, "not found")
		return
	}

	c.JSON(http.StatusOK, &userOrders)

	utils.ResponseSuccessText(c, "ok")
}
