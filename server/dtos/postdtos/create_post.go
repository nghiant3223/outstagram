package postdtos

import (
	"outstagram/server/enums/postvisibility"
)

type CreatePostRequest struct {
	Content    *string                   `form:"content"`
	Visibility postVisibility.Visibility `form:"visibility"`
}

type CreatePostResponse struct {
	ID         uint                      `json:"id"`
	Content    *string                   `json:"content"`
	NumViewed  int                       `json:"numViewed"`
	Visibility postVisibility.Visibility `json:"visibility"`
	Images     []PostImage               `json:"images"`
}
