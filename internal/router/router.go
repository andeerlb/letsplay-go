package router

import (
	"letsplay-microservice/internal/bootstrap"
	"letsplay-microservice/internal/config"
	"letsplay-microservice/internal/handler"
	"letsplay-microservice/internal/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewRouter(c *bootstrap.Container, logg *zap.Logger, cfg *config.Config) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())

	public := router.Group("/api/public")
	{
		public.GET("/health", handler.HealthCheck)
		public.POST("/newplayer", c.PlayerHandler.SignUp)
	}

	protected := router.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware(&cfg.JwtSecret))
	{
		// protected.GET("/user", ...)
	}

	return router
}
