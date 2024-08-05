package routes

import (
	"go-event-api/models"
	"log"
	"net/http"
	"strconv"

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
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id"})
		return
	}

	event, err := models.GetEvent(eventId)

	if err != nil {
		log.Print("ERROR: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save all the event"})
		return
	}

	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not find event with id provided"})
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

	err = event.Save()
	if err != nil {
		if err != nil {
			log.Print("ERROR: ", err)
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save all the event"})
			return
		}
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event was created", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id"})
		return
	}

	event, err := models.GetEvent(eventId)

	if err != nil {
		log.Print("ERROR: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the event"})
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

	updatedEvent.Id = eventId

	err = updatedEvent.Update()

	if err != nil {
		log.Print("ERROR: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event was updated", "event": updatedEvent})

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id"})
		return
	}

	event, err := models.GetEvent(eventId)

	if err != nil {
		log.Print("ERROR: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the event"})
		return
	}

	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not find event with id provided"})
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
