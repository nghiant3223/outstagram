package postcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/dtos/postdtos"
	"outstagram/server/models"
	"outstagram/server/utils"
)

// GetPosts retrieves posts of an authenticated user
// Returns a few of post's comments
func (pc *Controller) GetPosts(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	var reqBody postdtos.GetPostRequest
	var resBody postdtos.GetPostResponse
	var posts []models.Post
	var err error

	if err := c.ShouldBindQuery(&reqBody); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid query parameter", err.Error())
		return
	}

	// If limit and offset are not specified
	if reqBody.Offset == 0 && reqBody.Limit == 0 {
		posts, err = pc.postService.GetUserPosts(userID)
	} else {
		posts, err = pc.postService.GetUsersPostsWithLimit(userID, reqBody.Limit, reqBody.Offset)
	}

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithSuccess(c, http.StatusNoContent, "No posts", nil)
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving user's posts", err.Error())
		return
	}

	for _, post := range posts {
		dtoPost, err := pc.getDTOPost(&post)
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				utils.ResponseWithError(c, http.StatusNotFound, "Post not found", err.Error())
				return
			}

			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
			return
		}

		resBody.Posts = append(resBody.Posts, *dtoPost)
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Fetch user's posts successfully", resBody)
}

// GetPostComments retrieves specific post
// User may not see the post due to the visibility of the post
// Returns a few of post's comments
func (pc *Controller) GetPost(c *gin.Context) {
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

	dtoPost, err := pc.getDTOPost(post)
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
