package routes

import (
	"events-rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {

	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event ID.", "error": err.Error()})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {

	var event models.Event
	// Works a lil'bit like the Scan.
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"message": "No funcionó tu puta madre.",
			"error":   err.Error(),
		})
		return // This exits the function execution.
	}

	// event.ID = 1
	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create evenet."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse event ID.", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the event",
			"error":   err.Error(),
		})
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "You're not authorized to update the event",
		})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindBodyWithJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
			"error":   err.Error(),
		})
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update the event.",
			"error":   err.Error(),
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully.",
	})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not parse the event.",
			"error":   err.Error(),
		})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the event.",
			"error":   err.Error(),
		})
	}
	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete the event.",
			"error":   err.Error(),
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully.",
	})
}
