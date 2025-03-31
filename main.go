package main

import (
	"log"
	"net/http"

	"events-rest-api/models"

	"github.com/gin-gonic/gin"
)

// handler = functions (in Go) = controller in MVC (Model-View-Controller)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents) // Using () (like getEvenets() executes the function immediately.
	server.POST("/events", createEvent)
	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	log.Println("The /events handler has been executed.", context)
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	// Works a lil'bit like the Scan.
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "No funcionó tu puta madre."})
		return // This exits the function execution.
	}

	event.ID = 1
	event.UserID = 1
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
