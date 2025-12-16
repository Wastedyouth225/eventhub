package model

import "time"

type Event struct {
	ID        int64     `json:"id"`
	Type      string    `json:"type"`
	Payload   string    `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}
