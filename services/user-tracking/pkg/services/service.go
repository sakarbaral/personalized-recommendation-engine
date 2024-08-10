package services

import (
	"context"

	"github.com/sakarbaral/personalized-recommendation-engine/services/user-tracking/pkg/models"
)

type UserService interface {
	TrackUserBehavior(ctx context.Context, event models.UserEvent) error
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}
