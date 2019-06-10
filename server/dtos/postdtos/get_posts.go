package postdtos

type GetPostRequest struct {
	Limit  uint `form:"limit"`
	Offset uint `form:"offset"`
}

type GetPostResponse struct {
	Posts []Post `json:"posts"`
}
