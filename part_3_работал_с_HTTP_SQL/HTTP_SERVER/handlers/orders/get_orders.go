package orders

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	readOrdersErr := data.DbInst.ReadDatabase()
	utils.HandleServerError(readOrdersErr, w, "Failed to read orders")
	query := r.URL.Query().Get("UserID")

	var resp []byte
	var err error

	if query == "" {
		resp, err = json.Marshal(data.DbInst.Orders)
	} else {
		queryId, _ := strconv.Atoi(query)
		var filterOrders []data.Order
		for _, order := range data.DbInst.Orders {
			if queryId == order.UserID {
				filterOrders = append(filterOrders, order)
			}
		}
		resp, err = json.Marshal(filterOrders)
	}

	utils.HandleServerError(err, w, "Failed to marshal orders data")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
