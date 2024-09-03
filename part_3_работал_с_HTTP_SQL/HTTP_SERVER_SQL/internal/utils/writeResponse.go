package utils

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func ResponseErrorText(c *gin.Context, err error, message string) {
	if message == "" {
		message = "An error occurred"
	}

	errorMessage := message
	if err != nil {
		errorMessage = message + ": " + err.Error()
		log.Error("Error: %s\n", err)
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": errorMessage,
	})
}

func ResponseSuccessText(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
	log.Info("Success: %s\n", message)
}
