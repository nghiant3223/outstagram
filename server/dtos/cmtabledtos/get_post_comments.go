package cmtabledtos

import "outstagram/server/dtos"

type GetPostCommentsRequest struct {
	Limit  uint `form:"limit"`
	Offset uint `form:"offset"`
}

type GetPostCommentsResponse struct {
	Comments     []dtos.Comment `json:"comments"`
	CommentCount int            `json:"commentCount"`
}
