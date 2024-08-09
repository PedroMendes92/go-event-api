package middleware

import (
	"fmt"
	"go-event-api/models"
	serverError "go-event-api/server-error"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadEventById(context *gin.Context) {
	eventId := context.GetInt64("eventId")
	event, err := models.GetEvent(eventId)

	if err != nil {
		context.Error(serverError.NewHttpError(
			fmt.Sprintf("Could not get the event with id %v", eventId),
			err.Error(),
			http.StatusInternalServerError,
		))
		context.Abort()
		return
	}

	if event == nil {
		context.Error(serverError.NewHttpError(
			fmt.Sprintf("Could not find the event with id %v", eventId),
			"",
			http.StatusNotFound,
		))
		context.Abort()
		return
	}

	context.Set("event", event)
	context.Next()
}
