package server

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"my_project/urlgen/internal/service"
	"net/http"
)

// Url - Тип данных, описывающий структуру для представления ссылки
type Url struct {
	Data string `json:"data"` // Ссылка
}

// initRoutes - Метод, инициализирующий обработчики запросов
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
		return
	}

	answer, err := service.CreateShortUrl(s.service, inUrl.Data)
	if err != nil {
		http.Error(w, "Error: Failed to get short url (status code: 500)", http.StatusInternalServerError)
		return
	}

	// Запись ответа
	_, err = w.Write(answer)
	if err != nil {
		http.Error(w, "Error: Failed to write response (status code: 500)", http.StatusInternalServerError)
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

	answer, err := service.FindOriginalUrl(s.service, inShortUrl.Data)
	if err != nil {
		http.Error(w, "Error: Failed to find original url (status code: 404)", http.StatusNotFound)
		return
	}

	// Запись ответа
	_, err = w.Write(answer)
	if err != nil {
		http.Error(w, "Error: Failed to write response (status code: 500)", http.StatusInternalServerError)
	}
}
