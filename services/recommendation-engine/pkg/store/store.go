package store

import (
	"database/sql"
	"fmt"

	"github.com/sakarbaral/personalized-recommendation-engine/services/recommendation-engine/pkg/db"
)

type recommendationStore struct {
	DB *sql.DB
}

var recStore RecommendationStore

func GetStore() RecommendationStore {
	if recStore == nil {
		return &recommendationStore{
			DB: db.GetDB(),
		}
	}
	return recStore
}

type RecommendationStore interface {
	SaveRecommendation(userID string, productID string, recommendation float64, timestamp uint64) error
}

func (rs *recommendationStore) SaveRecommendation(userID string, productID string, recommendation float64, timestamp uint64) error {
	_, err := rs.DB.Exec("INSERT INTO recommendations (user_id, product_id, score, timestamp) VALUES (?, ?, ?, ?)",
		userID, productID, recommendation, timestamp)
	if err != nil {
		return fmt.Errorf("error inserting recommendation: %w", err)
	}

	return nil
}
