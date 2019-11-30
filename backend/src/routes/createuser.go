package routes

import (
	"fmt"
	"net/http"

	"../../models"
	"github.com/gin-gonic/gin"
)

//Creates a new user
//Todo generate usertoken

func CreateUser(c *gin.Context) {
	db := models.GetDB()
	var user models.UserDetail
	c.BindJSON(&user)
	err := db.Create(&user)
	if err.Error == nil {
		SendResponse(c, http.StatusCreated, "User created successfully!")
	} else {
		fmt.Println(err.Error)
		SendResponse(c, http.StatusBadRequest, "User already registered")
	}
}
