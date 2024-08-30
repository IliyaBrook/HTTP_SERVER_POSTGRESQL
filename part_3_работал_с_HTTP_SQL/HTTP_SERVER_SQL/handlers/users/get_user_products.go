package users

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/utils"
	"encoding/json"
	"errors"
	"net/http"
)

func GetUserProducts(w http.ResponseWriter, r *http.Request) {
	var userOrders []data.ProductStruct

	var err error
	var ordersJson []byte
	var userId string

	if userId = r.URL.Query().Get("id"); userId == "" {
		if userId = r.Header.Get("x-id"); userId == "" {
			err = errors.New("failed to get id from body")
			utils.ResponseErrorText(err, w, "failed to get id from body")
			return
		}
	}

	err = data.DB.Select(&userOrders,
		`
		SELECT p.*
		FROM products p
		JOIN user_orders uo ON p.id = uo.product_id
		WHERE uo.user_id = $1
		`,
		userId,
	)

	if err != nil || len(userOrders) == 0 {
		utils.ResponseErrorText(err, w, "not found")
		return
	}

	ordersJson, _ = json.Marshal(&userOrders)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ordersJson)
}
