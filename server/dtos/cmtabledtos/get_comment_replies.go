package cmtabledtos

import "outstagram/server/dtos"

type GetCommentRepliesRequest struct {
	Limit  uint `form:"limit"`
	Offset uint `form:"offset"`
}

type GetCommentRepliesResponse struct {
	Replies    []dtos.Reply `json:"replies"`
	ReplyCount int          `json:"replyCount"`
}
