package routes

import (
	"go-event-api/models"
	serverError "go-event-api/server-error"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerUserToEvent(context *gin.Context) {
	event := context.MustGet("event").(*models.Event)
	userId := context.GetInt64("userId")

	err := event.Register(userId)

	if err != nil {
		context.Error(serverError.NewHttpError(
			"Could not register user to event",
			"",
			http.StatusInternalServerError,
		))
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User was registered for event"})

}

func removeUserRegistration(context *gin.Context) {
	event := context.MustGet("event").(*models.Event)
	userId := context.GetInt64("userId")

	err := event.DeleteRegistration(userId)

	if err != nil {
		context.Error(serverError.NewHttpError(
			"Could not cancel the user registration to event",
			"",
			http.StatusInternalServerError,
		))
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User registration was canceled for event"})
}
