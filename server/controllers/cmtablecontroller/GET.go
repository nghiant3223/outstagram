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

	postID, err := utils.StringToUint(c.Param("postID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid parameter", err.Error())
		return
	}

	post, err := pc.postService.GetPostByID(postID, userID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "Post not found", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
		return
	}

	// If limit and offset are not specified
	if reqBody.Offset == 0 && reqBody.Limit == 0 {
		commentable, err = pc.commentableService.GetComments(post.CommentableID)
	} else {
		commentable, err = pc.commentableService.GetCommentsWithLimit(post.CommentableID, reqBody.Limit, reqBody.Offset)
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
		resBody.Comments = append(resBody.Comments, pc.getDTOComment(&comment))
	}

	resBody.CommentCount = commentable.CommentCount
	utils.ResponseWithSuccess(c, http.StatusOK, "Retrieve comments successfully", resBody)
}