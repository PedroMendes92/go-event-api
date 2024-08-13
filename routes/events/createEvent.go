package events

import (
	"go-event-api/models"
	serverError "go-event-api/server-error"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type createEventInput struct {
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"date"`
}

type createEventResponse struct {
	Event models.Event `json:"event"`
}

// @Summary Create an event
// @Tags         events
// @Accept       json
// @Produce      json
// @Description This endpoint will create an event and associate it to the user
// @Param event body createEventInput true "new event data"
// @Success 200 {object} createEventResponse
// @Router /events [post]
// @Security Bearer
func CreateEvent(context *gin.Context) {
	var eventInput createEventInput

	err := context.ShouldBindBodyWithJSON(&eventInput)

	if err != nil {
		context.Error(serverError.NewHttpError(
			"Could not parse data into an event",
			err.Error(),
			http.StatusBadRequest,
		))
		return
	}

	event := models.Event{
		Name:        eventInput.Name,
		Description: eventInput.Description,
		Location:    eventInput.Location,
		DateTime:    eventInput.DateTime,
		UserId:      context.GetInt64("userId"),
	}

	err = event.Save()

	if err != nil {
		context.Error(serverError.NewHttpError(
			"Could not save the event",
			err.Error(),
			http.StatusInternalServerError,
		))
		return
	}
	context.JSON(http.StatusCreated, createEventResponse{
		Event: event,
	})
}
