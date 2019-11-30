package routes

import (
	"net/http"

	"../../models"
	"github.com/gin-gonic/gin"
)

//Login route

func LoginUser(c *gin.Context) {
	db := models.GetDB()
	var user models.UserDetail
	var login models.Loginmodel
	c.BindJSON(&login)
	err := db.Where("email = ?", login.Email).First(&user)
	if err.Error != nil {
		SendResponse(c, http.StatusBadRequest, "User not registered")
		return
	}
	if user.Password == login.Password {
		SendResponse(c, http.StatusCreated, "Logged in successfully!")
	} else {
		SendResponse(c, http.StatusForbidden, "Wrong password")
	}

}
