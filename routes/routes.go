package routes

import "github.com/gin-gonic/gin"

// handler = functions (in Go) = controller in MVC (Model-View-Controller)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) // Using () (like getEvenets() executes the function immediately.
	server.POST("/events", createEvent)
	server.GET("/events/:eventId", getEvent)
	server.PUT("/events/:id", updateEvent)

}
