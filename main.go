package main

import (
	"go-event-api/db"
	docs "go-event-api/docs"

	"go-event-api/routes"
	"go-event-api/utils"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	utils.Env.InitEnvironment()
	db.InitDB()
	utils.InitLogger()
	server := gin.Default()

	routes.RegisterRoutes(server)
	docs.SwaggerInfo.BasePath = "/"

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	server.Run(":8080")
}
