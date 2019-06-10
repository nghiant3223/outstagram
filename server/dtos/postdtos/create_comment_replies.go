package postdtos

type CreateReplyRequest struct {
	Content *string `form:"content" required:"true"`
}
