package handler

import (
	"github.com/bahodurnazarov/to-do/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignUpUser(c *gin.Context) {
	var payload *models.SignUpInput
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err := h.Service.SignUpUser(*payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": payload})
}
