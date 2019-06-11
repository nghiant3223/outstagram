package cmtabledtos

import "outstagram/server/dtos"

type CreateCommentRequest struct {
	Content *string `form:"content" required:"true"`
}

type CreateCommentResponse struct {
	Comment dtos.Comment `json:"comment"`
}
