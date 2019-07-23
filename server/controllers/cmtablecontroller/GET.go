package cmtablecontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/dtos/cmtabledtos"
	"outstagram/server/models"
	"outstagram/server/utils"
)

// GetPostComments retrieves comments of a post
// User may not see the post's comment due to the visibility of the post
func (cc *Controller) GetComments(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs VerifyToken middleware")
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

	if err := cc.checkUserAuthorizationForCommentable(cmtableID, userID); err != nil {
		utils.ResponseWithError(c, err.StatusCode, err.Message, err.Data)
		return
	}
	// If limit and offset are not specified
	if reqBody.Offset == 0 && reqBody.Limit == 0 {
		commentable, err = cc.commentableService.GetComments(cmtableID, userID)
	} else {
		commentable, err = cc.commentableService.GetCommentsWithLimit(cmtableID, userID, reqBody.Limit, reqBody.Offset)
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
		dtoComment := comment.ToDTO()
		dtoComment.Reacted = cc.reactableService.CheckUserReaction(userID, comment.ReactableID)
		resBody.Comments = append(resBody.Comments, dtoComment)
	}

	resBody.CommentCount = commentable.CommentCount
	utils.ResponseWithSuccess(c, http.StatusOK, "Retrieve comments successfully", resBody)
}

// GetCommentReplies retrieves replies of a comment
// User may not see the comment's replies due to the visibility of the comment
func (cc *Controller) GetCommentReplies(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs VerifyToken middleware")
	}

	var reqBody cmtabledtos.GetCommentRepliesRequest
	var resBody cmtabledtos.GetCommentRepliesResponse
	var comment *models.Comment
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
		utils.ResponseWithError(c, http.StatusConflict, "Comment do not belong to commentable", nil)
		return
	}

	if reqBody.Offset == 0 && reqBody.Limit == 0 {
		comment, err = cc.commentService.GetReplies(cmtID, userID)
	} else {
		comment, err = cc.commentService.GetRepliesWithLimit(cmtID, userID, reqBody.Limit, reqBody.Offset)
	}

	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving comment", err.Error())
		return
	}

	resBody.ReplyCount = comment.ReplyCount
	for _, reply := range comment.Replies {
		dtoReply := reply.ToDTO()
		dtoReply.Reacted = cc.reactableService.CheckUserReaction(userID, reply.ReactableID)
		resBody.Replies = append(resBody.Replies, dtoReply)
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Retrieve replies successfully", resBody)
}
