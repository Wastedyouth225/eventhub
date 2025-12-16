package main

import (
	"fmt"
	"log"
	"net/http"

	"eventhub/internal/handler"
	"eventhub/internal/middleware"
	"eventhub/internal/storage"
)

func main() {
	// 1. Создаём хранилище
	storage := storage.NewMemoryStorage()

	// 2. Создаём обработчик для событий
	eventHandler := handler.NewEventHandler(storage)

	// 3. Создаём HTTP mux (маршрутизатор)
	mux := http.NewServeMux()

	// 4. Подключаем /events с POST и GET
	mux.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			eventHandler.CreateEvent(w, r)
		} else if r.Method == http.MethodGet {
			eventHandler.ListEvents(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// 5. Подключаем /health
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// 6. Подключаем middleware логирования
	loggedMux := middleware.Logging(mux)

	// 7. Запуск сервера
	port := ":8080"
	fmt.Println("Server started on port", port)
	err := http.ListenAndServe(port, loggedMux)
	if err != nil {
		log.Fatal(err)
	}
}
