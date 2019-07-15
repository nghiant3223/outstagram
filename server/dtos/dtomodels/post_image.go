package dtomodels

import (
	"time"
)

type PostImage struct {
	ID            uint         `json:"id"`
	CreatedAt     time.Time    `json:"createdAt"`
	ImageID       uint         `json:"imageID"`
	Comments      []Comment    `json:"comments"`
	CommentCount  int          `json:"commentCount"`
	Content       *string      `json:"content"`
	Reactors      []SimpleUser `json:"reactors"`
	ReactCount    int          `json:"reactCount"`
	OwnerID       uint         `json:"ownerID"`
	OwnerFullname string       `json:"ownerFullname"`
	OwnerUsername string       `json:"ownerUsername"`
	CommentableID uint         `json:"commentableID"`
	ReactableID   uint         `json:"reactableID"`
	ViewableID    uint         `json:"viewableID"`
	Reacted       bool         `json:"reacted"`
}

type SimplePostImage struct {
	ID      uint `json:"id"`
	ImageID uint `json:"imageID"`
}
