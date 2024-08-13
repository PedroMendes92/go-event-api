package events

import (
	"errors"
	"fmt"
	"go-event-api/models"
	serverError "go-event-api/server-error"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type updateEventInput struct {
	Name        *string
	Description *string
	Location    *string
	Date        *string
}

func (input updateEventInput) mergeWith(event *models.Event) error {

	mergeError := errors.New("nothing has changed")

	if input.Name != nil {
		event.Name = *input.Name
		mergeError = nil
	}
	if input.Description != nil {
		event.Description = *input.Description
		mergeError = nil
	}
	if input.Location != nil {
		event.Location = *input.Location
		mergeError = nil
	}
	if input.Date != nil {
		dateTime, err := time.Parse(time.DateTime, *input.Date)
		if err != nil {
			mergeError = err
		} else {
			event.DateTime = dateTime
			mergeError = nil
		}
	}

	return mergeError
}

type updateEventResponse struct {
	Event models.Event `json:"event"`
}

// @Summary Update an event
// @Tags         events
// @Accept       json
// @Produce      json
// @Description This endpoint will update an event if associated with the user
// @Param event body updateEventInput true "event data to update"
// @Success 200 {object} updateEventResponse
// @Router /events/{event_id} [put] integer
// @Param event_id path integer true "Event ID"
// @Security Bearer
func UpdateEvent(context *gin.Context) {
	var inputData updateEventInput

	err := context.ShouldBindBodyWithJSON(&inputData)

	if err != nil {
		context.Error(serverError.NewHttpError(
			fmt.Sprintf("Could not parse data into an event. %v", err.Error()),
			"",
			http.StatusBadRequest,
		))
		return
	}
	event := context.MustGet("event").(*models.Event)

	err = inputData.mergeWith(event)

	if err != nil {
		context.Error(serverError.NewHttpError(
			err.Error(),
			"",
			http.StatusBadRequest,
		))
		return
	}

	err = event.Update()

	if err != nil {
		context.Error(serverError.NewHttpError(
			fmt.Sprintf("Could not update the event with id %v", event.Id),
			err.Error(),
			http.StatusInternalServerError,
		))
		return
	}

	context.JSON(http.StatusCreated, updateEventResponse{
		Event: *event,
	})
}
