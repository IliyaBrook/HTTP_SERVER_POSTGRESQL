package users

import (
	"github.com/gin-gonic/gin"
	"main/internal/db"
	"main/internal/utils"
)

// @Summary Create user
// @Description create user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body db.UserStruct true "User data"
// @Param x-id header string true "X-ID" default(1)
// @Success 200 {string} string "User created successfully"
// @Failure 400 {string} string "Failed to create user"
// @Router /users [post]
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
