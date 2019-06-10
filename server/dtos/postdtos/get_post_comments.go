package postdtos

type GetPostCommentsRequest struct {
	Limit  uint `form:"limit"`
	Offset uint `form:"offset"`
}

type GetPostCommentsResponse struct {
	Comments []Comment `json:"comments"`
}