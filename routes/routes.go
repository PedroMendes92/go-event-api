package routes

import (
	"go-event-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticatedRoute := server.Group("/")
	authenticatedRoute.Use(middleware.Authenticate)

	// EVENT ROUTES
	server.GET("/events", getEvents)
	server.GET("events/:eventId", getEvent)

	authenticatedRoute.POST("/events", createEvent)
	authenticatedRoute.PUT("/events/:eventId", updateEvent)
	authenticatedRoute.DELETE("/events/:eventId", deleteEvent)
	authenticatedRoute.POST("/events/:eventId/register", registerUserToEvent)
	authenticatedRoute.DELETE("/events/:eventId/register", removeUserRegistration)
	//USER ROUTES
	server.POST("/signup", createUser)
	server.POST("/login", login)
}
