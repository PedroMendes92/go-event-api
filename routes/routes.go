package routes

import (
	"go-event-api/middleware"
	serverError "go-event-api/server-error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorMeta struct {
	status int
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case serverError.Http:
				c.AbortWithStatusJSON(e.StatusCode, e)
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Service Unavailable"})
			}
		}
	}
}

func RegisterRoutes(server *gin.Engine) {
	server.Use(ErrorHandler())

	authenticatedRoute := server.Group("/")
	authenticatedRoute.Use(middleware.Authenticate)

	checkEventIdParam := middleware.ValidateParam("eventId", "int64")

	// EVENT ROUTES
	server.GET("/events", getEvents)
	server.GET("events/:eventId", checkEventIdParam, getEvent)

	authenticatedRoute.POST("/events", createEvent)
	authenticatedRoute.PUT("/events/:eventId", checkEventIdParam, updateEvent)
	authenticatedRoute.DELETE("/events/:eventId", checkEventIdParam, deleteEvent)
	authenticatedRoute.POST("/events/:eventId/register", checkEventIdParam, registerUserToEvent)
	authenticatedRoute.DELETE("/events/:eventId/register", checkEventIdParam, removeUserRegistration)
	//USER ROUTES
	server.POST("/signup", createUser)
	server.POST("/login", login)
}
