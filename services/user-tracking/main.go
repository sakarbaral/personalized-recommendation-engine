package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sakarbaral/personalized-recommendation-engine/services/user-tracking/pkg/config"
	"github.com/sakarbaral/personalized-recommendation-engine/services/user-tracking/pkg/handlers"
	"github.com/sakarbaral/personalized-recommendation-engine/services/user-tracking/pkg/kafka"
)

func main() {
	// Load configuration from YAML file
	config.LoadConfig("pkg/config/config.yaml")

	// Initialize Kafka client
	kafka.InitKafkaClient(config.AppConfig.Kafka.Brokers, config.AppConfig.Kafka.Topic)

	// Initialize Gin router
	r := gin.Default()

	// Set up routes
	handler := handlers.NewUserHandler()
	r.POST("/track", handler.TrackUserBehavior)

	// Start server
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
