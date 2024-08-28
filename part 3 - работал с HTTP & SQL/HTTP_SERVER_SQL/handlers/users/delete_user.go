package users

import (
	"HTTP_SERVER/data"
	"HTTP_SERVER/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	var userToDeleteStruct struct {
		ID int `db:"id"`
	}

	var deletedUserId int
	tx, err := data.DB.Beginx()

	err = json.NewDecoder(r.Body).Decode(&userToDeleteStruct)
	defer r.Body.Close()
	if err != nil {
		utils.ResponseErrorText(err, w, "Failed to marshal body")
		return
	}
	userToDelete := userToDeleteStruct.ID
	fmt.Println("user To Delete", userToDelete)

	err = tx.Get(&deletedUserId, "DELETE FROM users WHERE ID=$1 RETURNING id;", userToDelete)
	if err != nil || deletedUserId != userToDelete {
		_ = tx.Rollback()
		utils.ResponseErrorText(err, w, "delete user failed")
		return
	}

	_, err = tx.Exec("DELETE FROM user_orders WHERE user_id=$1;", deletedUserId)

	if err != nil {
		_ = tx.Rollback()
		utils.ResponseErrorText(err, w, "delete user order failed")
		return
	}

	if err = tx.Commit(); err != nil {
		utils.ResponseErrorText(err, w, "Failed to commit transaction")
		return
	}

	utils.ResponseSuccessText(w, "User deleted successfully")
}
