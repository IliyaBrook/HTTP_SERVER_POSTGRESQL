package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {
	userId := c.GetHeader("x-id")
	if userId == "" {
		log.Printf("[%s] %s - error: userID is not provided\n",
			c.Request.Method, c.Request.RequestURI,
		)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.Set("ID", userId)
	c.Next()
}
