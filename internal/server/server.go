package server

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"my_project/urlgen/internal/service"
)

// Server - Тип данных, описывающий структуру сервера
type Server struct {
	context context.Context    // Контекст сервера
	router  *httprouter.Router // Маршрутизатор
	service *service.Service   // Сервис для работы с хранилищем
}

// NewServer - Функция, позволяющая создать новый сервер
func NewServer(service *service.Service, ctx context.Context) (*Server, error) {

	// Создание сервера
	s := Server{
		context: ctx,
		router:  httprouter.New(),
		service: service,
	}

	// Инициализация маршрутов
	s.initRoutes()

	return &s, nil
}

// GetRouter - Функция, позволяющая получить маршрутизатор сервера
func (s *Server) GetRouter() *httprouter.Router {
	return s.router
}
