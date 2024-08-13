package middleware

import (
	"go-event-api/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	authString := context.Request.Header.Get("Authorization")
	values := strings.Split(authString, " ")

	if len(values) != 2 {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid Token string"})
		return
	}

	token := values[1]

	if token == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "request is missing authorization header"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized access"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
