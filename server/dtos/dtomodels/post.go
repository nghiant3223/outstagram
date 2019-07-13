package dtomodels

import (
	"outstagram/server/enums/postprivacy"
	"time"
)

type Post struct {
	ID            uint                `json:"id"`
	CreatedAt     time.Time           `json:"createdAt"`
	Images        []PostImage         `json:"images"`
	ImageCount    int                 `json:"imageCount"`
	Comments      []Comment           `json:"comments"`
	CommentCount  int                 `json:"commentCount"`
	Visibility    postPrivacy.Privacy `json:"visibility"`
	Content       *string             `json:"content"`
	Reactors      []BasicUser         `json:"reactors"`
	ReactCount    int                 `json:"reactCount"`
	OwnerFullname string              `json:"ownerFullname"`
	OwnerID       uint                `json:"ownerID"`
	CommentableID uint                `json:"commentableID"`
	ReactableID   uint                `json:"reactableID"`
	ViewableID    uint                `json:"viewableID"`
	Reacted       bool                `json:"reacted"`
}
