package users

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"main/internal/db"
	"main/pkg"
	"net/http"
	"strconv"
)

func GetUsers(c *gin.Context) {
	userId := c.Query("id")
	var err error

	if userId == "" {
		var users []db.UserStruct

		err = db.DB.Select(&users, "SELECT id, name, email, password, registered_at FROM users")
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				pkg.ResponseErrorText(c, err, "No rows")
				return
			}
			pkg.ResponseErrorText(c, err, "Failed to load users")
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
