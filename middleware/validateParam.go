package middleware

import (
	"errors"
	serverError "go-event-api/server-error"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ValidateParam(paramName string, paramType string) gin.HandlerFunc {
	return func(context *gin.Context) {
		value := context.Param(paramName)

		if value == "" {
			context.AbortWithError(500, errors.New("Could not find param "+paramName+" in the URL"))
			return
		}

		switch paramType {
		case "int64":
			param, err := strconv.ParseInt(value, 10, 64)

			if err != nil {
				context.Error(serverError.NewHttpError("Could not parse the param. It must be a number", "", http.StatusBadRequest))
				context.Abort()
				return
			}
			context.Set(paramName, param)
			context.Next()
			return
		default:
			context.Next()
		}

	}
}
