package routes

import (
	"events-rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

// handler = functions (in Go) = controller in MVC (Model-View-Controller)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) // Using () (like getEvenets() executes the function immediately.
	server.GET("/events/:eventId", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent) // PROTECTED
	authenticated.PUT("/events/:eventId", updateEvent)
	authenticated.DELETE("/events/:eventId", deleteEvent)

	server.POST("/signup", signup)
	server.GET("/users", getUsers)
	server.POST("/login", login)
}
