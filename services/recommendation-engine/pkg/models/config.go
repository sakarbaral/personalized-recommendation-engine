package models

type AppConfig struct {
	Kafka     KafkaConfig `yaml:"kafka"`
	SQLConfig SQLConfig   `yaml:"sql"`
	External  External    `yaml:"external"`
}

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
	Topic   string   `yaml:"topic"`
}

type SQLConfig struct {
	ConnectionURL string `yaml:"connection_url"`
}

type External struct {
	PythonURL string `yaml:"python_url"`
}
