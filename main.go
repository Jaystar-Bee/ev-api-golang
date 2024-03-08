package main

import (
	"github.com/gin-gonic/gin"
	"test.com/event-api/db"
	"test.com/event-api/routes"
)

func main() {
	db.InitDatabase()
	server := gin.Default()
	routes.RegisterEventRoutes(server)
	server.Run(":8000")
}
