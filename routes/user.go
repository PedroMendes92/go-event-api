package routes

import (
	"fmt"
	"go-event-api/models"
	serverError "go-event-api/server-error"
	"go-event-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User

	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.Error(serverError.NewHttpError(
			fmt.Sprintf("Could not parse user data. %v", err.Error()),
			"",
			http.StatusBadRequest,
		))
		return
	}

	err = user.Save()

	if err != nil {
		context.Error(serverError.NewHttpError(
			"Could not create user",
			err.Error(),
			http.StatusInternalServerError,
		))
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user was created", "user": user})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.Error(serverError.NewHttpError(
			fmt.Sprintf("Could not parse login data. %v", err.Error()),
			err.Error(),
			http.StatusBadRequest,
		))
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.Error(serverError.NewHttpError(
			"Not authorized",
			err.Error(),
			http.StatusUnauthorized,
		))
		return
	}

	jwtToken, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.Error(serverError.NewHttpError(
			"Could not authenticate user",
			err.Error(),
			http.StatusInternalServerError,
		))
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": jwtToken})

}
