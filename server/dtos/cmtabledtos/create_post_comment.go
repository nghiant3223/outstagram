package cmtabledtos

type CreateCommentRequest struct {
	Content *string `form:"content" required:"true"`
}

type CreateCommentResponse struct {
	Comment Comment `json:"comment"`
}
