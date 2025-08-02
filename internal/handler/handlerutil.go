package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"letsplay-microservice/internal/locale"
	"letsplay-microservice/internal/middleware"
	"time"
)

func WithTimeoutFromGin(c *gin.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	base := c.Request.Context()
	ctx, cancel := context.WithTimeout(base, timeout)

	if userID, ok := c.Get(middleware.UserIDKey); ok {
		ctx = context.WithValue(ctx, middleware.UserIDKey, userID)
	}

	if localizer, ok := c.Get(locale.LocalizerKey); ok {
		ctx = context.WithValue(ctx, locale.LocalizerKey, localizer)
	}

	return ctx, cancel
}
