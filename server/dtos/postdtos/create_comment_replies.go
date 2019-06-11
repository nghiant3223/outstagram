package postdtos

type CreateReplyRequest struct {
	Content *string `form:"content" required:"true"`
}

type CreateReplyResponse struct {
	Reply Reply `json:"reply"`
}
