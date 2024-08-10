package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sakarbaral/personalized-recommendation-engine/services/recommendation-engine/pkg/config"
	"github.com/sakarbaral/personalized-recommendation-engine/services/recommendation-engine/pkg/kafka"
	"github.com/sakarbaral/personalized-recommendation-engine/services/recommendation-engine/pkg/services"
)

func main() {
	config.LoadConfig("pkg/config/config.yaml")

	r := gin.Default()
	go func() {
		kafka.ConsumeEvents(config.AppConfig.Kafka.Brokers, config.AppConfig.Kafka.Topic, services.NewRecommendationService())
	}()

	log.Println("Starting server on :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}

}
