package routes

import (
	"github.com/gin-gonic/gin"
	"main/internal/middlewares"
)

func WithMwsRoute(r *gin.Engine) gin.IRoutes {
	route := r.Group("")
	return route.Use(middlewares.AuthMiddleware, middlewares.LoggerMiddleware)
}
