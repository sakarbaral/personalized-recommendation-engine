package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

var writer *kafka.Writer

func InitKafkaClient(brokers []string, topic string) {
	var logger = logrus.New()

	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:     brokers,
		Topic:       topic,
		Balancer:    &kafka.LeastBytes{},
		Logger:      kafka.LoggerFunc(logger.Infof),
		ErrorLogger: kafka.LoggerFunc(logger.Errorf),
	})
	fmt.Println(writer)
}

func PublishEvent(ctx context.Context, message []byte) error {
	fmt.Println("Message published to kafka", message)
	return writer.WriteMessages(ctx, kafka.Message{Value: message})
}
