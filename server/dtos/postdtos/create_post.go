package postdtos

import (
	"outstagram/server/dtos/dtomodels"
	"outstagram/server/enums/postvisibility"
)

type CreatePostRequest struct {
	Content    *string                   `form:"content"`
	Visibility postVisibility.Visibility `form:"visibility"`
}

type CreatePostResponse struct {
	//ID         uint                      `json:"id"`
	//Content    *string                   `json:"content"`
	//NumViewed  int                       `json:"numViewed"`
	//Visibility postVisibility.Visibility `json:"visibility"`
	//Images     []dtomodels.PostImage     `json:"images"`
	Post 	dtomodels.Post `json:"post"`
}
