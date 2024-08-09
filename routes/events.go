package routes

import (
	"fmt"
	"go-event-api/models"
	serverError "go-event-api/server-error"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.Error(serverError.NewHttpError(
			"Could not get all the events",
			err.Error(),
			http.StatusInternalServerError,
		))
		return
	}

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	event := context.MustGet("event").(*models.Event)

	context.JSON(http.StatusOK, gin.H{"message": "event was found", "event": event})
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.Error(serverError.NewHttpError(
			"Could not parse data into an event",
			err.Error(),
			http.StatusBadRequest,
		))
		return
	}

	userId := context.GetInt64("userId")
	event.UserId = userId

	err = event.Save()

	if err != nil {
		context.Error(serverError.NewHttpError(
			"Could not save the event",
			err.Error(),
			http.StatusInternalServerError,
		))
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event was created", "event": event})
}

func updateEvent(context *gin.Context) {
	event := context.MustGet("event").(*models.Event)

	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.Error(serverError.NewHttpError(
			fmt.Sprintf("Could not parse data into an event. %v", err.Error()),
			"",
			http.StatusBadRequest,
		))
		return
	}

	event.Id = context.GetInt64("eventId")

	err = event.Update()

	if err != nil {
		context.Error(serverError.NewHttpError(
			fmt.Sprintf("Could not update the event with id %v", event.Id),
			err.Error(),
			http.StatusInternalServerError,
		))
		return
	}

	context.JSON(http.StatusCreated, event)
}

func deleteEvent(context *gin.Context) {
	event := context.MustGet("event").(*models.Event)

	err := event.Delete()

	if err != nil {
		context.Error(serverError.NewHttpError(
			fmt.Sprintf("Could not delete event with id %v", event.Id),
			err.Error(),
			http.StatusInternalServerError,
		))
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event was deleted"})
}
