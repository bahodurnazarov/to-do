package main

import (
	_ "github.com/bahodurnazarov/to-do/cmd/app/docs"
	"github.com/bahodurnazarov/to-do/pkg/app"
	"github.com/bahodurnazarov/to-do/pkg/config"
	"github.com/bahodurnazarov/to-do/pkg/db"
	"log"
)

//@title To DO
//@version 1.0
//@description To-do List
//@termsOfService  http://localhost:8089/swagger/index.html#/

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalln("Failed init cfg", err.Error())
	}
	a := app.New(cfg)
	defer db.CloseDB()
	log.Fatalln(a.Run(cfg))
}
