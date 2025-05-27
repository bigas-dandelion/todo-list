package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var dbServiceAddress = "http://db-service:8082"

func proxyRequest(w http.ResponseWriter, r *http.Request) {

	targetURL := fmt.Sprintf("%s%s", dbServiceAddress, r.URL.Path)

	req, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	req.Header = r.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	defer resp.Body.Close()

	// Копируем заголовки ответа от БД-сервиса
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Передаём HTTP-статус от БД-сервиса
	w.WriteHeader(resp.StatusCode)

	// Копируем тело ответа от БД-сервиса клиенту
	io.Copy(w, resp.Body)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /create", proxyRequest)
	mux.HandleFunc("GET /list", proxyRequest)
	mux.HandleFunc("DELETE /delete/{id}", proxyRequest)
	mux.HandleFunc("PATCH /done/{id}", proxyRequest)

	log.Fatal(http.ListenAndServe(":8081", mux))
}
