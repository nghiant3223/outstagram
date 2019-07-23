package roomdtos

import "outstagram/server/dtos/dtomodels"

type CreateMessageRequest struct {
	Content string `form:"content"`
	Type    int8   `gorm:"type"`
}

type CreateMessageResponse struct {
	Message dtomodels.Message `json:"message"`
}
