package middleware

import (
	"fmt"
	"go-event-api/models"
	serverError "go-event-api/server-error"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsEventOwner(context *gin.Context) {
	userId := context.GetInt64("userId")
	event := context.MustGet("event").(*models.Event)

	if event.UserId != userId {
		context.Error(serverError.NewHttpError(
			fmt.Sprintf("Not authorized to change the event with id %v", event.Id),
			"",
			http.StatusUnauthorized,
		))
		context.Abort()
	}
	context.Next()
}
