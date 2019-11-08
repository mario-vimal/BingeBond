package main

import (
	"os"

	"./models"
	"./src"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	env := os.Getenv("APP_ENV")

	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	models.InitDb()

	app := gin.Default()

	app = src.InitRoutes(app)

	defer models.CloseDB()

	app.Use(cors.Default())

	//App listening on port 3000!
	app.Run(":3000")
}
