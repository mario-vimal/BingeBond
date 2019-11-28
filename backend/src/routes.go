package src

import (
	"../src/routes"
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) *gin.Engine {

	api := app.Group("/api")
	{
		api.POST("/sample", routes.Sample)
	}

	return app
}
