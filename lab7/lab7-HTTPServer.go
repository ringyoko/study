package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// обрабатывает HTTP-запросы и возвращает ошибку.
type HandlerFunc func(http.ResponseWriter, *http.Request) error

// функция-обертка для обработчиков.
type Middleware func(HandlerFunc) HandlerFunc

// логирует запросы.
func loggingMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		start := time.Now()
		log.Printf("Method: %s, URL: %s, Started at: %s", r.Method, r.URL.Path, start.Format(time.RFC3339))

		err := next(w, r)

		log.Printf("Completed in %v", time.Since(start))
		return err
	}
}

// обрабатывает GET-запросы /hello.
func helloHandler(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Hello, welcome to our server!")
	return err
}

// структура для данных, отправляемых в POST-запросе.
type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// обрабатывает POST-запросы /data.
func dataHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return nil
	}

	var data Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return nil
	}
	defer r.Body.Close()

	log.Printf("Received POST data: %+v", data) // Вывод структуры в консоль
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Data received successfully")
	return err
}

// преобразует HandlerFunc в http.HandlerFunc.
func adapter(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			log.Printf("Error handling request: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func main() {
	mux := http.NewServeMux()
	var loggingMW Middleware = loggingMiddleware

	mux.HandleFunc("/hello", adapter(loggingMW(helloHandler)))
	mux.HandleFunc("/data", adapter(loggingMW(dataHandler)))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server is running on port 8080...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// GET в POSTMAN: http://localhost:8080/hello
// POST в POSTMAN: http://localhost:8080/data . В raw: {
//    "name": "John",
//    "age": 111
//}
