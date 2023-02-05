package server

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"my_project/urlgen/config"
	"my_project/urlgen/database"
	"my_project/urlgen/pkg/cache_manager"
)

// Url - Тип данных, описывающий структуру для представления ссылки
type Url struct {
	Data string // Ссылка
}

// Server - Тип данных, описывающий структуру сервера
type Server struct {
	context context.Context    // Контекст сервера
	router  *httprouter.Router // Маршрутизатор

	db                      *database.Database   // Подключение к БД
	cacheWithShortUrlKey    *cache_manager.Cache // Кеш с ключами вида "короткая ссылка"
	cacheWithOriginalUrlKey *cache_manager.Cache // Кеш с ключами вида "оригинальная ссылка"
}

// NewServer - Функция, позволяющая создать новый сервер
func NewServer(db *database.Database, ctx context.Context) (*Server, error) {

	// Создание сервера
	s := Server{
		context: ctx,
		router:  httprouter.New(),

		db:                      db,
		cacheWithShortUrlKey:    cache_manager.CacheCreate(config.CacheDefaultExpiration, config.CacheCleanupTime),
		cacheWithOriginalUrlKey: cache_manager.CacheCreate(config.CacheDefaultExpiration, config.CacheCleanupTime),
	}

	// Инициализация маршрутов
	s.initRoutes()

	return &s, nil
}

// GetRouter - Функция, позволяющая получить маршрутизатор сервера
func (s *Server) GetRouter() *httprouter.Router {
	return s.router
}
