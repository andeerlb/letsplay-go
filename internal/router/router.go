package router

import (
	"letsplay-microservice/internal/bootstrap"
	"letsplay-microservice/internal/handler"
	"letsplay-microservice/internal/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewRouter(c *bootstrap.Container, logg *zap.Logger) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())

	router.GET("/health", handler.HealthCheck)
	router.POST("/newplayer", c.PlayerHandler.SignUp)

	protected := router.Group("/s")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		// protected.GET("/user", ...)
	}

	return router
}
