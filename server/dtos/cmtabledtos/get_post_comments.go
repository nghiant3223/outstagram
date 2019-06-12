package cmtabledtos

import "outstagram/server/dtos/dtomodels"

type GetPostCommentsRequest struct {
	Limit  uint `form:"limit"`
	Offset uint `form:"offset"`
}

type GetPostCommentsResponse struct {
	Comments     []dtomodels.Comment `json:"comments"`
	CommentCount int                 `json:"commentCount"`
}
