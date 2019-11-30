package src

import (
	"../src/routes"
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) *gin.Engine {

	api := app.Group("/api")
	{
		api.POST("/sample", routes.Sample)
		api.POST("/createuser", routes.CreateUser)
		api.POST("/login", routes.LoginUser)

	}

	return app
}
