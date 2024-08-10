package services

import "github.com/sakarbaral/personalized-recommendation-engine/services/recommendation-engine/pkg/store"

type RecommendationService interface {
	ProcessEvent(event []byte)
}

type recommendationService struct {
	Store store.RecommendationStore
}

func NewRecommendationService() RecommendationService {
	return &recommendationService{
		Store: store.GetStore(),
	}
}
