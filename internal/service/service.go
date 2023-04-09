package service

import (
	"errors"
	"my_project/urlgen/database"
	"my_project/urlgen/pkg/cache_manager"
	"my_project/urlgen/pkg/hash_generator"
)

// Service - тип данных для работы
type Service struct {
	db                *database.Database
	cacheWithOriginal *cache_manager.Cache
	cacheWithShort    *cache_manager.Cache
}

// NewService - Метод, реализующий создание нового сервиса
func NewService(db *database.Database) (*Service, error) {

	// Создание кеша для полных ссылок
	cacheWithOriginal := cache_manager.CacheCreate()

	// Создание кеша для сокращенных ссылок
	cacheWithShort := cache_manager.CacheCreate()

	return &Service{
		db:                db,
		cacheWithOriginal: cacheWithOriginal,
		cacheWithShort:    cacheWithShort,
	}, nil
}

// CreateShortUrl - Функция, реализующая создание короткой ссылке по заданной полной
func CreateShortUrl(s *Service, inUrl string) ([]byte, error) {

	// Поиск в кеше
	if shrUrl, isExist := s.cacheWithOriginal.Get(inUrl); isExist {
		return []byte(shrUrl), nil
	}

	// Поиск в БД
	var answer []byte

	row, isExist := s.db.GetUrlRow(inUrl)
	if isExist {
		answer = []byte(row.ShortUrl)
	} else {

		// Генерация новой ссылки с последующим добавлением в БД, если значение не найдено
		answer = []byte(hash_generator.GenerateShortUrl(inUrl))
		err := s.db.SaveShortUrl(database.RowData{
			Id:       0,
			Url:      inUrl,
			ShortUrl: string(answer),
		})
		if err != nil {
			return nil, err
		}
	}

	// Добавление новых значений в кеш
	s.cacheWithShort.Set(string(answer), inUrl, 0)
	s.cacheWithOriginal.Set(inUrl, string(answer), 0)

	return answer, nil
}

// FindOriginalUrl - Функция, реализующая поиск полной ссылки по заданной короткой
func FindOriginalUrl(s *Service, inShortUrl string) ([]byte, error) {

	// Поиск в кеше
	if origUrl, isExist := s.cacheWithShort.Get(inShortUrl); isExist {
		return []byte(origUrl), nil
	}

	// Поиск в БД
	row, isExist := s.db.GetShortUrlRow(inShortUrl)
	if !isExist {
		return nil, errors.New("url not found")
	}

	// Добавление значений в кеш
	s.cacheWithShort.Set(inShortUrl, row.Url, 0)
	s.cacheWithOriginal.Set(row.Url, inShortUrl, 0)

	return []byte(row.Url), nil
}
