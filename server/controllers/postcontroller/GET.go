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

// GetPostComments retrieves comments of a post
// User may not see the post's comment due to the visibility of the post
func (pc *Controller) GetPostComments(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	var reqBody postdtos.GetPostCommentsRequest
	var resBody postdtos.GetPostCommentsResponse
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

	post, err := pc.postService.GetPostByID(userID, postID)
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

	utils.ResponseWithSuccess(c, http.StatusOK, "Retrieve comments successfully", resBody)
}

// GetPostComments retrieves specific post
// User may not see the post due to the visibility of the post
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

	post, err := pc.postService.GetPostByID(userID, postID)
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

// getDTOPost maps basic information of post, including post's images, post's comments into a DTO object
func (pc *Controller) getDTOPost(post *models.Post) (*postdtos.Post, error) {
	// Set basic post's info
	dtoPost := postdtos.Post{
		ID:            post.ID,
		Content:       post.Content,
		Visibility:    post.Visibility,
		ImageCount:    len(post.Images),
		NumRead:       post.NumRead,
		OwnerID:       post.UserID,
		OwnerFullname: post.User.Fullname,
		ReactCount:    pc.reactableService.GetReactCount(post.ReactableID),
		Reactors:      pc.reactableService.GetReactors(post.ReactableID)}

	// Map post's images to DTO
	for _, postImage := range post.Images {
		image := postImage.Image
		dtoPostImage := postdtos.PostImage{
			ID:     postImage.ID,
			Tiny:   image.Tiny,
			Origin: image.Origin,
			Huge:   image.Huge, Big: image.Huge,
			Medium: image.Medium,
			Small:  image.Small,
		}
		dtoPost.Images = append(dtoPost.Images, dtoPostImage)
	}

	// Get post's comments
	commentable, err := pc.commentableService.GetCommentsWithLimit(post.CommentableID, 5, 0)
	if err != nil {
		return nil, err
	}

	// Mapping post's comments to DTO
	dtoPost.CommentCount = commentable.CommentCount
	for _, comment := range commentable.Comments {
		dtoComment := pc.getDTOComment(&comment)
		dtoPost.Comments = append(dtoPost.Comments, dtoComment)
	}

	return &dtoPost, nil
}

// getDTOComment maps information of a comment into a DTO object
func (pc *Controller) getDTOComment(comment *models.Comment) postdtos.Comment {
	return postdtos.Comment{
		ID:            comment.ID,
		Content:       comment.Content,
		ReplyCount:    comment.ReplyCount,
		CreatedAt:     comment.CreatedAt,
		OwnerFullname: comment.User.Fullname,
		OwnerID:       comment.UserID,
		ReactCount:    pc.reactableService.GetReactCount(comment.ReactableID),
		Reactors:      pc.reactableService.GetReactors(comment.ReactableID)}
}
