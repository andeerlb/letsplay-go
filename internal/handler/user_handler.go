package handler

import (
	"letsplay-microservice/internal/locale"
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

	ctx, cancel := WithTimeoutFromGin(c, time.Second*6)
	defer cancel()

	created, err := h.service.SignUp(ctx, payload)
	if err != nil {
		if err.Error() == "USER_ALREADY_EXISTS" {
			msg := locale.Msg(ctx, "user_handler.user-already-exists")
			c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}
		msg := locale.Msg(ctx, "user_service.failed-to-delete-user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (h *UserHandler) Get(c *gin.Context) {
	ctx, cancel := WithTimeoutFromGin(c, time.Second*6)
	defer cancel()

	response, err := h.service.GetUserDefinitions(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "FAILED_TO_GET_USER_DEFINITIONS"})
		return
	}

	if response == nil {
		c.JSON(http.StatusNoContent, nil)
	} else {
		c.JSON(http.StatusOK, response)
	}
}
