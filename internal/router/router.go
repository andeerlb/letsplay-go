package router

import (
	"letsplay-microservice/internal/bootstrap"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewRouter(c *bootstrap.Container, logg *zap.Logger) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	router.POST("/newplayer", c.PlayerHandler.SignUp)

	return router
}
