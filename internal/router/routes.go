package router

import (
	"github.com/bahodurnazarov/to-do/internal/app/handler"
	"github.com/bahodurnazarov/to-do/pkg/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(h *handler.Handler) *gin.Engine {

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/ping", h.Ping)
	r.POST("/addTask", h.AddTask)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": "works!", "message": "not route!"})
	})

	return r
}
