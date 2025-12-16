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
	storage := storage.NewMemoryStorage()

	eventHandler := handler.NewEventHandler(storage)

	mux := http.NewServeMux()

	//  events with POST and GET
	mux.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			eventHandler.CreateEvent(w, r)
		} else if r.Method == http.MethodGet {
			eventHandler.ListEvents(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	loggedMux := middleware.Logging(mux)

	// server
	port := ":8080"
	fmt.Println("Server started on port", port)
	err := http.ListenAndServe(port, loggedMux)
	if err != nil {
		log.Fatal(err)
	}
}
