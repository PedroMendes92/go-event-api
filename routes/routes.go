package routes

import (
	"go-event-api/middleware"
	serverError "go-event-api/server-error"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	//MIDDLEWARE Setup
	checkEventIdParam := middleware.ValidateParam("eventId", "int64")

	authenticatedRoute := server.Group("/")
	authenticatedRoute.Use(middleware.Authenticate)

	authenticatedRouteWithEventId := server.Group("/")
	authenticatedRouteWithEventId.Use(middleware.Authenticate, checkEventIdParam, middleware.LoadEventById)

	routeWithEventId := server.Group("/")
	routeWithEventId.Use(checkEventIdParam, middleware.LoadEventById)

	server.POST("/signup", createUser)
	server.POST("/login", login)

	server.GET("/events", getEvents)

	routeWithEventId.GET("events/:eventId", getEvent)

	authenticatedRouteWithEventId.POST("/events", createEvent)
	authenticatedRouteWithEventId.PUT("/events/:eventId", middleware.IsEventOwner, updateEvent)
	authenticatedRouteWithEventId.DELETE("/events/:eventId", middleware.IsEventOwner, deleteEvent)
	authenticatedRouteWithEventId.POST("/events/:eventId/register", registerUserToEvent)
	authenticatedRouteWithEventId.DELETE("/events/:eventId/register", removeUserRegistration)
}
