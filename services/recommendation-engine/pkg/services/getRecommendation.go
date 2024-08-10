package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sakarbaral/personalized-recommendation-engine/services/recommendation-engine/pkg/config"
)

func (s *recommendationService) getRecommendation(userID string, productID string) float64 {
	url := config.AppConfig.External.PythonURL + "/recommend"

	body, _ := json.Marshal(map[string]string{
		"user_id":    userID,
		"product_id": productID,
	})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error calling Python service: %v", err)
		return 0.0
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Error decoding response: %v", err)
		return 0.0
	}

	return result["recommendation"].(float64)
}
