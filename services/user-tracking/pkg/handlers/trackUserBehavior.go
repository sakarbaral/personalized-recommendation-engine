package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sakarbaral/personalized-recommendation-engine/services/user-tracking/pkg/models"
)

func (h *UserHandler) TrackUserBehavior(c *gin.Context) {
	ctx := context.Context(c)
	var event models.UserEvent

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.Service.TrackUserBehavior(ctx, event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to track user behavior"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User behavior tracked"})
}
