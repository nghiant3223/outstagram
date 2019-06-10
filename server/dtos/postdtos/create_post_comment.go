package postdtos

type CreateCommentRequest struct {
	Content *string `form:"content" required:"true"`
}