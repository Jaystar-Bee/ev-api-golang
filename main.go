package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"test.com/event-api/models"
)

func main() {

	server := gin.Default()
	server.GET("/events", getEvent)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvent(context *gin.Context) {

	events := models.GetAllEvents()
	context.JSON(http.StatusOK, gin.H{
		"message": "Events fetched successfully",
		"data":    events,
	})
}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not process the request data",
			"data":    nil,
		})
		return
	}

	event.ID = 2
	event.UserId = 10
	event.Save()
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"data":    event,
	})
}
