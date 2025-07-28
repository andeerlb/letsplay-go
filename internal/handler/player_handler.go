package handler

import (
	"context"
	"net/http"
	"time"

	"letsplay-microservice/internal/model"
	"letsplay-microservice/internal/service"

	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	service *service.PlayerService
}

func NewPlayerHandler(service *service.PlayerService) *PlayerHandler {
	return &PlayerHandler{service: service}
}

func (h *PlayerHandler) SignUp(c *gin.Context) {
	var payload model.SignUp
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "invalid_request",
			"message": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 6*time.Second)
	defer cancel()

	created, err := h.service.CreateNewPlayer(ctx, payload.SignUpAuthServer)
	if err != nil {
		if err.Error() == "USER_ALREADY_EXISTS" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "USER_ALREADY_EXISTS"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}
