package storycontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/dtos/storydtos"
	"outstagram/server/models"
	"outstagram/server/utils"
)

func (sc *Controller) CreateStory(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	var reqBody storydtos.CreateStoryRequest
	var resBody storydtos.CreateStoryResponse

	if err := c.ShouldBind(&reqBody); err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Some required fields missing", err.Error())
		return
	}

	user, err := sc.userService.FindByID(userID)
	if err != nil {
		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving user", err.Error())
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid form data", err.Error())
		return
	}

	files := form.File["images"]
	if len(files) < 1 {
		utils.ResponseWithError(c, http.StatusBadRequest, "Missing story's images", nil)
		return
	}

	for _, file := range files {
		image, err := sc.imageService.Save(file, userID)
		if err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving story's image", err.Error())
			return
		}

		story := models.Story{ImageID: image.ID, StoryBoardID: user.StoryBoard.ID, Duration: 3000}
		if err := sc.storyBoardService.SaveStory(&story); err != nil {
			utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving story's image", err.Error())
			return
		}

		savedStory, _ := sc.storyBoardService.GetStoryByID(story.ID)
		resBody.Stories = append(resBody.Stories, savedStory.ToDTO())
	}

	utils.ResponseWithSuccess(c, http.StatusCreated, "Create story successfully", resBody)
}

func (sc *Controller) ViewStory(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs verifyToken middleware")
	}

	postID, err := utils.StringToUint(c.Param("storyID"))
	if err != nil {
		utils.ResponseWithError(c, http.StatusBadRequest, "Invalid parameter", err.Error())
		return
	}

	story, err := sc.storyBoardService.GetStoryByID(postID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "Post not found", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while retrieving post", err.Error())
		return
	}

	if err := sc.viewableService.SaveView(userID, story.ViewableID); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "Post not found", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while saving view", err.Error())
		return
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Save view successfully", nil)
}
