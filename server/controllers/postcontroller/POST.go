package postcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/db"
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
	imageURLs := reqBody.ImageURLs

	if len(files)+len(imageURLs) < 1 {
		utils.ResponseWithError(c, http.StatusBadRequest, "Missing post's images", nil)
		return
	}

	if len(files)+len(imageURLs) == 1 {
		post := models.Post{Content: reqBody.Content, Privacy: reqBody.Visibility, UserID: userID}

		var image *models.Image
		var err error

		if len(files) > 0 {
			image, err = pc.imageService.Save(files[0], userID, false)
		} else if len(imageURLs) > 0 {
			image, err = pc.imageService.SaveURL(imageURLs[0], userID, false)
		}

		if err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving post's image", err.Error())
			return
		}

		post.ImageID = image.ID

		if err := pc.postService.Save(&post); err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving post", err.Error())
			return
		}

		savedPost, err := pc.postService.GetPostByID(post.ID, userID)
		dtoPost, err := pc.postService.GetDTOPost(savedPost, userID, userID)
		if err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving post", err.Error())
			return
		}

		resBody.Post = *dtoPost
		followers := pc.userService.GetFollowers(userID)
		for _, follower := range followers {
			redisSupplier, _ := db.NewRedisSupplier()
			redisSupplier.LPush(fmt.Sprintf("newsfeed:%v", follower.ID), savedPost.ID)
		}

		utils.ResponseWithSuccess(c, 200, "Create post successfully", resBody)
		return
	}

	post := models.Post{Content: reqBody.Content, Privacy: reqBody.Visibility, UserID: userID}
	if err := pc.postService.Save(&post); err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving post", err.Error())
		return
	}

	for _, file := range files {
		image, err := pc.imageService.Save(file, userID, false)
		if err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving post's image", err.Error())
			return
		}

		postImage := models.PostImage{ImageID: image.ID, PostID: post.ID}
		if err := pc.postImageService.Save(&postImage); err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving post's image", err.Error())
			return
		}
	}

	for _, url := range imageURLs {
		image, err := pc.imageService.SaveURL(url, userID, false)
		if err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving post's image", err.Error())
			return
		}

		postImage := models.PostImage{ImageID: image.ID, PostID: post.ID}
		if err := pc.postImageService.Save(&postImage); err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving post's image", err.Error())
			return
		}
	}

	savedPost, err := pc.postService.GetPostByID(post.ID, userID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "Post not found", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
		return
	}
	
	dtoPost, err := pc.postService.GetDTOPost(savedPost, userID, userID)
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving post", err.Error())
		return
	}

	resBody.Post = *dtoPost

	followers := pc.userService.GetFollowers(userID)
	for _, follower := range followers {
		redisSupplier, _ := db.NewRedisSupplier()
		redisSupplier.LPush(fmt.Sprintf("newsfeed:%v", follower.ID), savedPost.ID)
	}

	utils.ResponseWithSuccess(c, http.StatusCreated, "Create post successfully", resBody)
}

func (pc *Controller) ViewPost(c *gin.Context) {
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

	if err := pc.viewableService.IncrementView(userID, post.ViewableID); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "Post not found", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving view", err.Error())
		return
	}

	utils.ResponseWithSuccess(c, http.StatusCreated, "Save view successfully", nil)
}
