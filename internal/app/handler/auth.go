package handler

import (
	_ "github.com/bahodurnazarov/to-do/cmd/app/docs"
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

func (h *Handler) SignInUser(c *gin.Context) {
	var payload *models.SignInInput
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail1", "message": err.Error()})
		return
	}

	token, err := h.Service.SignInUser(*payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail2", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "token": token})
}
