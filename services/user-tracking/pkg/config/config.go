package config

import (
	"log"

	"os"

	"github.com/sakarbaral/personalized-recommendation-engine/services/user-tracking/pkg/models"
	"gopkg.in/yaml.v2"
)

var AppConfig models.AppConfig

func LoadConfig(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("error opening config file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("error decoding config file: %v", err)
	}
}
