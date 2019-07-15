package rctdtos

import "outstagram/server/dtos/dtomodels"

type GetReactionsRequest struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

type GetReactionsResponse struct {
	Reactors   []dtomodels.SimpleUser `json:"reactors"`
	ReactCount int                    `json:"reactCount"`
}
