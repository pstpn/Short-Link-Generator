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

// Создание кеша (в котором в роли ключа выступает короткая ссылка)
// с очисткой устаревших ссылок через 20 секунд
// (предположим, что ссылка действует столько времени)
var firstCache = cache_manager.CacheCreate(20*time.Second, 20*time.Second)

// Создание кеша (в котором в роли ключа выступает оригинальная ссылка)
// с очисткой устаревших ссылок через 20 секунд
// (предположим, что ссылка действует столько времени)
var secondCache = cache_manager.CacheCreate(20*time.Second, 20*time.Second)

type Url struct {
	Data string
}

func TestConnection(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Connection success")
	log.Println("[SUCCESS] Test connection")
}

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
		log.Println("[SUCCESS] Url found in cache: ", shrUrl, "    ||    ", inUrl.Data)
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
		log.Println("[SUCCESS] Url found in database: ", string(answer), "     ||    ", inUrl.Data)
	} else {
		//
		// Если значение не найдено, то генерация новой ссылки и добавление в БД
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
		log.Println("[SUCCESS] Url was generated successfully: ", string(answer), "     ||    ", inUrl.Data)
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
		log.Println("[SUCCESS] Url found in cache: ", origUrl, "    ||    ", inShortUrl.Data)
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
		log.Println("[SUCCESS] Url found in database: ", string(answer), "     ||    ", inShortUrl.Data)
		//
		// Запись ответа
		//
		w.Write(answer)
	} else {
		//
		// Если значение не найдено, то возврат ошибки
		//
		log.Println("[ERROR] Url not found: ")
		http.Error(w, "Error: Url not found (status code: 404)", http.StatusNotFound)
	}
}

func main() {

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
