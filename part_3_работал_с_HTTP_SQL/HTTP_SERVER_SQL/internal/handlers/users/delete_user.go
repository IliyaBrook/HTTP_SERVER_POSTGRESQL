package users

import (
	"github.com/gin-gonic/gin"
	"main/internal/db"
	"main/internal/utils"
)

// @Summary Delete user
// @Description delete user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id body int true "User ID"
// @Param x-id header string true "X-ID" default(1)
// @Success 200 {string} string "User deleted successfully"
// @Failure 400 {string} string "Failed to marshal body"
// @Failure 500 {string} string "Failed to commit transaction"
// @Router /users [delete]
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
