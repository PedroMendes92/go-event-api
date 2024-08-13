package user

import (
	"fmt"
	"go-event-api/models"
	serverError "go-event-api/server-error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createUserResponse struct {
	User models.User `json:"user"`
}

// @Summary Create a new user
// @Tags         user
// @Produce      json
// @Description
// @Success 200 {object} createUserResponse
// @Router /signup [post]
// @Param data body createUserInput true "user data"
func CreateUser(context *gin.Context) {
	var userInput createUserInput

	err := context.ShouldBindBodyWithJSON(&userInput)

	if err != nil {
		context.Error(serverError.NewHttpError(
			fmt.Sprintf("Could not parse user data. %v", err.Error()),
			"",
			http.StatusBadRequest,
		))
		return
	}

	user := models.User{
		Email:    userInput.Email,
		Password: userInput.Password,
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

	context.JSON(http.StatusCreated, createUserResponse{User: user})
}
