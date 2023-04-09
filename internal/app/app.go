package app

import (
	"fmt"
	"my_project/urlgen/config"
	"my_project/urlgen/database"
	"my_project/urlgen/internal/server"
	"my_project/urlgen/internal/service"
	"my_project/urlgen/pkg/logger"
	"net/http"
)

// Run - Функция запуска проекта
func Run(config *config.Config) error {

	l := logger.New(config.Log.Level)

	// Подключение к БД
	db, err := database.GetConnection()
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - database.GetConnection: %w", err))
		return err
	}
	defer db.CloseConnection()

	// Создание сервиса
	newService, err := service.NewService(db)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - service.NewService: %w", err))
		return err
	}

	// Создание сервера
	newServer, err := server.NewServer(newService, nil)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - server.NewServer: %w", err))
		return err
	}

	// Запуск сервера
	err = http.ListenAndServe(config.HTTP.Port, newServer.GetRouter())
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - http.ListenAndServe: %w", err))
		return err
	}

	return nil
}
