package routes

import (
	"fmt"
	"go-event-api/models"
	serverError "go-event-api/server-error"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		log.Print("ERROR: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get all the events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId := context.GetInt64("eventId")
	event, err := models.GetEvent(eventId)

	if err != nil {
		log.Print("ERROR: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save all the event"})
		return
	}

	if event == nil {
		context.Error(serverError.NewHttpError(
			fmt.Sprintf("Could not find event with id %v", eventId),
			"",
			http.StatusNotFound,
		))
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event was found", "event": event})
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userId := context.GetInt64("userId")
	event.UserId = userId

	err = event.Save()

	if err != nil {
		log.Print("ERROR: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save all the event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event was created", "event": event})
}

func updateEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	fmt.Print(context.GetInt64("eventId"))
	event, err := models.GetEvent(context.GetInt64("eventId"))

	if err != nil {
		log.Print("ERROR: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the event"})
		return
	}
	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not find the event"})
		return
	}
	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to change this event"})
		return
	}

	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not find event with id provided"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindBodyWithJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	updatedEvent.Id = context.GetInt64("eventId")

	err = updatedEvent.Update()

	if err != nil {
		log.Print("ERROR: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event was updated", "event": updatedEvent})

}

func deleteEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	event, err := models.GetEvent(context.GetInt64("eventId"))

	if err != nil {
		log.Print("ERROR: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the event"})
		return
	}

	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not find event with id provided"})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to DELETE this event"})
		return
	}
	err = event.Delete()

	if err != nil {
		log.Print("ERROR:", err)
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not delete event with id provided"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event was deleted"})
}
