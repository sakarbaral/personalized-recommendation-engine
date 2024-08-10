package services

import (
	"encoding/json"
	"log"
)

func (s *recommendationService) ProcessEvent(event []byte) {
	var eventData map[string]interface{}
	if err := json.Unmarshal(event, &eventData); err != nil {
		log.Printf("Error unmarshalling event data: %v", err)
		return
	}

	userID := eventData["user_id"].(string)
	timestamp := (uint64)(123)
	productID := eventData["product_id"].(string)

	// Call the Python API to get the recommendation
	recommendation := s.getRecommendation(userID, productID)
	log.Printf("Recommendation for user %s: %v and timestamp, : %v", userID, recommendation, timestamp)

	err := s.Store.SaveRecommendation(userID, productID, recommendation, timestamp)
	if err != nil {
		log.Fatal("Error in saving recommendation ", err)
	}
}
