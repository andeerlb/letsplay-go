package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"letsplay-microservice/internal/model"
	"letsplay-microservice/internal/service"
	"net/http"
	"time"
)

type SettingsHandlers struct {
	service *service.SettingsService
}

func NewSettingsHandler(service *service.SettingsService) *SettingsHandlers {
	return &SettingsHandlers{service: service}
}

func (h *SettingsHandlers) Get(c *gin.Context) {
	response, err := h.service.Get(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "FAILED_TO_GET_SETTINGS"})
		return
	}

	if response == nil {
		c.JSON(http.StatusNoContent, nil)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func (h *SettingsHandlers) Save(c *gin.Context) {
	var payload model.Settings
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "invalid_request",
			"message": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 6*time.Second)
	defer cancel()

	newSettings, err := h.service.Save(ctx, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ERROR_CREATING_SETTINGS"})
		return
	}

	c.JSON(http.StatusCreated, newSettings)
}
