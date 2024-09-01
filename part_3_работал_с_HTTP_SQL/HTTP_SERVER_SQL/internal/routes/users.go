package routes

import (
	"github.com/gin-gonic/gin"
	"main/internal/handlers/users"
)

func RegisterUserRoutes(r *gin.Engine) {
	route := WithMwsRoute(r)
	{
		route.GET("/users", users.GetUsers)
		route.POST("/users", users.CreateUser)
		route.PATCH("/users", users.UpdateUser)
		route.DELETE("/users", users.DeleteUser)
		route.GET("/user-products", users.GetUserProducts)
	}
}
