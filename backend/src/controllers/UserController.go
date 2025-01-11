package controllers

import (
	"backend/domain/requests"
	"backend/domain/services"
	"backend/lib/helpers"
	"backend/lib/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginUser(c *gin.Context) {

	var request requests.Login

	if err := c.ShouldBindJSON(&request); err != nil {
		response.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"error": err.Error(),
		})
		return
	}

	user, err := services.Login(request.Email, request.Password)
	if err != nil {
		response.ERROR(c.Writer, http.StatusUnauthorized, err)
		return
	}

	token, err := helpers.JWTGenerator(user)

	if err != nil {
		response.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": "Could not generate token",
			"error": err. Error(),
		})
		return
	}

	response.JSON(c.Writer, http.StatusOK, gin.H{
		"message": "Login User successfully",
		"token":   token,
	})
}

func RegisterUser(c *gin.Context) {
	var request requests.Register

	if err := c.ShouldBindJSON(&request); err != nil {
		response.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
			"error": err.Error(),
		})
		return
	}

	user, err := services.SignUp(&request)
	if err != nil {
		response.ERROR(c.Writer, http.StatusInternalServerError, err)
		return
	}

	response.JSON(c.Writer, http.StatusCreated, gin.H{"message": "Register User Succesfully", "user": user})
}
