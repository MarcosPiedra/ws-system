package models

import "time"

type Message struct {
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}
