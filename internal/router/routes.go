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

	auth := r.Group("/auth")
	{
		auth.POST("/register", h.SignUpUser)
	}

	api := r.Group("/task")
	{
		api.POST("/add", h.AddTask)
		api.GET("/get-all", h.AllTasks)
		api.PUT("/edit/:id", h.EditTask)
		api.DELETE("/remove/:id", h.RemoveTask)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    "works!",
			"message": "route not found!"})
	})

	return r
}
