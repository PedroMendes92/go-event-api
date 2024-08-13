package registration

import (
	"go-event-api/models"
	serverError "go-event-api/server-error"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary remove a registration
// @Tags         events/registration
// @Accept       json
// @Produce      json
// @Description This endpoint allows teh logged in user to remove a registration to an event
// @Success 204
// @Router /events/{event_id}/register [delete] integer
// @Param event_id path integer true "Event ID"
// @Security Bearer
func RemoveUserRegistration(context *gin.Context) {
	event := context.MustGet("event").(*models.Event)
	userId := context.GetInt64("userId")

	err := event.DeleteRegistration(userId)

	if err != nil {
		context.Error(serverError.NewHttpError(
			"Could not cancel the user registration to event",
			"",
			http.StatusInternalServerError,
		))
		return
	}

	context.Status(http.StatusNoContent)
}
