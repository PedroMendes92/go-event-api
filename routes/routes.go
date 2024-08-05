package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// EVENT ROUTES
	server.GET("/events", getEvents)
	server.GET("events/:eventId", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:eventId", updateEvent)
	server.DELETE("/events/:eventId", deleteEvent)

	//USER ROUTES
	server.POST("/signup", createUser)
	server.POST("/login", login)
}
