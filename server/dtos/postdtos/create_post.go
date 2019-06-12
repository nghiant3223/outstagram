package postdtos

import (
	"outstagram/server/dtos/dtomodels"
	"outstagram/server/enums/postprivacy"
)

type CreatePostRequest struct {
	Content    *string             `form:"content"`
	Visibility postPrivacy.Privacy `form:"visibility"`
}

type CreatePostResponse struct {
	//ID         uint                      `json:"id"`
	//Content    *string                   `json:"content"`
	//NumViewed  int                       `json:"numViewed"`
	//Privacy postVisibility.Privacy `json:"visibility"`
	//Images     []dtomodels.PostImage     `json:"images"`
	Post 	dtomodels.Post `json:"post"`
}
