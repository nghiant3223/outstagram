package cmtabledtos

type GetCommentRepliesRequest struct {
	Limit  uint `form:"limit"`
	Offset uint `form:"offset"`
}

type GetCommentRepliesResponse struct {
	Replies    []Reply `json:"replies"`
	ReplyCount int     `json:"replyCount"`
}
