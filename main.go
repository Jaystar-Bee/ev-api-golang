package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"test.com/event-api/db"
	"test.com/event-api/models"
)

func main() {
	db.InitDatabase()

	server := gin.Default()
	server.GET("/events", getEvent)
	server.POST("/events", createEvent)
	server.Run(":8000")
}

func getEvent(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events",
			"data":    nil,
		})
		return
	}

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

	event.UserId = 10
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event",
			"data":    nil,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"data":    event,
	})
}
