package main

import (
	"log"
	"my_project/urlgen/config"
	"my_project/urlgen/database"
	"my_project/urlgen/internal/server"
	"net/http"
)

// Главная функция проекта
func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

// Функция запуска проекта
func run() error {

	// Подключение к БД
	db, err := database.GetConnection()
	if err != nil {
		log.Println("[ERROR] Failed to connect to database")
		return err
	}
	defer func() {
		err = db.CloseConnection()
		if err != nil {
			log.Fatal("[ERROR] Failed to close database")
		}
	}()

	// Создание сервера
	newServer, err := server.NewServer(&db, nil)
	if err != nil {
		log.Println("[ERROR] Failed to create server")
		return err
	}

	// Запуск сервера
	err = http.ListenAndServe(config.ServerPort, newServer.GetRouter())
	if err != nil {
		log.Println("[ERROR] Failed to start server")
		return err
	}

	return nil
}
