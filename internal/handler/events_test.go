package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"eventhub/internal/model"
	"eventhub/internal/storage"
)

func TestCreateEvent(t *testing.T) {
	store := storage.NewMemoryStorage()
	h := NewEventHandler(store)

	// test event
	event := model.Event{
		Type:    "test",
		Payload: "hello",
	}
	body, _ := json.Marshal(event)

	req := httptest.NewRequest(http.MethodPost, "/events", bytes.NewReader(body))
	w := httptest.NewRecorder()

	h.CreateEvent(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}

	if len(store.GetEvents()) != 1 {
		t.Errorf("Expected 1 event in storage, got %d", len(store.GetEvents()))
	}
}
