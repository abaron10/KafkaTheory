package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reservation/models"
	"reservation/producer"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	producer := producer.NewProducer(producer.ProducerConfig{Hosts: "localhost:9092", OutgoingTopic: "message-log"})

	defer producer.Close()

	// Define a POST endpoint for adding new users
	router.POST("/users", func(c *gin.Context) {
		// Parse the request body into a User struct
		var newUser models.User
		err := c.BindJSON(&newUser)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		// Produce message
		producer.Emit(newUser)

		c.JSON(http.StatusOK, "message produced successfully")
	})

	// Define a GET endpoint for retrieving all users
	router.GET("/users", func(c *gin.Context) {
		// Return the list of users
		c.JSON(http.StatusOK, "ok")
	})

	// Start the server
	fmt.Println("Server running ...")
	router.Run(":8080")
}
