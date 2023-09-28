package handler

import (
	"github.com/bahodurnazarov/to-do/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AddTask(c *gin.Context) {
	var task models.Tasks
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}
	if err := h.Service.AddTask(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Added New Task"})
}

func (h *Handler) AllTasks(c *gin.Context) {
	tasks, err := h.Service.AllTasks()
	if len(tasks) == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "message": "Not found any record"})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": tasks})
}

func (h *Handler) EditTask(c *gin.Context) {
	var updateTask models.Tasks
	if err := c.BindJSON(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	var id string
	id = c.Param("id")
	if err := h.Service.EditTask(updateTask, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Task Updated"})
}

func (h *Handler) RemoveTask(c *gin.Context) {
	var id string
	id = c.Param("id")

	if err := h.Service.RemoveTask(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Task Removed"})
}
