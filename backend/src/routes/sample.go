package routes

import (
	"github.com/gin-gonic/gin"
)

//Sample file for all routes

func Sample(c *gin.Context) {
	SendResponse(c, 200, "HOLA")
}
