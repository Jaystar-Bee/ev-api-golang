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
	eventToUpdate.UserId = event.UserId // Set the user ID to the original user ID of the event to update. This ensures that the user ID of the updated event is the same as the original user ID. Without this, the updated event would have a different user ID than the original event. This would result in a security vulnerability if the
	eventToUpdate.CreatedAt = event.CreatedAt
	err = eventToUpdate.UpdateEvent()

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

// REGISTER USER FOR AN EVENT
func RegisterForEvent(context *gin.Context) {
	parsedEventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id",
		})
		return
	}

	event, err := models.GetEvent(parsedEventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Could not find event",
		})
		return
	}
	userId := context.GetInt64("userId")
	_, err = event.GetRegistrationByEvent(userId)
	if err == nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Registration already exists",
		})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error registering for the event!",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Registration for " + event.Name + " was successful!",
	})

}

func CancelRegistrationForEvent(context *gin.Context) {
	parsedEventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id",
		})
		return
	}
	event, err := models.GetEvent(parsedEventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Could not find event",
		})
		return
	}

	userId := context.GetInt64("userId")
	_, err = event.GetRegistrationByEvent(userId)
	if err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"message": "You have not registered for this event",
		})
		return
	}

	err = event.Cancel(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error cancelling registration for the event!",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Registration for " + event.Name + " was cancelled!",
	})
}
