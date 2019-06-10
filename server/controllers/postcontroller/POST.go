package postcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"log"
	"net/http"
	"outstagram/server/dtos/postdtos"
	"outstagram/server/models"
	"outstagram/server/utils"
)

func (pc *Controller) CreatePost(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

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

	post := models.Post{Content: reqBody.Content, Visibility: reqBody.Visibility, UserID: userID}
	if err := pc.postService.Save(&post); err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving post", err.Error())
		return
	}

	resBody.ID = post.ID
	if err = copier.Copy(&resBody, &post); err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while copying entity to dto", err.Error())
		return
	}

	for _, file := range files {
		image, err := pc.imageService.Save(file, userID)
		if err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving post's image", err.Error())
			return
		}

		postImage := models.PostImage{ImageID: image.ID, PostID: post.ID}
		if err := pc.postImageService.Save(&postImage); err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving post's image", err.Error())
			return
		}

		dtoPostImage := postdtos.PostImage{ID: postImage.ID}
		if err = copier.Copy(&dtoPostImage, &image); err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while copying entity to dto", err.Error())
			return
		}

		resBody.Images = append(resBody.Images, dtoPostImage)
	}

	utils.ResponseWithSuccess(c, 200, "Create post successfully", resBody)
}