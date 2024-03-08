package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"test.com/event-api/models"
)

// GET ALL EVENTS
func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Could not fetch events",
			"data":    nil,
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Events fetched successfully",
		"data":    events,
	})
}

// GET EVENT BY ID
func GetEvent(context *gin.Context) {
	id := context.Param("id")

	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id",
			"data":    nil,
			"error":   err,
		})
		return
	}

	event, err := models.GetEvent(parsedId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Could not fetch event",
			"data":    nil,
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event fetched successfully",
		"data":    event,
	})
}

// DELETE EVENT BY ID
func DeleteEvent(context *gin.Context) {
	id := context.Param("id")

	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id",
			"data":    nil,
			"error":   err,
		})
		return
	}

	_, err = models.GetEvent(parsedId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Could not fetch event to edit",
			"data":    nil,
			"error":   err,
		})
		return
	}

	err = models.DeleteEvent(parsedId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Could not delete event",
			"data":    nil,
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
		"data":    nil,
	})
}

// CREATE AN EVENT
func CreateEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not process the request data",
			"data":    nil,
			"error":   err,
		})
		return
	}

	event.UserId = 10
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event",
			"data":    nil,
			"error":   err,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"data":    event,
	})
}

// UPDATE AN EVENT BY ID

func UpdateEventByID(context *gin.Context) {
	id := context.Param("id")
	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id",
			"data":    nil,
			"error":   err,
		})
		return
	}

	_, err = models.GetEvent(parsedId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Could not fetch event to edit",
			"data":    nil,
			"error":   err,
		})
		return
	}

	var eventToUpdate models.Event
	err = context.ShouldBindJSON(&eventToUpdate)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not process the request data",
			"data":    nil,
			"error":   err,
		})
		return
	}
	eventToUpdate.ID = parsedId
	err = eventToUpdate.UpdateEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event",
			"data":    nil,
			"error":   err,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"data":    eventToUpdate,
	})
}