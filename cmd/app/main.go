package main

import (
	"log"
	"my_project/urlgen/config"
	"my_project/urlgen/internal/app"
)

// Главная функция проекта
func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error getting config: %s", err)
	}

	if err := app.Run(cfg); err != nil {
		log.Fatalf(err.Error())
	}
}
