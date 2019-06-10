package postdtos

import (
	"outstagram/server/enums/postenums"
)

type CreatePostRequest struct {
	Content    *string              `form:"content"`
	Visibility postenums.Visibility `form:"visibility"`
}

type CreatePostResponse struct {
	ID         uint                 `json:"id"`
	Content    *string              `json:"content"`
	NumViewed  int                  `json:"numViewed"`
	Visibility postenums.Visibility `json:"visibility"`
	Images     []PostImage          `json:"images"`
}
