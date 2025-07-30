package handler

import (
	"context"
	"net/http"
	"time"

	"letsplay-microservice/internal/model"
	"letsplay-microservice/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) SignUp(c *gin.Context) {
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

	created, err := h.service.SignUp(ctx, payload)
	if err != nil {
		if err.Error() == "USER_ALREADY_EXISTS" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "USER_ALREADY_EXISTS"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ERROR_CREATING_PLAYER"})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (h *UserHandler) Get(c *gin.Context) {
	response, err := h.service.GetUserDefinitions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "FAILED_TO_GET_USER_DEFINITIONS"})
	}
	c.JSON(http.StatusOK, response)
}
