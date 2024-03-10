package handler

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"test.com/event-api/models"
)

// REGISTER USER
func RegisterUser(context *gin.Context) {

	var user models.User

	// CHECK FOR DATA PROCESSING
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not process the data",
			"error":   err.Error(),
		})
		return
	}

	// CHECK IF IT IS VALID EMAIL
	_, err = mail.ParseAddress(user.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user mail",
			"error":   err.Error(),
		})
		return
	}

	// CHECK IF USER ALREADY EXIST WITH THE EMAIL
	userFound, _ := models.GetUserByEmail(user.Email)
	if userFound != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "User already exist",
			"error":   nil,
		})
		return
	}

	// SVE THE USER
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save the data",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})

}

// LOGIN USER

func LoginUser(context *gin.Context) {
	var login models.Login

	err := context.ShouldBindJSON(&login)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not process the data",
			"error":   err.Error(),
		})
		return
	}

	err = login.ValidateUserCredentials()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
			"error":   "Cound not authenticate user",
		})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{
		"message": "User logged in successfully",
	})

	// Login user

}
