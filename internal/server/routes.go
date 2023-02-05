package server

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"my_project/urlgen/database"
	"my_project/urlgen/pkg/generator"
	"net/http"
)

// InitRoutes - Метод, инициализирующий обработчики запросов
func (s *Server) initRoutes() {
	s.router.POST("/get-short", s.GetShortUrl)
	s.router.GET("/get-original", s.GetOriginalUrl)
}

// GetShortUrl - Метод, реализующий обработку "Post" запроса на сервер (возврат сокращенной ссылки)
func (s *Server) GetShortUrl(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	inUrl := Url{}

	// Получение входной ссылки
	err := json.NewDecoder(r.Body).Decode(&inUrl)
	if err != nil {
		http.Error(w, "Error: Failed to read request (status code: 400)", http.StatusBadRequest)
		log.Println("[ERROR] Failed to read request")
		return
	}

	var answer []byte

	// Поиск в кеше
	if shrUrl, isExist := s.cacheWithOriginalUrlKey.Get(inUrl.Data); isExist {
		answer = []byte(shrUrl)

		log.Println("[SUCCESS] Url found in cache: ", shrUrl, "(In URL: ", inUrl.Data, ")")

		_, err := w.Write(answer)
		if err != nil {
			http.Error(w, "Error: Failed to write response (status code: 500)", http.StatusInternalServerError)
			log.Println("[ERROR] Failed to write response")
		}

		return
	}

	// Поиск в БД
	row, isExist := s.db.GetUrlRow(inUrl.Data)
	if isExist {
		answer = []byte(row.ShortUrl)

		log.Println("[SUCCESS] Url found in database: ", string(answer), "(In URL: ", inUrl.Data, ")")
	} else {

		// Генерация новой ссылки с последующим добавлением в БД, если значение не найдено
		answer = []byte(generator.GenerateShortUrl(inUrl.Data))
		err = s.db.SaveShortUrl(database.RowData{
			Id:       0,
			Url:      inUrl.Data,
			ShortUrl: string(answer),
		})
		if err != nil {
			http.Error(w, "Error: Failed to save url in database (status code: 500)", http.StatusInternalServerError)
			log.Println("[ERROR] Failed to save url in database")
			return
		}

		log.Println("[SUCCESS] Url was generated successfully: ", string(answer), "(In URL: ", inUrl.Data, ")")
	}

	// Добавление новых значений в кеш
	s.cacheWithShortUrlKey.Set(string(answer), inUrl.Data, 0)
	s.cacheWithOriginalUrlKey.Set(inUrl.Data, string(answer), 0)

	// Запись ответа
	_, err = w.Write(answer)
	if err != nil {
		http.Error(w, "Error: Failed to write response (status code: 500)", http.StatusInternalServerError)
		log.Println("[ERROR] Failed to write response")
	}
}

// GetOriginalUrl - Метод, реализующий обработку "Get" запроса на сервер (возврат исходной ссылки, если она есть)
func (s *Server) GetOriginalUrl(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	inShortUrl := Url{}

	// Получение входной короткой ссылки
	err := json.NewDecoder(r.Body).Decode(&inShortUrl)
	if err != nil {
		http.Error(w, "Error: Failed to read request (status code: 400)", http.StatusBadRequest)
		log.Println("[ERROR] Failed to read request")
		return
	}

	var answer []byte

	// Поиск в кеше
	if origUrl, isExist := s.cacheWithShortUrlKey.Get(inShortUrl.Data); isExist {
		answer = []byte(origUrl)

		log.Println("[SUCCESS] Url found in cache: ", origUrl, "(Short URL: ", inShortUrl.Data, ")")

		_, err := w.Write(answer)
		if err != nil {
			http.Error(w, "Error: Failed to write response (status code: 500)", http.StatusInternalServerError)
			log.Println("[ERROR] Failed to write response")
		}

		return
	}

	// Поиск в БД
	row, isExist := s.db.GetShortUrlRow(inShortUrl.Data)
	if isExist {
		answer = []byte(row.Url)

		log.Println("[SUCCESS] Url found in database: ", string(answer), "(Short URL: ", inShortUrl.Data, ")")

		// Добавление значений в кеш
		s.cacheWithShortUrlKey.Set(inShortUrl.Data, string(answer), 0)
		s.cacheWithOriginalUrlKey.Set(string(answer), inShortUrl.Data, 0)

		// Запись ответа
		_, err := w.Write(answer)
		if err != nil {
			http.Error(w, "Error: Failed to write response (status code: 500)", http.StatusInternalServerError)
			log.Println("[ERROR] Failed to write response")
		}
	} else {

		// Возврат ошибки, если значение не найдено
		http.Error(w, "Error: Url not found (status code: 404)", http.StatusNotFound)
		log.Println("[ERROR] Url not found")
	}
}
