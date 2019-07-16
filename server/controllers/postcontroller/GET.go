package postcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/utils"
)

// GetPostComments retrieves specific post
// User may not see the post due to the visibility of the post
// Returns a few of post's comments
func (pc *Controller) GetPost(c *gin.Context) {
	audienceUserID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
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

	dtoPost, err := pc.postService.GetDTOPost(post, userID, audienceUserID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "Post not found", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
		return
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Retrieve post successfully", dtoPost)
}

func (pc *Controller) GetPostImage(c *gin.Context) {
	audienceUserID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	postImageID, err := utils.StringToUint(c.Param("postImageID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid parameter", err.Error())
		return
	}

	postImage, err := pc.postImageService.GetPostImageByID(postImageID)
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
		return
	}

	dtoPostImage, err := pc.postImageService.GetDTOPostImage(postImage, postImage.Post.UserID, audienceUserID)
	utils.ResponseWithSuccess(c, http.StatusOK, "Retrieve post successfully", dtoPostImage)

}
