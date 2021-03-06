package postdtos

import (
	"outstagram/server/dtos/dtomodels"
	"outstagram/server/enums/postprivacy"
)

type CreatePostRequest struct {
	Content    *string             `form:"content"`
	Visibility postPrivacy.Privacy `form:"visibility"`
	ImageURLs  []string            `form:"imageURLs"`
}

type CreatePostResponse struct {
	Post dtomodels.Post `json:"post"`
}
