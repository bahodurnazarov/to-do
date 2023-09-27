package handler

import (
	"github.com/bahodurnazarov/to-do/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) AddTask(c *gin.Context) {
	var task models.Tasks
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.AddTask(task); err != nil {
		log.Println(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
