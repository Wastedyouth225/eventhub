package storage

import (
	"sync"
	"time"

	"eventhub/internal/model"
)

type MemoryStorage struct {
	mu     sync.Mutex
	events []model.Event
	nextID int64
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		events: []model.Event{},
		nextID: 1,
	}
}

func (s *MemoryStorage) AddEvent(event model.Event) model.Event {
	s.mu.Lock()
	defer s.mu.Unlock()

	event.ID = s.nextID
	event.CreatedAt = time.Now() // позже объясню почему
	s.nextID++
	s.events = append(s.events, event)
	return event
}

func (s *MemoryStorage) GetEvents() []model.Event {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.events
}
