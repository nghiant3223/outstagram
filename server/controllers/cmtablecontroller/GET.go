package cmtablecontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/dtos/cmtabledtos"
	postVisibility "outstagram/server/enums/postvisibility"
	"outstagram/server/models"
	"outstagram/server/utils"
)

func (cc *Controller) GetComments(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	var reqBody cmtabledtos.GetPostCommentsRequest
	var resBody cmtabledtos.GetPostCommentsResponse
	var commentable *models.Commentable
	var err error

	if err := c.ShouldBindQuery(&reqBody); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid query parameter", err.Error())
		return
	}

	cmtableID, err := utils.StringToUint(c.Param("cmtableID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid parameter", err.Error())
		return
	}

	visibility, ownerID, err := cc.commentableService.GetVisibilityByID(cmtableID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "Commentable not found", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
		return
	}

	if visibility == postVisibility.Private {
		if ownerID != userID {
			utils.ResponseWithError(c, http.StatusForbidden, "Cannot access this commentable", nil)
			return
		}
	} else if visibility == postVisibility.OnlyFollowers {
		followed, err := cc.userService.CheckFollow(userID, ownerID)
		if err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while checking follow", err.Error())
			return
		}

		if !followed {
			utils.ResponseWithError(c, http.StatusForbidden, "Cannot access this commentable", nil)
			return
		}
	} else if visibility != postVisibility.Public {
		utils.ResponseWithError(c, http.StatusConflict, "Invalid visibility of a commentable", visibility)
		return
	}


	// If limit and offset are not specified
	if reqBody.Offset == 0 && reqBody.Limit == 0 {
		commentable, err = cc.commentableService.GetComments(cmtableID)
	} else {
		commentable, err = cc.commentableService.GetCommentsWithLimit(cmtableID, reqBody.Limit, reqBody.Offset)
	}

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "Post not found", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
		return
	}

	for _, comment := range commentable.Comments {
		resBody.Comments = append(resBody.Comments, cc.getDTOComment(&comment))
	}

	resBody.CommentCount = commentable.CommentCount
	utils.ResponseWithSuccess(c, http.StatusOK, "Retrieve comments successfully", resBody)
}
