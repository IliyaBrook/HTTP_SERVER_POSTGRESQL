package users

import (
	"github.com/gin-gonic/gin"
	"main/internal/db"
	"main/internal/utils"
)

func CreateUser(c *gin.Context) {
	var newUser db.UserStruct
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		utils.ResponseErrorText(c, err, "Failed to marshal orders data")
	}

	rows, errInsert := db.DB.NamedQuery(
		"INSERT INTO users (name, email, password) VALUES (:name, :email, :password) RETURNING id",
		map[string]interface{}{
			"name":     newUser.Name,
			"email":    newUser.Email,
			"password": newUser.Password,
		},
	)
	defer rows.Close()

	if errInsert != nil {
		utils.ResponseErrorText(c, err, "Failed to create user")
		return
	}

	utils.ResponseSuccessText(c, "User created successfully.")
}
