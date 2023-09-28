package handler

import (
	"github.com/bahodurnazarov/to-do/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AddTask(c *gin.Context) {
	var task models.Tasks
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.AddTask(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *Handler) AllTasks(c *gin.Context) {
	tasks, err := h.Service.AllTasks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) RemoveTask(c *gin.Context) {
	var id string
	id = c.Param("id")

	if err := h.Service.RemoveTask(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Removed"})
}
