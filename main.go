package main

import (
	"go-event-api/db"
	"go-event-api/routes"
	"go-event-api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.Env.InitEnvironment()
	db.InitDB()
	utils.InitLogger()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
