package events

import (
	"fmt"
	"go-event-api/models"
	serverError "go-event-api/server-error"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Delete event
// @Tags         events
// @Produce      json
// @Description This endpoint will delete a specific event
// @Success 200
// @Router /events/{event_id} [delete] integer
// @Param event_id path integer true "Event ID"
// @Security Bearer
func DeleteEvent(context *gin.Context) {
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

	context.Status(http.StatusOK)
}
