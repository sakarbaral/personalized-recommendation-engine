package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/sakarbaral/personalized-recommendation-engine/services/recommendation-engine/pkg/services"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func ConsumeEvents(brokers []string, topic string, service services.RecommendationService) {
	var logger = logrus.New()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		Logger:      kafka.LoggerFunc(logger.Infof),
		ErrorLogger: kafka.LoggerFunc(logger.Errorf),
	})

	for {
		msg, err := reader.ReadMessage(context.Background())
		fmt.Println("The message from kafka is", msg)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		// Process the message
		service.ProcessEvent(msg.Value)
	}
}
