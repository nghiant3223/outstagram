package cmtabledtos

import "outstagram/server/dtos/dtomodels"

type CreateReplyRequest struct {
	Content *string `form:"content" required:"true"`
}

type CreateReplyResponse struct {
	Reply dtomodels.Reply `json:"reply"`
}
