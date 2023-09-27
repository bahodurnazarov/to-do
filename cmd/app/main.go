package main

import (
	"github.com/bahodurnazarov/to-do/pkg/app"
	"github.com/bahodurnazarov/to-do/pkg/config"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalln("Failed init cfg", err.Error())
	}
	a := app.New(cfg)
	log.Println("DataBase :", a)
}
