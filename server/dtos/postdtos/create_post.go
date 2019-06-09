package postdtos

import (
	"outstagram/server/enums/postenums"
)

type CreatePostResponse struct {
	ID         uint                 `json:"id"`
	Content    *string              `json:"content"`
	NumRead    int                  `json:"numRead"`
	Visibility postenums.Visibility `json:"visibility"`
	Images     []PostImage          `json:"images"`
}

type CreatePostRequest struct {
	Content    *string              `form:"content"`
	Visibility postenums.Visibility `form:"visibility"`
}
