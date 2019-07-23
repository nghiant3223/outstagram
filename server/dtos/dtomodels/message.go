package dtomodels

import "time"

type Message struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Content   string    `json:"content"`
	Type      int8      `json:"type"`
	RoomID    uint      `json:"roomID"`
	AuthorID  uint      `json:"authorID"`
}
