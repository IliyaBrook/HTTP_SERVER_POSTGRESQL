package users

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"main/internal/db"
	"main/internal/utils"
	"net/http"
	"strconv"
)

// @Summary Get users
// @Description get users, If no id is provided, returns all users.
// @Tags users
// @Accept  json
// @Produce  json
// @Param id query string false "User ID"
// @Param x-id header string true "X-ID" default(1)
// @Success 200 {array} db.UserStruct "Users"
// @Failure 400 {string} string "Invalid user ID"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Failed to load users"
// @Router /users [get]
func GetUsers(c *gin.Context) {
	userId := c.Query("id")
	var err error

	if userId == "" {
		var users []db.UserStruct

		err = db.DB.Select(&users, "SELECT id, name, email, password, registered_at FROM users")
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				utils.ResponseErrorText(c, err, "No rows")
				return
			}
			utils.ResponseErrorText(c, err, "Failed to load users")
			return
		}

		c.JSON(http.StatusOK, users)
	} else {
		id, err := strconv.Atoi(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		var user db.UserStruct

		err = db.DB.Get(&user, "SELECT id, name, email, password, registered_at FROM users WHERE id=$1", id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
