package routes

import (
	"github.com/gin-gonic/gin"
	handler "test.com/event-api/handlers/event"
)

func RegisterEventRoutes(server *gin.Engine) {

	server.GET("/events", handler.GetEvents)
	server.POST("/events", handler.CreateEvent)
	server.GET("/events/:id", handler.GetEvent)
	server.DELETE("/events/:id", handler.DeleteEvent)
}
