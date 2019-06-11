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

func (cc *Controller) CreateComment(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	var reqBody cmtabledtos.CreateCommentRequest
	var resBody cmtabledtos.CreateCommentResponse

	if err := c.ShouldBind(&reqBody); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Some required fields missing", err.Error())
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

	comment := models.Comment{Content: reqBody.Content, UserID: userID, CommentableID: cmtableID}
	if err := cc.commentService.Save(&comment); err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving comment", err.Error())
		return
	}

	resBody.Comment = cc.getDTOComment(&comment)
	utils.ResponseWithSuccess(c, http.StatusCreated, "Save comment successfully", resBody)
}
