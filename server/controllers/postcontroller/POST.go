package postcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"outstagram/server/dtos/postdtos"
	"outstagram/server/models"
	"outstagram/server/utils"
)

func (pc *Controller) CreatePost(c *gin.Context) {
	var reqBody postdtos.CreatePostRequest
	var resBody postdtos.CreatePostResponse

	if err := c.ShouldBind(&reqBody); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Some required fields missing", err.Error())
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid form data", err.Error())
		return
	}

	files := form.File["images"]
	if len(files) < 1 {
		utils.ResponseWithError(c, http.StatusBadRequest, "Missing post's images", nil)
		return
	}

	userID, _ := utils.RetrieveUserID(c)
	post := models.Post{Content: reqBody.Content, Visibility: reqBody.Visibility, UserID: userID}
	if err := pc.postService.Save(&post); err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error when saving post", err.Error())
		return
	}

	resBody.Visibility = post.Visibility
	resBody.NumRead = post.NumRead
	resBody.Content = post.Content

	for _, file := range files {
		image, err := pc.imageService.Save(file, userID)
		if err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error when saving post's image", err.Error())
			return
		}

		postImage := models.PostImage{ImageID: image.ID, PostID: post.ID}
		if err := pc.postImageService.Save(&postImage); err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error when saving post's image", err.Error())
			return
		}

		resBody.PostImages = append(resBody.PostImages, postdtos.PostImage{
			ID:       postImage.ID,
			Small:    image.Small,
			Tiny:     image.Tiny,
			Original: image.Origin,
			Medium:   image.Medium,
			Huge:     image.Huge,
			Big:      image.Big,
		})
	}

	utils.ResponseWithSuccess(c, 200, "Create post successfully", resBody)
}
