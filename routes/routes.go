package routes

import (
	"go-event-api/middleware"
	"go-event-api/routes/events"
	"go-event-api/routes/registration"
	"go-event-api/routes/user"
	serverError "go-event-api/server-error"
	"go-event-api/utils"
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
	p.SetMetricsPathWithAuth(server, gin.Accounts{
		"admin": utils.Env.MetricsPassword,
	})

	server.POST("/signup", user.CreateUser)
	server.POST("/login", user.Login)

	server.GET("/events", events.GetEvents)

	//MIDDLEWARE Setup
	checkEventIdParam := middleware.ValidateParam("eventId", "int64")

	routeWithEventId := server.Group("/")
	routeWithEventId.Use(checkEventIdParam, middleware.LoadEventById)

	routeWithEventId.GET("events/:eventId", events.GetEvent)

	authenticatedRoute := server.Group("/")
	authenticatedRoute.Use(middleware.Authenticate)

	authenticatedRoute.POST("/events", events.CreateEvent)

	authenticatedRouteWithEventId := authenticatedRoute.Group("/")
	authenticatedRouteWithEventId.Use(checkEventIdParam, middleware.LoadEventById)

	authenticatedRouteWithEventId.PUT("/events/:eventId", middleware.IsEventOwner, events.UpdateEvent)
	authenticatedRouteWithEventId.DELETE("/events/:eventId", middleware.IsEventOwner, events.DeleteEvent)
	authenticatedRouteWithEventId.POST("/events/:eventId/register", registration.RegisterUserToEvent)
	authenticatedRouteWithEventId.DELETE("/events/:eventId/register", registration.RemoveUserRegistration)
}
