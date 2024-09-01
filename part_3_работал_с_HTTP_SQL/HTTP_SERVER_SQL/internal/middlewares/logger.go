package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LoggerMiddleware(c *gin.Context) {
	userId, _ := c.Get("ID")
	method := c.Request.Method
	url := c.Request.URL
	if userId == "" {
		log.Printf("[%s] %s - error: userID is invalid", method, url)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	log.Printf("[%s] %s by userId %s\n", method, url, userId)
	c.Next()
}
