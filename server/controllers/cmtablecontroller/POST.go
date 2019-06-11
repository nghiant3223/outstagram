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

func (cc *Controller) CreateComment(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	var reqBody cmtabledtos.CreateCommentRequest
	var resBody cmtabledtos.CreateCommentResponse

	cmtableID, err := utils.StringToUint(c.Param("cmtableID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid parameter", err.Error())
		return
	}

	if err := c.ShouldBind(&reqBody); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Some required fields missing", err.Error())
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

	comment := models.Comment{Content: reqBody.Content, UserID: userID, CommentableID: post.CommentableID}
	if err := pc.commentService.Save(&comment); err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving comment", err.Error())
		return
	}

	resBody.Comment = pc.getDTOComment(&comment)
	utils.ResponseWithSuccess(c, http.StatusCreated, "Save comment successfully", resBody)

}