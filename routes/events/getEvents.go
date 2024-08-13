package events

import (
	"go-event-api/models"
	serverError "go-event-api/server-error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getEventsResponse struct {
	Events []models.Event `json:"events"`
}

// @Summary get all events
// @Tags         events
// @Accept       json
// @Produce      json
// @Description This endpoint will get all available events
// @Success 200 {object} getEventsResponse
// @Router /events [get]
func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.Error(serverError.NewHttpError(
			"Could not get all the events",
			err.Error(),
			http.StatusInternalServerError,
		))
		return
	}

	context.JSON(http.StatusOK, getEventsResponse{
		Events: events,
	})
}
