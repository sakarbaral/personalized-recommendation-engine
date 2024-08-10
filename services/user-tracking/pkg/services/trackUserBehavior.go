package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sakarbaral/personalized-recommendation-engine/services/user-tracking/pkg/kafka"
	"github.com/sakarbaral/personalized-recommendation-engine/services/user-tracking/pkg/models"
)

func (s *userService) TrackUserBehavior(ctx context.Context, event models.UserEvent) error {
	// Convert event to JSON

	eventJSON, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Publish to Kafka
	if err := kafka.PublishEvent(ctx, eventJSON); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
