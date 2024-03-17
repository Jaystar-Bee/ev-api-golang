package routes

import (
	"github.com/gin-gonic/gin"
	handler "test.com/event-api/handlers"
	"test.com/event-api/middlewares"
)

func RegisterEventRoutes(server *gin.Engine) {

	server.GET("/events", handler.GetEvents)
	server.GET("/events/:id", handler.GetEvent)

	{
		authenticated := server.Group("/events")

		authenticated.Use(middlewares.Authenticate)
		authenticated.POST("", handler.CreateEvent)
		authenticated.DELETE("/:id", handler.DeleteEvent)
		authenticated.PUT("/:id", handler.UpdateEventByID)
	}
}
