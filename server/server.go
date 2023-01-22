package main

import (
	"encoding/json"
	"fmt"
	"log"
	"my_project/urlgen/config"
	"my_project/urlgen/generator"
	"net/http"
)

type Url struct {
	Data string
}

func TestConnection(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Connection success")
}

func GetShortUrl(w http.ResponseWriter, r *http.Request) {
	inUrl := Url{}

	err := json.NewDecoder(r.Body).Decode(&inUrl)
	if err != nil {
		http.Error(w, "Error: Bad request (status code: 400)", http.StatusBadRequest)
		return
	}

	// Поиск в кеше

	// Поиск в БД

	// Если не найдено, то генерация новой ссылки и добавление в БД

	b, _ := json.Marshal(generator.GenerateShortUrl(inUrl.Data))

	w.Write(b)
}

func GetOriginalUrl(w http.ResponseWriter, r *http.Request) {
	inShortUrl := Url{}

	err := json.NewDecoder(r.Body).Decode(&inShortUrl)
	if err != nil {
		http.Error(w, "Error: Bad request (status code: 400)", http.StatusBadRequest)
		return
	}

	// Поиск в кеше

	// Поиск в БД

	// Если не найдена, то возврат ошибки
}

func main() {
	http.HandleFunc("/", TestConnection)
	http.HandleFunc("/getshort", GetShortUrl)

	err := http.ListenAndServe(config.ServerPort, nil)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//
	// HTTPS Server Version
	//
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", HelloAnswer)
	//
	//certManager := autocert.Manager{
	//	Prompt: autocert.AcceptTOS,
	//	Cache:  autocert.DirCache(".cache/certs"),
	//}
	//
	//server := &http.Server{
	//	Addr:    ":443",
	//	Handler: mux,
	//	TLSConfig: &tls.Config{
	//		GetCertificate: certManager.GetCertificate,
	//	},
	//}
	//
	//go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
	//server.ListenAndServeTLS("", "")
}
