package handlers

import (
	"github.com/sakarbaral/personalized-recommendation-engine/services/user-tracking/pkg/services"
)

type UserHandler struct {
	Service services.UserService
}

var userHandler *UserHandler

func NewUserHandler() *UserHandler {
	if userHandler == nil {
		return &UserHandler{
			Service: services.NewUserService(),
		}
	} else {
		return userHandler
	}
}
