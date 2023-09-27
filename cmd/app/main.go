package main

import (
	"github.com/bahodurnazarov/to-do/pkg/config"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalln("Failed init cfg", err.Error())
	}
	log.Println("Config :", cfg)
}
