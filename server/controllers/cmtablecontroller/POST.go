package cmtablecontroller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"outstagram/server/dtos/cmtabledtos"
	"outstagram/server/models"
	"outstagram/server/utils"
)

// CreateComment create comments of a post
// User may not be able to create the post's comment due to the visibility of the post
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

	if err := cc.checkUserAuthorizationForCommentable(cmtableID, userID); err != nil {
		utils.ResponseWithError(c, err.StatusCode, err.Message, err.Data)
		return
	}

	comment := models.Comment{Content: reqBody.Content, UserID: userID, CommentableID: cmtableID}
	if err := cc.commentService.Save(&comment); err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving comment", err.Error())
		return
	}

	resBody.Comment = comment.ToDTO()
	utils.ResponseWithSuccess(c, http.StatusCreated, "Save comment successfully", resBody)
}

// CreateComment create replies of a post
// User may not be able to create the comment's reply due to the visibility of the comment
func (cc *Controller) CreateCommentReplies(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	var reqBody cmtabledtos.CreateReplyRequest
	var resBody cmtabledtos.CreateReplyResponse
	var err error

	if err := c.ShouldBind(&reqBody); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid query parameter", err.Error())
		return
	}

	cmtableID, err := utils.StringToUint(c.Param("cmtableID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid parameter", err.Error())
		return
	}

	if err := cc.checkUserAuthorizationForCommentable(cmtableID, userID); err != nil {
		utils.ResponseWithError(c, err.StatusCode, err.Message, err.Data)
		return
	}

	cmtID, err := utils.StringToUint(c.Param("cmtID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid parameter", err.Error())
		return
	}

	if !cc.commentableService.CheckHasComment(cmtableID, cmtID) {
		utils.ResponseWithError(c, http.StatusConflict, "Comment does not belong to commentable", nil)
		return
	}

	reply := models.Reply{UserID: userID, Content: reqBody.Content, CommentID: cmtID}
	if err := cc.replyService.Save(&reply); err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving reply", err.Error())
		return
	}

	resBody.Reply = reply.ToDTO()
	utils.ResponseWithSuccess(c, http.StatusCreated, "Save reply successfully", resBody)
}
