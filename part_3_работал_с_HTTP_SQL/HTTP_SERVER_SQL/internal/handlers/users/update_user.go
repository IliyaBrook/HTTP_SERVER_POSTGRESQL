package users

import (
	"errors"
	"github.com/gin-gonic/gin"
	"main/internal/db"
	"main/pkg"
)

func UpdateUser(c *gin.Context) {
	updateUserId := c.Query("id")
	var updatedUserData map[string]interface{}

	if updateUserId == "" {
		noIdErr := errors.New("id not found")
		pkg.ResponseErrorText(c, noIdErr, "id not found in URL query")
		return
	}

	if err := c.ShouldBindJSON(&updatedUserData); err != nil {
		pkg.ResponseErrorText(c, err, "failed to decode request body")
		return
	}

	query, args, err := pkg.BuildSQLDynamic("UPDATE", "users", updatedUserData, "id = $1", updateUserId)
	if err != nil {
		pkg.ResponseErrorText(c, err, "failed to build update query")
		return
	}

	_, err = db.DB.Exec(query, args...)
	if err != nil {
		pkg.ResponseErrorText(c, err, "Failed to update user")
		return
	}

	pkg.ResponseSuccessText(c, "User successfully updated")
}
