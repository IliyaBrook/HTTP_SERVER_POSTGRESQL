package pkg

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ResponseErrorText(c *gin.Context, err error, message string) {
	if message == "" {
		message = "An error occurred"
	}

	errorMessage := message
	if err != nil {
		errorMessage = message + ": " + err.Error()
		log.Printf("Error: %s\n", err)
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": errorMessage,
	})
}

func ResponseSuccessText(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
	log.Printf("Success: %s\n", message)
}
