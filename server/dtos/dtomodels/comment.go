package dtomodels

import "time"

type Comment struct {
	ID            uint      `json:"id"`
	Content       *string   `json:"content"`
	CreatedAt     time.Time `json:"createdAt"`
	ReplyCount    int       `json:"replyCount"`
	OwnerFullname string    `json:"ownerFullname"`
	OwnerID       uint      `json:"ownerID"`
	OwnerUsername string    `json:"ownerUsername"`
	Reactors      []string  `json:"reactors"`
	ReactCount    int       `json:"reactCount"`
	CommentableID uint      `json:"commentableID"`
	ReactableID   uint      `json:"reactableID"`
	Reacted       bool      `json:"reacted"`
}
