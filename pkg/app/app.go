package app

import (
	"github.com/bahodurnazarov/to-do/internal/app/handler"
	"github.com/bahodurnazarov/to-do/internal/app/repository"
	"github.com/bahodurnazarov/to-do/internal/app/service"
	"github.com/bahodurnazarov/to-do/internal/router"
	"github.com/bahodurnazarov/to-do/pkg/config"
	"github.com/bahodurnazarov/to-do/pkg/db"
	"github.com/bahodurnazarov/to-do/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	r *repository.Repository
	s *service.Service
	h *handler.Handler
	g *gin.Engine
}

func New(cfg *models.Settings) *App {
	app := &App{}
	connDB, err := db.ConnectDB(cfg)
	if err != nil {
		log.Println(err)
		return nil
	}

	app.r = repository.New(connDB)
	app.s = service.New(app.r)
	app.h = handler.New(app.s)
	app.g = router.Init(app.h)

	return app
}

func (a *App) Run(cfg *models.Settings) error {
	//return a.g.Run(net.JoinHostPort(cfg.Server.Host, cfg.Server.Port))
	return a.g.Run(":" + config.Cnfg.App.PortRun)
}
