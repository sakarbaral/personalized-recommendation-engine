package models

type AppConfig struct {
	Kafka KafkaConfig `yaml:"kafka"`
}
