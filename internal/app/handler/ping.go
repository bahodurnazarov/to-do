package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}
