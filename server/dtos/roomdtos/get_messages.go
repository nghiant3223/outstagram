package roomdtos

import "outstagram/server/dtos/dtomodels"

type GetMessagesRequest struct {
	Limit  uint `form:"limit"`
	Offset uint `form:"offset"`
}

type GetMessageResponse struct {
	Messages []dtomodels.Message `json:"message"`
}
