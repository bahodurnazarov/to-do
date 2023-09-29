package router

import (
	_ "github.com/bahodurnazarov/to-do/cmd/app/docs"
	"github.com/bahodurnazarov/to-do/internal/app/handler"
	"github.com/bahodurnazarov/to-do/pkg/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Init(h *handler.Handler) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/ping", h.Ping)
	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8089/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	auth := r.Group("/auth")
	{
		auth.POST("/register", h.SignUpUser)
		auth.POST("/login", h.SignInUser)
	}

	api := r.Group("/task", middleware.TokenMiddleware())
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
