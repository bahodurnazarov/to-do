package main

import (
	"github.com/bahodurnazarov/to-do/pkg/app"
	"github.com/bahodurnazarov/to-do/pkg/config"
	"github.com/bahodurnazarov/to-do/pkg/db"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalln("Failed init cfg", err.Error())
	}
	a := app.New(cfg)
	defer db.CloseDB()
	log.Fatalln(a.Run(cfg))
}
