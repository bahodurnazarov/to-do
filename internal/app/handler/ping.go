package handler

import (
	_ "github.com/bahodurnazarov/to-do/cmd/app/docs"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingExample godoc
// @Summary Ping
// @Schemes
// @Description do ping
// @Tags Routes
// @Accept json
// @Produce json
// @Success 200
// @Router /ping [get]
// Init foutes function
func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}
