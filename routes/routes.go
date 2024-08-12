package routes

import (
	"go-event-api/middleware"
	serverError "go-event-api/server-error"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case serverError.Http:
				c.AbortWithStatusJSON(e.StatusCode, e)
			default:
				sentry.CaptureException(e)
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Server error. Try again later"})
			}
		}
	}
}

func RegisterRoutes(server *gin.Engine) {
	server.Use(ErrorHandler())

	p := ginprometheus.NewPrometheus("gin")
	p.Use(server)

	server.POST("/signup", createUser)
	server.POST("/login", login)

	server.GET("/events", getEvents)

	//MIDDLEWARE Setup
	checkEventIdParam := middleware.ValidateParam("eventId", "int64")

	routeWithEventId := server.Group("/")
	routeWithEventId.Use(checkEventIdParam, middleware.LoadEventById)

	routeWithEventId.GET("events/:eventId", getEvent)

	authenticatedRoute := server.Group("/")
	authenticatedRoute.Use(middleware.Authenticate)

	authenticatedRoute.POST("/events", createEvent)

	authenticatedRouteWithEventId := authenticatedRoute.Group("/")
	authenticatedRouteWithEventId.Use(checkEventIdParam, middleware.LoadEventById)

	authenticatedRouteWithEventId.PUT("/events/:eventId", middleware.IsEventOwner, updateEvent)
	authenticatedRouteWithEventId.DELETE("/events/:eventId", middleware.IsEventOwner, deleteEvent)
	authenticatedRouteWithEventId.POST("/events/:eventId/register", registerUserToEvent)
	authenticatedRouteWithEventId.DELETE("/events/:eventId/register", removeUserRegistration)
}
