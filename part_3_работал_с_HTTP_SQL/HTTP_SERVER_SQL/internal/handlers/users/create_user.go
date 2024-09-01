package users

import (
	"github.com/gin-gonic/gin"
	"main/internal/db"
	"main/pkg"
)

func CreateUser(c *gin.Context) {
	var newUser db.UserStruct
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		pkg.ResponseErrorText(c, err, "Failed to marshal orders data")
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
		pkg.ResponseErrorText(c, err, "Failed to create user")
		return
	}

	pkg.ResponseSuccessText(c, "User created successfully.")
}
