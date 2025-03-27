package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handler = functions (in Go) = controller in MVC (Model-View-Controller)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents) // Using () (like getEvenets() executes the function immediately.
	server.Run(":8080")              // localhost:8080
}

func getEvents(context *gin.Context) {
	log.Println("The /events handler has been executed.", context)
	context.JSON(http.StatusOK, gin.H{"message": "Hi, Human!"})
}
