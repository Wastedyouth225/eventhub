package model

import "time"

// Event представляет одно событие в системе
type Event struct {
	ID        int64     `json:"id"`
	Type      string    `json:"type"`
	Payload   string    `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}
