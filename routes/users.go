package routes

import (
	"first-gin/models"
	"first-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Validation wrong."})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message" : "User created!", "user" : user})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Validation wrong."})
		return
	}

	err = user.ValidateUser()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message" : err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message" : "Internal server error."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message" : "Login success.", "token" : token})
}