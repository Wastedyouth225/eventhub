package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"eventhub/internal/model"
	"eventhub/internal/storage"
)

type EventHandler struct {
	storage *storage.MemoryStorage
}

// Конструктор handler
func NewEventHandler(s *storage.MemoryStorage) *EventHandler {
	return &EventHandler{storage: s}
}

// POST /events
func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var e model.Event
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	created := h.storage.AddEvent(e)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

// GET /events с пагинацией
func (h *EventHandler) ListEvents(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limit, _ := strconv.Atoi(query.Get("limit"))
	offset, _ := strconv.Atoi(query.Get("offset"))

	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	events := h.storage.GetEvents()

	if offset > len(events) {
		offset = len(events)
	}
	end := offset + limit
	if end > len(events) {
		end = len(events)
	}

	result := events[offset:end]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
