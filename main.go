package main

import (
	"go-event-api/db"
	"go-event-api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		log.Print("ERROR: ", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get all the events"})
		return
	}
	context.JSON(http.StatusOK, events)
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
