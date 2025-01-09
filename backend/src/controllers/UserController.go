package controllers

import (
	"backend/src/config"
	"backend/src/helpers"
	"backend/src/model"
	"backend/src/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context)  {
	
	var request requests.Login
	
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": "Invalid request" })
		return
	}

	var user model.User
	if err := config.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{ "message": "Invalid Credentials" })
		return
	}

	if user.Password != request.Password {
		c.JSON(http.StatusUnauthorized, gin.H{ "message": "Invalid Credentials" })
		return
	}

	token, err := helpers.JWTGenerator(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ "message": "Could not generate token" })
		return	
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successfully",
		"token": token,
	})
}

func RegisterUser(c *gin.Context)  {
	var request 
}