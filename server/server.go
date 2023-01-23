package main

import (
	"encoding/json"
	"fmt"
	"log"
	"my_project/urlgen/cache_manager"
	"my_project/urlgen/config"
	"my_project/urlgen/database"
	"my_project/urlgen/generator"
	"net/http"
	"time"
)

/*
Создание кеша (в котором в роли ключа выступает короткая ссылка)
с очисткой устаревших ссылок через 20 секунд
(предположим, что ссылка действует столько времени)
*/
var firstCache = cache_manager.CacheCreate(20*time.Second, 20*time.Second)

/*
Создание кеша (в котором в роли ключа выступает оригинальная ссылка)
с очисткой устаревших ссылок через 20 секунд
(предположим, что ссылка действует столько времени)
*/
var secondCache = cache_manager.CacheCreate(20*time.Second, 20*time.Second)

// Url - Тип данных, описывающий структуру для представления ссылки
type Url struct {
	Data string
}

// TestConnection - Функция проверки работоспособности сервера
func TestConnection(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Connection success")
	log.Println("[SUCCESS] Test connection")
}

// GetShortUrl - Функция, реализующая обработку "Post" запроса на сервер (возврат сокращенной ссылки)
func GetShortUrl(w http.ResponseWriter, r *http.Request) {

	inUrl := Url{}

	//
	// Получение входной ссылки
	//
	err := json.NewDecoder(r.Body).Decode(&inUrl)
	if err != nil {
		http.Error(w, "Error: Bad request (status code: 400)", http.StatusBadRequest)
		log.Println("[ERROR] Bad request")
		return
	}

	var answer []byte

	//
	// Поиск в кеше
	//
	if shrUrl, isExist := secondCache.Get(inUrl.Data); isExist {
		answer = []byte(shrUrl)
		log.Println("[SUCCESS] Url found in cache: ", shrUrl, "(In URL: ", inUrl.Data, ")")
		w.Write(answer)
		return
	}

	//
	// Подключение к БД
	//
	db, err := database.GetConnection()
	if err != nil {
		http.Error(w, "Error: Failed to connect to database (status code: 408)", http.StatusRequestTimeout)
		log.Println("[ERROR] Failed to connect to database")
		return
	}

	defer func() {
		err = db.CloseConnection()
		if err != nil {
			http.Error(w, "Error: Failed to connect to database (status code: 408)", http.StatusRequestTimeout)
			log.Println("[ERROR] Failed to close database")
			return
		}
	}()

	//
	// Поиск в БД
	//
	row, isExist := db.GetUrlRow(inUrl.Data, false)
	if isExist {
		answer = []byte(row.ShortUrl)
		log.Println("[SUCCESS] Url found in database: ", string(answer), "(In URL: ", inUrl.Data, ")")
	} else {
		//
		// Генерация новой ссылки с последующим добавлением в БД, если значение не найдено
		//
		answer = []byte(generator.GenerateShortUrl(inUrl.Data))
		err = db.SaveShortUrl(database.RowData{
			Id:       0,
			Url:      inUrl.Data,
			ShortUrl: string(answer),
		})
		if err != nil {
			http.Error(w, "Error: Failed to add data to database (status code: 304)", http.StatusNotModified)
			log.Println("[ERROR] Failed to add data to database")
			return
		}
		log.Println("[SUCCESS] Url was generated successfully: ", string(answer), "(In URL: ", inUrl.Data, ")")
	}

	//
	// Добавление новых значений в кеш
	//
	firstCache.Set(string(answer), inUrl.Data, 0)
	secondCache.Set(inUrl.Data, string(answer), 0)

	//
	// Запись ответа
	//
	w.Write(answer)
}

// GetOriginalUrl - Функция, реализующая обработку "Get" запроса на сервер (возврат исходной ссылки, если она есть)
func GetOriginalUrl(w http.ResponseWriter, r *http.Request) {

	inShortUrl := Url{}

	//
	// Получение входной короткой ссылки
	//
	err := json.NewDecoder(r.Body).Decode(&inShortUrl)
	if err != nil {
		http.Error(w, "Error: Bad request (status code: 400)", http.StatusBadRequest)
		return
	}

	var answer []byte

	//
	// Поиск в кеше
	//
	if origUrl, isExist := firstCache.Get(inShortUrl.Data); isExist {
		answer = []byte(origUrl)
		log.Println("[SUCCESS] Url found in cache: ", origUrl, "(Short URL: ", inShortUrl.Data, ")")
		w.Write(answer)
		return
	}

	//
	// Подключение к БД
	//
	db, err := database.GetConnection()
	if err != nil {
		http.Error(w, "Error: Failed to connect to database (status code: 408)", http.StatusRequestTimeout)
		log.Println("[ERROR] Failed to connect to database")
		return
	}

	defer func() {
		err = db.CloseConnection()
		if err != nil {
			http.Error(w, "Error: Failed to connect to database (status code: 408)", http.StatusRequestTimeout)
			log.Println("[ERROR] Failed to close database")
			return
		}
	}()

	//
	// Поиск в БД
	//
	row, isExist := db.GetUrlRow(inShortUrl.Data, true)
	if isExist {
		answer = []byte(row.Url)
		log.Println("[SUCCESS] Url found in database: ", string(answer), "(Short URL: ", inShortUrl.Data, ")")
		//
		// Добавление значений в кеш
		//
		firstCache.Set(inShortUrl.Data, string(answer), 0)
		secondCache.Set(string(answer), inShortUrl.Data, 0)
		//
		// Запись ответа
		//
		w.Write(answer)
	} else {
		//
		// Возврат ошибки, если значение не найдено
		//
		log.Println("[ERROR] Url not found")
		http.Error(w, "Error: Url not found (status code: 404)", http.StatusNotFound)
	}
}

// Главная функция проекта
func main() {

	//
	// Реализованные обработчики запросов
	//
	http.HandleFunc("/", TestConnection)
	http.HandleFunc("/getshort", GetShortUrl)
	http.HandleFunc("/getoriginal", GetOriginalUrl)

	err := http.ListenAndServe(config.ServerPort, nil)
	if err != nil {
		log.Println("[ERROR] Failed to start server")
		log.Fatal(err.Error())
		return
	}
}
