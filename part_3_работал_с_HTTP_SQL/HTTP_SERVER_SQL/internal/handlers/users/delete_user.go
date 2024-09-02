package users

import (
	"github.com/gin-gonic/gin"
	"main/internal/db"
	"main/internal/utils"
)

func DeleteUser(c *gin.Context) {

	var userToDeleteStruct struct {
		ID int `db:"id"`
	}

	var deletedUserId int
	tx, err := db.DB.Beginx()
	defer tx.Rollback()

	err = c.ShouldBindJSON(&userToDeleteStruct)
	if err != nil {
		utils.ResponseErrorText(c, err, "Failed to marshal body")
		return
	}
	userToDelete := userToDeleteStruct.ID

	err = tx.Get(&deletedUserId, "DELETE FROM users WHERE ID=$1 RETURNING id;", userToDelete)
	if err != nil || deletedUserId != userToDelete {
		utils.ResponseErrorText(c, err, "delete user failed")
		return
	}

	_, err = tx.Exec("DELETE FROM user_orders WHERE user_id=$1;", deletedUserId)

	if err != nil {
		utils.ResponseErrorText(c, err, "delete user order failed")
		return
	}

	if err = tx.Commit(); err != nil {
		utils.ResponseErrorText(c, err, "Failed to commit transaction")
		return
	}

	utils.ResponseSuccessText(c, "User deleted successfully")
}
