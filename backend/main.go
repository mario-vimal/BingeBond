package main

import (
	"./src"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	env := os.Getenv("APP_ENV")

	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()

	app = src.InitRoutes(app)

	app.Use(cors.Default())

	//App listening on port 3000!
	app.Run(":3000")
}