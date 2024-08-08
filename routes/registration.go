package routes

import (
	"go-event-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerUserToEvent(context *gin.Context) {
	event, err := models.GetEvent(context.GetInt64("eventId"))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Could not get the event"})
	}

	if event == nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Could not find event with id provided"})
	}

	userId := context.GetInt64("userId")
	err = event.Register(userId)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Could not register user to event"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User was registered for event"})

}

func removeUserRegistration(context *gin.Context) {
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

	userId := context.GetInt64("userId")
	err = event.DeleteRegistration(userId)
	if err != nil {
		log.Print("ERROR: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel the user registration to event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User registration was canceled for event"})
}
