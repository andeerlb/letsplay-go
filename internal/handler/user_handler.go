package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
