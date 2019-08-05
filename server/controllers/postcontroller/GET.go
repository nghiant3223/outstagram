package postcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/dtos/dtomodels"
	"outstagram/server/dtos/postdtos"
	"outstagram/server/models"
	"outstagram/server/utils"
)

// GetPostComments retrieves specific post
// User may not see the post due to the visibility of the post
// Returns a few of post's comments
func (pc *Controller) GetPost(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs VerifyToken middleware")
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

	dtoPost, err := pc.postService.GetDTOPost(post, userID, userID)
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
		log.Fatal("This route needs VerifyToken middleware")
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

func (pc *Controller) SearchPost(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs VerifyToken middleware")
	}

	var req postdtos.SearchPostRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	posts, err := pc.postService.Search(req.Filter)
	if err != nil {
		appError := models.NewAppError(err.Error(), "SearchPost", "Cannot search for post", nil, http.StatusInternalServerError)
		utils.ResponseWithAppError(c, appError)
		return
	}

	var dtoPosts []*dtomodels.Post

	for _, post := range posts {
		dtoPost, err := pc.postService.GetDTOPost(post, post.UserID, userID)
		if err != nil {
			appError := models.NewAppError(err.Error(), "SearchPost", "Cannot search for post", nil, http.StatusInternalServerError)
			utils.ResponseWithAppError(c, appError)
			return
		}
		dtoPosts = append(dtoPosts, dtoPost)
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Fetching user successfully", dtoPosts)
}
