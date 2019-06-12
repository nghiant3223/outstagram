package cmtabledtos

import "outstagram/server/dtos/dtomodels"

type CreateCommentRequest struct {
	Content *string `form:"content" required:"true"`
}

type CreateCommentResponse struct {
	Comment dtomodels.Comment `json:"comment"`
}
