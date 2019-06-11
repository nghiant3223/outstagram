package dtos

import (
	"outstagram/server/dtos/postdtos"
	"outstagram/server/enums/postvisibility"
	"time"
)

type Post struct {
	ID            uint                      `json:"id"`
	CreatedAt     time.Time                 `json:"createdAt"`
	Images        []postdtos.PostImage      `json:"images"`
	ImageCount    int                       `json:"imageCount"`
	Comments      []Comment                 `json:"comments"`
	CommentCount  int                       `json:"commentCount"`
	Visibility    postVisibility.Visibility `json:"visibility"`
	Content       *string                   `json:"content"`
	NumViewed     int                       `json:"numViewed"`
	Reactors      []string                  `json:"reactors"`
	ReactCount    int                       `json:"reactCount"`
	OwnerFullname string                    `json:"ownerFullname"`
	OwnerID       uint                      `json:"ownerID"`
}
