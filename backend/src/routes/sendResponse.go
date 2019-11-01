package routes

import (
	"github.com/gin-gonic/gin"
)

// Wrapper function to send a JSON response

func SendResponse(c *gin.Context, status_code int, message interface{}) {

	c.JSON(status_code, gin.H{
		"status_code": status_code,
		"message":     message,
	})

	c.Abort()
	return
}
