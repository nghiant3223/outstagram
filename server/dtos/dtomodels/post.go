package dtomodels

import (
	"outstagram/server/enums/postprivacy"
	"time"
)

type Post struct {
	ID            uint                `json:"id"`
	CreatedAt     time.Time           `json:"createdAt"`
	Images        []SimplePostImage   `json:"images,omitempty"`
	ImageID       *uint               `json:"imageID,omitempty"`
	ImageCount    int                 `json:"imageCount"`
	Comments      []Comment           `json:"comments"`
	CommentCount  int                 `json:"commentCount"`
	Visibility    postPrivacy.Privacy `json:"visibility"`
	Content       *string             `json:"content"`
	Reactors      []SimpleUser        `json:"reactors"`
	ReactCount    int                 `json:"reactCount"`
	OwnerFullname string              `json:"ownerFullname"`
	OwnerUsername string              `json:"ownerUsername"`
	OwnerID       uint                `json:"ownerID"`
	CommentableID uint                `json:"commentableID"`
	ReactableID   uint                `json:"reactableID"`
	ViewableID    uint                `json:"viewableID"`
	Reacted       bool                `json:"reacted"`
}
