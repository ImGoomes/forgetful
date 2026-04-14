package model

import "time"

type Entry struct {
	ID          int       `json:"id"`
	Command     string    `json:"command"`
	Tag         string    `json:"tag"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

type Store struct {
	Entries []*Entry `json:"entries"`
	NextID  int      `json:"next_id"`
}
