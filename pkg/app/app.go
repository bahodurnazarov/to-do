package app

import (
	"github.com/bahodurnazarov/to-do/pkg/db"
	"github.com/bahodurnazarov/to-do/pkg/models"
	"log"
)

type App struct {
}

func New(cfg *models.Settings) *App {
	app := &App{}
	connDB, err := db.ConnectDB(cfg)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("connect :", connDB)

	return app
}
