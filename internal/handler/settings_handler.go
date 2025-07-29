package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SettingsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
