package src

import (
	"../handler"
	"../src/routes"
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) *gin.Engine {

	api := app.Group("/api")
	{
		api.POST("/sample", routes.Sample)
		api.POST("/query", handler.GraphqlHandler())
	}

	return app
}
