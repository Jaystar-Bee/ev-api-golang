package routes

import (
	"github.com/gin-gonic/gin"
	handler "test.com/event-api/handlers"
)

func RegisterUserRoutes(server *gin.Engine) {
	user := server.Group("/user")
	{
		user.POST("/signup", handler.RegisterUser)
		user.POST("/login", handler.LoginUser)
	}
}
