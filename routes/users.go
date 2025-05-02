package routes

import (
	"events-rest-api/models"
	"events-rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save user",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully.",
	})
}

func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not return all users",
			"error":   err.Error(),
		})
	}
	context.JSON(http.StatusOK, users)
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not do the thing",
			"error":   err.Error(),
		})
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not authorize the user.",
			"error":   err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "There was an error generating the token.",
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
	// I used to have a panic because I tried to send an error here, but this should not be.

}
