package products

import (
	"encoding/json"
	"log"
	"main/data"
	"main/pkg"
	"net/http"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var deletedProduct struct {
		ID int `db:"id"`
	}

	tx, err := data.DB.Beginx()
	defer tx.Rollback()

	if err != nil {
		pkg.ResponseErrorText(err, w, "failed to begin transaction")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&deletedProduct)
	defer r.Body.Close()
	log.Println("product id from body:", deletedProduct.ID)
	if err != nil || deletedProduct.ID == 0 {
		pkg.ResponseErrorText(err, w, "failed to get deleted product id")
		return
	}

	_, err = tx.NamedExec("DELETE FROM products WHERE id=:id", &deletedProduct)
	if err != nil {
		pkg.ResponseErrorText(err, w, "failed to delete product")
		return
	}

	_, err = tx.Exec("DELETE FROM user_orders WHERE product_id=$1", &deletedProduct.ID)
	if err != nil {
		pkg.ResponseErrorText(err, w, "failed to delete product")
		return
	}

	if err = tx.Commit(); err != nil {
		pkg.ResponseErrorText(err, w, "failed to commit transaction")
		return
	}

	pkg.ResponseSuccessText(w, "Product successfully deleted")
}
