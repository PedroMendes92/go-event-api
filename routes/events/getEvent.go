package events

import (
	"go-event-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getEventResponse struct {
	Event models.Event `json:"event"`
}

// @Summary get all events
// @Tags         events
// @Produce      json
// @Description This endpoint will get all available events
// @Success 200 {object} getEventResponse
// @Router /events/{event_id} [get] integer
// @Param event_id path integer true "Event ID"
func GetEvent(context *gin.Context) {
	event := context.MustGet("event").(*models.Event)

	context.JSON(http.StatusOK, getEventResponse{
		Event: *event,
	})
}
