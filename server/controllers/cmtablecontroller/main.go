package cmtablecontroller

import "outstagram/server/services/cmtableservice"

type Controller struct {
	commentableService *cmtableservice.CommentableService
}

func New(commentableService *cmtableservice.CommentableService) *Controller {
	return &Controller{commentableService: commentableService}
}
