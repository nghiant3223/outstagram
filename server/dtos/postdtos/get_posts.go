package postdtos

import "outstagram/server/dtos"

type GetPostRequest struct {
	Limit  uint `form:"limit"`
	Offset uint `form:"offset"`
}

type GetPostResponse struct {
	Posts []dtos.Post `json:"posts"`
}
