package cmtablecontroller

import (
	"github.com/jinzhu/gorm"
	"net/http"
	postVisibility "outstagram/server/enums/postprivacy"
	"outstagram/server/services/cmtableservice"
	"outstagram/server/services/cmtservice"
	"outstagram/server/services/rctableservice"
	"outstagram/server/services/replyservice"
	"outstagram/server/services/userservice"
	"outstagram/server/utils"
)

type Controller struct {
	commentableService *cmtableservice.CommentableService
	commentService     *cmtservice.CommentService
	userService        *userservice.UserService
	reactableService   *rctableservice.ReactableService
	replyService       *replyservice.ReplyService
}

func New(commentableService *cmtableservice.CommentableService, commentService *cmtservice.CommentService, userService *userservice.UserService, reactableService *rctableservice.ReactableService, replyService *replyservice.ReplyService) *Controller {
	return &Controller{commentableService: commentableService, commentService: commentService, userService: userService, reactableService: reactableService, replyService: replyService}
}

// checkUserAuthorizationForCommentable checks if user has the authorization to see the post
func (cc *Controller) checkUserAuthorizationForCommentable(cmtableID, userID uint) *utils.HttpError {
	visibility, ownerID, err := cc.commentableService.GetVisibilityByID(cmtableID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return utils.NewHttpError(http.StatusNotFound, "Commentable not found", err.Error())
		}

		return utils.NewHttpError(http.StatusInternalServerError, "Error while retrieving post", err.Error())
	}

	if visibility == postVisibility.Private {
		if ownerID != userID {
			return utils.NewHttpError(http.StatusForbidden, "Cannot access this commentable", nil)
		}
	} else if visibility == postVisibility.OnlyFollowers {
		followed, err := cc.userService.CheckFollow(userID, ownerID)
		if err != nil {
			return utils.NewHttpError(http.StatusInternalServerError, "Error while checking follow", err.Error())
		}

		if !followed {
			return utils.NewHttpError(http.StatusForbidden, "Cannot access this commentable", nil)

		}
	} else if visibility != postVisibility.Public {
		return utils.NewHttpError(http.StatusConflict, "Invalid visibility of a commentable", visibility)
	}

	return nil
}
