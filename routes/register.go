package routes

import (
	"events-rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	// Takes a param (in this case :eventId) and parses it.
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse event id",
			"error":   err.Error(),
		})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch event",
			"error":   err.Error(),
		})
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not register the event",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "The event has been registered!",
	})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "coud not parse event id",
			"error":   err.Error(),
		})
		return
	}

	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not cancel the registration",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "The registration has been canceled!",
	})
}
