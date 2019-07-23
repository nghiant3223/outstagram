package storycontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"outstagram/server/dtos/medtos"
	"outstagram/server/utils"
)

// GetStories returns all available user's stories
func (sc *Controller) GetMyStoryBoard(c *gin.Context) {
	userID, ok := utils.RetrieveUserID(c)
	if !ok {
		log.Fatal("This route needs VerifyToken middleware")
	}

	var resBody medtos.GetMyStoryBoard

	user, err := sc.userService.FindByID(userID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "Not found", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while sorting story board", err.Error())
		return
	}

	stories, err := sc.storyBoardService.GetStories(user.StoryBoard.ID)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.ResponseWithError(c, http.StatusNotFound, "Not found", err.Error())
			return
		}

		utils.ResponseWithError(c, http.StatusInternalServerError, "Error while sorting story board", err.Error())
		return
	}

	for _, story := range stories {
		resBody.Stories = append(resBody.Stories, story.ToDTO())
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "Get story board successfully", resBody)
}
