package postcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/dtos/postdtos"
	"outstagram/server/utils"
)

func (pc *Controller) GetPosts(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	var reqBody postdtos.GetPostRequest
	var resBody postdtos.GetPostResponse

	if err := c.ShouldBindQuery(&reqBody); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid query parameter", err.Error())
		return
	}

	posts, err := pc.postService.GetUsersPostsWithLimit(userID, reqBody.Limit, reqBody.Offset)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithSuccess(c, http.StatusNoContent, "No posts", nil)
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving user's posts", err.Error())
		return
	}

	for _, post := range posts {
		// Set basic post's info
		dtoPost := postdtos.Post{
			ID:         post.ID,
			Content:    post.Content,
			Visibility: post.Visibility,
			ImageCount: len(post.Images),
			NumRead:    post.NumRead,
			ReactCount: pc.reactableService.GetReactCount(post.ReactableID),
			Reactors:   pc.reactableService.GetReactors(post.ReactableID)}

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
			if gorm.IsRecordNotFoundError(err) {
				utils.ResponseWithError(c, http.StatusNotFound, "Post not found", err.Error())
				return
			}

			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
			return
		}

		// Mapping post's comments to DTO
		dtoPost.CommentCount = commentable.CommentCount
		for _, comment := range commentable.Comments {
			dtoComment := postdtos.Comment{
				ID:         comment.ID,
				Content:    comment.Content,
				ReplyCount: comment.ReplyCount,
				CreatedAt:  comment.CreatedAt,
				Fullname:   comment.User.Fullname,
				UserID:     comment.UserID,
				ReactCount: pc.reactableService.GetReactCount(comment.ReactableID),
				Reactors:   pc.reactableService.GetReactors(comment.ReactableID)}
			dtoPost.Comments = append(dtoPost.Comments, dtoComment)
		}

		resBody.Posts = append(resBody.Posts, dtoPost)
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Fetch user's posts successfully", resBody)
}
