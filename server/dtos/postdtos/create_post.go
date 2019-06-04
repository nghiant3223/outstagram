package postdtos

import (
	"outstagram/server/enums/postenums"
)

type CreatePostResponse struct {
	Content    *string              `json:"content" binding:"required"`
	NumRead    int                  `json:"numRead" binding:"required"`
	Visibility postenums.Visibility `json:"visibility" binding:"required"`
	PostImages []PostImage          `json:"postImages" binding:"required"`
}

type CreatePostRequest struct {
	Content    *string              `form:"content"`
	Visibility postenums.Visibility `form:"visibility"`
}
