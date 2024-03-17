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
			"error":   err.Error(),
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
			"error":   err.Error(),
		})
		return
	}

	event, err := models.GetEvent(parsedId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Could not fetch event",
			"error":   err.Error(),
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
			"error":   err.Error(),
		})
		return
	}

	eventToDelete, err := models.GetEvent(parsedId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Could not fetch event to edit",
			"error":   err.Error(),
		})
		return
	}

	userId := context.GetInt64("userId")
	if eventToDelete.UserId != userId {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "You are not authorized to delete this event",
			"data":    nil,
		})
		return
	}

	err = models.DeleteEvent(parsedId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Could not delete event",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
		"data":    eventToDelete,
	})
}

// CREATE AN EVENT
func CreateEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not process the request data",
			"error":   err.Error(),
		})
		return
	}
	userId := context.GetInt64("userId")
	event.UserId = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event",
			"data":    nil,
			"error":   err.Error(),
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
			"error":   err.Error(),
		})
		return
	}

	event, err := models.GetEvent(parsedId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Could not find event to edit",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}

	userId := context.GetInt64("userId")
	if event.UserId != userId {
		context.JSON(http.StatusForbidden, gin.H{
			"message": "You are not authorized to edit this event",
			"data":    nil,
		})
		return
	}

	var eventToUpdate models.Event
	err = context.ShouldBindJSON(&eventToUpdate)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not process the request data",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}
	eventToUpdate.ID = parsedId
	err = eventToUpdate.UpdateEvent(event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event",
			"data":    nil,
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"data":    eventToUpdate,
	})
}
