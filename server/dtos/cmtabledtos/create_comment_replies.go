package cmtabledtos

import "outstagram/server/dtos"

type CreateReplyRequest struct {
	Content *string `form:"content" required:"true"`
}

type CreateReplyResponse struct {
	Reply dtos.Reply `json:"reply"`
}
