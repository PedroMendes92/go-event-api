package user

import (
	"fmt"
	"go-event-api/models"
	serverError "go-event-api/server-error"
	"go-event-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	Token string `json:"token"`
}

// @Summary Login user
// @Tags         user
// @Produce      json
// @Description
// @Success 200 {object} loginResponse
// @Router /login [post]
// @Param data body loginInput true "login data"
func Login(context *gin.Context) {
	var inputData loginInput

	err := context.ShouldBindBodyWithJSON(&inputData)

	if err != nil {
		context.Error(serverError.NewHttpError(
			fmt.Sprintf("Could not parse login data. %v", err.Error()),
			err.Error(),
			http.StatusBadRequest,
		))
		return
	}

	user := models.User{
		Email:    inputData.Email,
		Password: inputData.Password,
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

	context.JSON(http.StatusOK, loginResponse{Token: jwtToken})

}
