package registration

import (
	"go-event-api/models"
	serverError "go-event-api/server-error"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary register a user
// @Tags         events/registration
// @Accept       json
// @Produce      json
// @Description This endpoint allows logged in user to register to an event
// @Success 204
// @Router /events/{event_id}/register [post] integer
// @Param event_id path integer true "Event ID"
// @Security Bearer
func RegisterUserToEvent(context *gin.Context) {
	event := context.MustGet("event").(*models.Event)
	userId := context.GetInt64("userId")

	err := event.Register(userId)

	if err != nil {
		context.Error(serverError.NewHttpError(
			"Could not register user to event",
			"",
			http.StatusInternalServerError,
		))
		return
	}

	context.Status(http.StatusNoContent)

}
