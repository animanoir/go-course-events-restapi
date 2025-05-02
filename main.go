package main

import (
	"events-rest-api/db"
	"events-rest-api/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	fmt.Print("Server has started!")
	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
