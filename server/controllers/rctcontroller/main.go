package rctcontroller

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	postVisibility "outstagram/server/enums/postprivacy"
	"outstagram/server/services/rctableservice"
	"outstagram/server/services/rctservice"
	"outstagram/server/services/userservice"
	"outstagram/server/utils"
)

type Controller struct {
	reactService     *rctservice.ReactService
	reactableService *rctableservice.ReactableService
	userService      *userservice.UserService
}

func New(reactService *rctservice.ReactService, reactableService *rctableservice.ReactableService, userService *userservice.UserService) *Controller {
	return &Controller{reactService: reactService, reactableService: reactableService, userService: userService}
}

// checkUserAuthorizationForCommentable checks if user has the authorization to see the post
func (rc *Controller) checkUserAuthorizationForReactable(reactableID, userID uint) *utils.HttpError {
	visibility, ownerID, err := rc.reactableService.GetVisibilityByID(reactableID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return utils.NewHttpError(http.StatusNotFound, "Commentable not found", err.Error())
		}

		fmt.Println(">>")
		
		return utils.NewHttpError(http.StatusInternalServerError, "Error while retrieving reactable", err.Error())
	}

	if visibility == postVisibility.Private {
		if ownerID != userID {
			return utils.NewHttpError(http.StatusForbidden, "Cannot access this reactable", nil)
		}
	} else if visibility == postVisibility.OnlyFollowers {
		followed, err := rc.userService.CheckFollow(userID, ownerID)
		if err != nil {
			return utils.NewHttpError(http.StatusInternalServerError, "Error while checking follow", err.Error())
		}

		if !followed {
			return utils.NewHttpError(http.StatusForbidden, "Cannot access this reactable", nil)

		}
	} else if visibility != postVisibility.Public {
		return utils.NewHttpError(http.StatusConflict, "Invalid visibility of a reactable", visibility)
	}

	return nil
}
