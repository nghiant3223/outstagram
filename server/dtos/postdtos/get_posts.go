package postdtos

import "outstagram/server/dtos/dtomodels"

type GetPostRequest struct {
	Limit  uint `form:"limit"`
	Offset uint `form:"offset"`
}

type GetPostResponse struct {
	Posts []dtomodels.Post `json:"posts"`
}
