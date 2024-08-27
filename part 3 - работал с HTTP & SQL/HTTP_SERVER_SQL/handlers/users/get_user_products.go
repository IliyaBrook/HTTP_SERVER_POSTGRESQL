package users

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/utils"
	"encoding/json"
	"net/http"
)

func GetUserProducts(w http.ResponseWriter, r *http.Request) {
	var userId struct {
		UserId string `db:"id"`
	}
	var userOrders data.ProductStruct

	tx, err := data.DB.Beginx()

	err = json.NewDecoder(r.Body).Decode(&userId)
	if err != nil {
		utils.ResponseErrorText(err, w, "failed to get id from body")
		return
	}

	err = tx.Select(&userOrders, "")
}
