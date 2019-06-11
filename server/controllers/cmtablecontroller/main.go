package cmtablecontroller

import (
	"outstagram/server/dtos/cmtabledtos"
	"outstagram/server/models"
	"outstagram/server/services/cmtableservice"
	"outstagram/server/services/cmtservice"
	"outstagram/server/services/rctableservice"
	"outstagram/server/services/userservice"
)

type Controller struct {
	commentableService *cmtableservice.CommentableService
	commentService     *cmtservice.CommentService
	userService        *userservice.UserService
	reactableService   *rctableservice.ReactableService
}

func New(commentableService *cmtableservice.CommentableService, commentService *cmtservice.CommentService, userService *userservice.UserService, reactableService *rctableservice.ReactableService) *Controller {
	return &Controller{commentableService: commentableService, commentService: commentService, userService: userService, reactableService: reactableService}
}

func (cc *Controller) getDTOComment(comment *models.Comment) cmtabledtos.Comment {
	return cmtabledtos.Comment{
		ID:            comment.ID,
		Content:       comment.Content,
		ReplyCount:    comment.ReplyCount,
		CreatedAt:     comment.CreatedAt,
		OwnerFullname: comment.User.Fullname,
		OwnerID:       comment.UserID,
		ReactCount:    cc.reactableService.GetReactCount(comment.ReactableID),
		Reactors:      cc.reactableService.GetReactors(comment.ReactableID)}
}
