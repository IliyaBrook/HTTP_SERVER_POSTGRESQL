package users

import (
	"encoding/json"
	"main/internal/db"
	"main/pkg"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	var userToDeleteStruct struct {
		ID int `db:"id"`
	}

	var deletedUserId int
	tx, err := db.DB.Beginx()
	defer tx.Rollback()

	err = json.NewDecoder(r.Body).Decode(&userToDeleteStruct)
	defer r.Body.Close()
	if err != nil {
		pkg.ResponseErrorText(err, w, "Failed to marshal body")
		return
	}
	userToDelete := userToDeleteStruct.ID

	err = tx.Get(&deletedUserId, "DELETE FROM users WHERE ID=$1 RETURNING id;", userToDelete)
	if err != nil || deletedUserId != userToDelete {
		pkg.ResponseErrorText(err, w, "delete user failed")
		return
	}

	_, err = tx.Exec("DELETE FROM user_orders WHERE user_id=$1;", deletedUserId)

	if err != nil {
		pkg.ResponseErrorText(err, w, "delete user order failed")
		return
	}

	if err = tx.Commit(); err != nil {
		pkg.ResponseErrorText(err, w, "Failed to commit transaction")
		return
	}

	pkg.ResponseSuccessText(w, "User deleted successfully")
}
